# EXPLORE

## Architecture (POC)
- services/trading: matching engine + order book (Go)
- services/ledger: balances + postings (Go)
- services/custody: mock wallet (Go)
- infra: Postgres (RLS), Redis
- gateway: REST + WebSocket (Go)
- docker-compose orchestrates all

## Rejected alternatives
- Rust matching engine: too slow to build for POC; documented as prod target.
- Kafka: unnecessary for POC; in-process channels.
- Real KYC provider: mocked.

## KISS rules (karpathy)
- No premature abstraction.
- One service boundary only where needed.
- Tests before features for the invariant.
- No admin dashboard, no fiat, no mobile.
