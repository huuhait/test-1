package main

import "github.com/huuhait/test-1/engine"

func main() {
	e := engine.NewEngine()

	order := &engine.Order{}

	e.AddOrder(order)
	e.RemoveOrder(order)
}
