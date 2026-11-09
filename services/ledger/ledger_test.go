package ledger

import "testing"

func TestInvariantSumConserved(t *testing.T) {
	l := New()
	l.Deposit("alice", 1000)
	l.Deposit("bob", 0)
	before := l.Total()
	if err := l.Transfer("alice", "bob", 300); err != nil {
		t.Fatal(err)
	}
	after := l.Total()
	if before != after {
		t.Fatalf("invariant broken: before=%d after=%d", before, after)
	}
	if l.Balance("alice") != 700 || l.Balance("bob") != 300 {
		t.Fatalf("bad balances: alice=%d bob=%d", l.Balance("alice"), l.Balance("bob"))
	}
}
