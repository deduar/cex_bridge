# CEX POC — Decision Log (Anchors)

## Goal
POC for job application: demonstrate P2P→CEX migration capability. Mocks allowed. Runnable repo + short recording.

## Anchors (non-negotiable, real, not mocked)
1. Matching engine core — order intake → match → fill (price/amount/timestamp). Go.
2. Double-entry ledger — balances update on fill; invariant test (sum conserved).
3. Live order book via WebSocket depth feed.
4. KYC gate (mocked provider) gating account creation.
5. Postgres with RLS (multi-tenancy) + Redis order book.
6. docker compose up → runnable; seed + demo script.

## Mocked (acceptable)
Fiat ramps, real custody hot/cold, admin dashboard, real KYC provider, Kafka, Kubernetes, production migrations.

## Stack decisions
- Matching engine: Go (POC speed); Rust documented as production target.
- Ledger store: Postgres (Docker) with RLS.
- Order book: Redis (with documented in-memory fallback).
- Auth: JWT (simple).
- Messaging: in-process channels (Kafka/Redpanda documented as production).
- Infra: docker-compose for all.

## POC flow
register (mock KYC) → deposit → place buy → place sell → match → balances update → WS depth.

## Scope cuts (NOT in POC)
Admin dashboard, fiat, mobile, full KYC provider, Kafka, K8s, full migration tooling.

## Proof criterion
Automated test: two users, opposing orders, match executes, ledger invariant holds (sum of balances conserved).
