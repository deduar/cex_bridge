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
