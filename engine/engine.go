package engine

type Engine struct {
	// Asks
	// Bids
	// Sequence
}

func NewEngine() *Engine {
	return &Engine{}
}

// AddOrder adds an order to the order book and returns a list of trades that were executed
func (e *Engine) AddOrder(order *Order) []Trade {
	return nil
}

func (e *Engine) RemoveOrder(order *Order) {
}

func (e *Engine) execute(order *Order) []Trade {
	return nil
}

func (e *Engine) PrintOrderBook() {
}
