# Evidence Log — CEX POC

## M1 — Infra skeleton
- docker compose up: postgres + redis + gateway healthy
- gateway /health returns ok
- custom bridge network (no default-route hijack)

## M2 — Ledger
- invariant test: sum of balances conserved after transfer
- test file: services/ledger/ledger_test.go

## Pending
- M3 matching engine
- M4 Redis order book + WS
- M5 gateway + mock KYC + JWT
- M6 seed + demo script + recording

## M3 — Matching engine
- in-memory order book + price-time matching
- test: services/matching/matching_test.go

## M3 probe
- POST /orders buy then sell -> match returned (amount 4 @ 100)
- endpoints: /health, /orders, /orderbook

## M4 — WebSocket feed
- /ws endpoint upgrades; /orders broadcasts matches to clients
- deps: gorilla/websocket (go-redis listed for M5)

## M5 — Mock KYC + auth
- /register (mock KYC approved) and /login; bearer token required on /orders
