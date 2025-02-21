package engine

type TradeType int

const (
	TradeTypeBuy TradeType = iota
	TradeTypeSell
)

type Trade struct{}

func (t Trade) String() string {
	return "print trade here"
}
