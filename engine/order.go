package engine

import "github.com/shopspring/decimal"

type OrderSide int

const (
	OrderSideBuy OrderSide = iota
	OrderSideSell
)

type OrderType int

const (
	OrderTypeLimit OrderType = iota
	OrderTypeMarket
)

type Order struct {
	ID int64
}

func (o Order) IsLimit() bool {
	return false
}

func (o Order) IsMarket() bool {
	return false
}

func (o Order) Fill(decimal.Decimal) {
}

func (o Order) IsFilled() bool {
	return false
}

func (o Order) String() string {
	return "print order here"
}
