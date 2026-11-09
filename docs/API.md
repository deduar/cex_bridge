# API Probe Guide (POC Demo)

Base URL: `http://localhost:8080`

## 1. Health
```bash
curl http://localhost:8080/health
```

## 2. Register (mock KYC)
```bash
curl -s -X POST http://localhost:8080/register \
  -H 'Content-Type: application/json' \
  -d '{"username":"alice"}'
```

## 3. Login
```bash
curl -s -X POST http://localhost:8080/login \
  -H 'Content-Type: application/json' \
  -d '{"username":"alice"}'
```
Save the returned `token`.

## 4. Place buy order (authenticated)
```bash
TOKEN="<paste token>"
curl -s -X POST http://localhost:8080/orders \
  -H "Authorization: Bearer $TOKEN" \
  -H 'Content-Type: application/json' \
  -d '{"id":"b1","user_id":"alice","side":"buy","price":100,"amount":10}'
```

## 5. Place sell order (triggers match)
```bash
curl -s -X POST http://localhost:8080/orders \
  -H "Authorization: Bearer $TOKEN" \
  -H 'Content-Type: application/json' \
  -d '{"id":"s1","user_id":"bob","side":"sell","price":95,"amount":4}'
```
Expected: `{"matches":[{"BuyOrderID":"b1","SellOrderID":"s1","Price":100,"Amount":4}]}`

## 6. WebSocket order book feed
```bash
# using websocat (optional)
websocat ws://localhost:8080/ws
```

## Demo flow summary
Register (mock KYC) → login → place buy → place sell → match executes → WS broadcast.
EOF