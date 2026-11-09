# PLAN

## Milestones
1. M1: docker-compose skeleton + Postgres/Redis + Go module init (day 1-2)
2. M2: Ledger service + invariant unit test (day 3-4)
3. M3: Matching engine + in-memory order book (day 5-7)
4. M4: Redis order book + WebSocket depth feed (day 8-9)
5. M5: REST gateway + mock KYC + JWT auth (day 10-11)
6. M6: Seed + demo script + recording (day 12-14)

## Definition of done (POC)
- `docker compose up` runs everything
- Demo flow executes end-to-end
- Invariant test passes
- 3-5 min recording captured
- README with narrative + decisions

## Out of scope (KISS)
Admin dashboard, fiat, real KYC, Kafka, K8s, full migration tooling.
