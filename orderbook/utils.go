package main

import (
	"fmt"
	"log"
	"reflect"
	"time"
)

func assertNoError(e error) {
	if e != nil {
		log.Fatalf("cannot recovered from this: %v", e)
	}
}

func convertToFloat64(value interface{}) float64 {
	switch raw := value.(type) {
	case uint64:
		return float64(raw)
	case float64:
		return raw
	default:
		assertNoError(fmt.Errorf("unsupported data type: %s", reflect.TypeOf(value).String()))
		return 0
	}
}

func populateDatabase() {
	depositCash(1000.0)
	setMarketPrice("BTC", 62000.00)
	setMarketPrice("ETH", 2410.00)

	pushOrder("ETH", "buy", 2410.0, 10)
	pushOrder("ETH", "buy", 2400.0, 20)
	pushOrder("BTC", "sell", 62000.00, 5)
	pushOrder("ETH", "sell", 2423.0, 30)
	pushOrder("BTC", "buy", 61900.00, 5)

	bookTrade("BTC", "buy", 61800.00, 2)
	bookTrade("BTC", "sell", 63300.00, 1)
	bookTrade("BTC", "sell", 63300.00, 1)
	bookTrade("ETH", "sell", 4250.00, 3)
	bookTrade("ETH", "buy", 4150.00, 4)
}

func pause() {
	time.Sleep(2 * time.Second)
}
