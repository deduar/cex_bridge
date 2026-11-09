# DISCOVER

## Problem
Migrate P2P system to CEX platform. POC must prove hardest part: matching engine + ledger.

## Anchors (from DECISIONS.md)
- Matching engine (Go)
- Double-entry ledger + invariant test
- WebSocket order book
- Mock KYC gate
- Postgres RLS + Redis
- docker compose runnable

## Constraints
- Part-time, 1-3 months total; POC 1-2 weeks.
- KISS: smallest thing that proves capability.
- Docker networking: use a custom bridge network (no default gateway) to avoid docker compose hijacking host default route / breaking internet (lesson from crownForge_TI).

## Open questions resolved
- Go over Rust for POC (Rust = production target).
- Redis included (low cost in compose) with in-memory fallback documented.
- Kafka skipped; in-process channels.
