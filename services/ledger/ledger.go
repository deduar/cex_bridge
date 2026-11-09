package ledger

import "fmt"

type Account struct {
	UserID string
	Balance int64 // minor units
}

type Ledger struct {
	accounts map[string]*Account
}

func New() *Ledger { return &Ledger{accounts: map[string]*Account{}} }

func (l *Ledger) Deposit(userID string, amount int64) {
	if _, ok := l.accounts[userID]; !ok {
		l.accounts[userID] = &Account{UserID: userID}
	}
	l.accounts[userID].Balance += amount
}

func (l *Ledger) Transfer(from, to string, amount int64) error {
	if l.accounts[from].Balance < amount {
		return fmt.Errorf("insufficient funds")
	}
	l.accounts[from].Balance -= amount
	l.accounts[to].Balance += amount
	return nil
}

func (l *Ledger) Total() int64 {
	var t int64
	for _, a := range l.accounts {
		t += a.Balance
	}
	return t
}

func (l *Ledger) Balance(userID string) int64 {
	if a, ok := l.accounts[userID]; ok {
		return a.Balance
	}
	return 0
}
