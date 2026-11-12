Repo: https://github.com/deduar/cex_bridge

POC demonstrating the hardest part of migrating a P2P trading system to a centralized exchange (CEX): **matching engine + double-entry ledger + real-time order book**.

## Quick start
```bash
docker compose up --build
# seed demo users
make seed
# open http://localhost:8080
```

## Demo flow
1. Register (mock KYC gate)
2. Deposit test funds
3. Place a buy order
4. Place a sell order
5. Match executes → balances update
6. Live order book depth via WebSocket

## Proof criterion
Automated test verifies the ledger invariant: sum of balances conserved after a match.

## Stack
Go (matching engine), PostgreSQL (RLS), Redis (order book), Docker Compose.

## Scope cuts (KISS)
No admin dashboard, fiat ramps, real KYC provider, Kafka, or Kubernetes in this POC. Documented as production targets.

## Docs
- `.planning/DECISIONS.md` — anchors and decisions
- `.planning/` — DISCOVER / EXPLORE / PLAN / KISS / SUMMARY
# cex_bridge
