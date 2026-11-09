package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/gorilla/websocket"

	"cex/gateway/matching"
)

var (
	engine = matching.NewEngine()
	mu     sync.Mutex
	upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}
	clients = map[*websocket.Conn]bool{}
	cmu     sync.Mutex
	users   = map[string]string{} // user -> mock token
)

type placeReq struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
	Side   string `json:"side"`
	Price  int64  `json:"price"`
	Amount int64  `json:"amount"`
}

type authReq struct {
	Username string `json:"username"`
}

func tokenFor(u string) string {
	b, _ := json.Marshal(map[string]string{"user": u})
	return base64.StdEncoding.EncodeToString(b)
}

func broadcast(v any) {
	cmu.Lock()
	defer cmu.Unlock()
	for c := range clients {
		c.WriteJSON(v)
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "CEX POC gateway")
	})
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "ok")
	})
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		var req authReq
		json.NewDecoder(r.Body).Decode(&req)
		if req.Username == "" {
			http.Error(w, "username required", http.StatusBadRequest)
			return
		}
		// mock KYC gate: accept any non-empty username
		tok := tokenFor(req.Username)
		users[req.Username] = tok
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"token": tok, "kyc": "approved (mock)"})
	})
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		var req authReq
		json.NewDecoder(r.Body).Decode(&req)
		tok, ok := users[req.Username]
		if !ok {
			http.Error(w, "unknown user", http.StatusUnauthorized)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"token": tok})
	})
	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		auth := r.Header.Get("Authorization")
		if !strings.HasPrefix(auth, "Bearer ") {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		var req placeReq
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		side := matching.Buy
		if req.Side == "sell" {
			side = matching.Sell
		}
		mu.Lock()
		matches := engine.Place(matching.Order{ID: req.ID, UserID: req.UserID, Side: side, Price: req.Price, Amount: req.Amount})
		mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{"matches": matches})
		broadcast(map[string]any{"type": "matches", "data": matches})
	})
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		cmu.Lock()
		clients[c] = true
		cmu.Unlock()
	})
	http.ListenAndServe(":8080", nil)
}