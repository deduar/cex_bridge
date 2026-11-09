# Testing

- Ledger invariant: services/ledger/ledger_test.go
- CI: .github/workflows/ci.yml (go vet, go test, docker build)
- To run locally: cd services/ledger && go test ./...
