package matching

import "testing"

func TestMatch(t *testing.T) {
	e := NewEngine()
	e.Place(Order{ID: "b1", UserID: "alice", Side: Buy, Price: 100, Amount: 10})
	m := e.Place(Order{ID: "s1", UserID: "bob", Side: Sell, Price: 95, Amount: 4})
	if len(m) != 1 || m[0].Amount != 4 {
		t.Fatalf("expected 1 match of 4, got %+v", m)
	}
}
