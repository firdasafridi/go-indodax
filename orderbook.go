package indodax

//
// OrderBook contains the data from order open buy(bid) and sell(ask).
//
type OrderBook struct {
	Buys  []*Order `json:"buy"`
	Sells []*Order `json:"sell"`
}
