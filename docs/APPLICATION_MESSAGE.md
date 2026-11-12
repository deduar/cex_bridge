# Application Message (for job post)

Hi,

I've analyzed the P2P→CEX migration and built a working POC proving the hardest part: the matching engine + double-entry ledger + real-time order book.

## POC (runnable)
- Repo: docker compose up --build → gateway on :8080
- Flow: register (mock KYC) → login → place buy → place sell → match executes → WS broadcast
- Probe: docs/API.md (curl commands included)
- Proof: automated invariant test (sum of balances conserved after a match)

## Plan (3–6 week, part-time)
- Deep audit of your codebase, gap analysis, prioritized backlog
- CEX core — matching engine (Go), double-entry ledger, custody service (3 independent domains)
- User migration, KYC/AML, fiat on-ramps, admin dashboard
- Load test, security audit, Kubernetes, handover

## Key decisions
- Matching engine in Go (Rust documented as production target)
- PostgreSQL with RLS for multi-tenancy; Redis for order book
- Stateless services scale horizontally; stateful scale vertically/sharded
- Mocks used in POC for KYC/fiat/custody; production swaps in real providers

## Communication
Daily async updates, weekly demo, shared repo + issue tracker, weekly written report.

I work in sprints with full transparency. Happy to walk through the POC live.

Best regards
EOF