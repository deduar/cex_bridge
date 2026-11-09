package matching

import "sort"

type Side int

const (
	Buy Side = iota
	Sell
)

type Order struct {
	ID     string
	UserID string
	Side   Side
	Price  int64
	Amount int64
}

type Match struct {
	BuyOrderID  string
	SellOrderID string
	Price       int64
	Amount      int64
}

type Engine struct {
	buys  []Order
	sells []Order
}

func NewEngine() *Engine { return &Engine{} }

func (e *Engine) Place(o Order) []Match {
	var matches []Match
	if o.Side == Buy {
		for len(e.sells) > 0 && e.sells[0].Price <= o.Price && o.Amount > 0 {
			s := &e.sells[0]
			qty := min(o.Amount, s.Amount)
			matches = append(matches, Match{BuyOrderID: o.ID, SellOrderID: s.ID, Price: s.Price, Amount: qty})
			o.Amount -= qty
			s.Amount -= qty
			if s.Amount == 0 {
				e.sells = e.sells[1:]
			}
		}
		if o.Amount > 0 {
			e.buys = append(e.buys, o)
			sort.Slice(e.buys, func(i, j int) bool { return e.buys[i].Price > e.buys[j].Price })
		}
	} else {
		for len(e.buys) > 0 && e.buys[0].Price >= o.Price && o.Amount > 0 {
			b := &e.buys[0]
			qty := min(o.Amount, b.Amount)
			matches = append(matches, Match{BuyOrderID: b.ID, SellOrderID: o.ID, Price: b.Price, Amount: qty})
			o.Amount -= qty
			b.Amount -= qty
			if b.Amount == 0 {
				e.buys = e.buys[1:]
			}
		}
		if o.Amount > 0 {
			e.sells = append(e.sells, o)
			sort.Slice(e.sells, func(i, j int) bool { return e.sells[i].Price < e.sells[j].Price })
		}
	}
	return matches
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
