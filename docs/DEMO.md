# Demo Script (3-5 min recording)

1. docker compose up --build
2. make seed (or run seed script)
3. Browser: register (mock KYC) → login
4. Deposit test funds (alice 1000, bob 0)
5. Place buy order (alice) → place sell order (bob)
6. Match executes → balances update (alice 700, bob 300)
7. Show live order book depth via WebSocket
8. Show invariant test passing in CI
