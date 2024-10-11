package main

import (
	"fmt"

	tarantool "github.com/tarantool/go-tarantool"
)

var cnx *tarantool.Connection = nil

func connectToDatabase(host string, port int, user string, password string) {
	address := fmt.Sprintf("%s:%d", host, port)
	var err error
	cnx, err = tarantool.Connect(address, tarantool.Opts{
		User: user,
		Pass: password,
	})
	assertNoError(err)
}

func disconnectFromDatabase() {
	cnx.Close()
}

func fetchTrades() []Trade {
	resp, err := cnx.Call("trades.list", []interface{}{})
	assertNoError(err)
	trades := []Trade{}
	tuples := resp.Tuples()

	if len(tuples) == 0 || tuples[0] == nil || len(tuples[0]) != 5 {
		return trades
	}

	for _, tuple := range tuples {
		id := tuple[0].(uint64)
		symbol := tuple[1].(string)
		direction := tuple[2].(string)
		price := convertToFloat64(tuple[3])
		quantity := convertToFloat64(tuple[4])

		trade := Trade{
			Id:        id,
			Symbol:    symbol,
			Direction: direction,
			Price:     price,
			Quantity:  quantity,
		}

		trades = append(trades, trade)
	}

	return trades
}

func fetchMarketPrices() map[string]float64 {
	resp, err := cnx.Call("market_prices.list", []interface{}{})
	assertNoError(err)
	marketPrices := map[string]float64{}
	tuples := resp.Tuples()

	if len(tuples) == 0 || tuples[0] == nil || len(tuples[0]) != 5 {
		return marketPrices
	}

	for _, tuple := range tuples {
		symbol := tuple[0].(string)
		price := convertToFloat64(tuple[1])
		marketPrices[symbol] = price
	}

	return marketPrices
}

func getMarketPrice(symbol string) float64 {
	resp, err := cnx.Call("market_prices.read", []interface{}{symbol})
	assertNoError(err)
	price := convertToFloat64(resp.Tuples()[0][0])
	return price
}

func setMarketPrice(symbol string, price float64) {
	_, err := cnx.Call("market_prices.update", []interface{}{symbol, price})
	assertNoError(err)
}

func pushOrder(symbol string, direction string, price float64, quantity float64) uint64 {
	resp, err := cnx.Call("order_queue.push", []interface{}{symbol, direction, price, quantity})
	assertNoError(err)
	orderId := resp.Tuples()[0][0]
	return orderId.(uint64)
}

func popOrder() *Order {
	resp, err := cnx.Call("order_queue.pop", []interface{}{})
	assertNoError(err)
	tuples := resp.Tuples()

	if len(tuples) == 0 || tuples[0] == nil || len(tuples[0]) != 5 {
		return nil
	}

	tuple := tuples[0]
	id := tuple[0].(uint64)
	symbol := tuple[1].(string)
	direction := tuple[2].(string)
	price := convertToFloat64(tuple[3])
	quantity := convertToFloat64(tuple[4])

	order := &Order{
		Id:        id,
		Symbol:    symbol,
		Direction: direction,
		Price:     price,
		Quantity:  quantity,
	}

	return order
}

func bookTrade(symbol string, direction string, price float64, quantity float64) uint64 {
	resp, err := cnx.Call("trades.book", []interface{}{symbol, direction, price, quantity})
	assertNoError(err)
	tradeId := resp.Tuples()[0][0]
	return tradeId.(uint64)
}

func cancelTrade(tradeId uint64) {
	_, err := cnx.Call("trades.cancel", []interface{}{tradeId})
	assertNoError(err)
}
