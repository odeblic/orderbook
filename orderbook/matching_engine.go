package main

import (
	"fmt"
)

type Order struct {
	Id        uint64
	Symbol    string
	Direction string
	Price     float64
	Quantity  float64
}

type Trade struct {
	Id        uint64
	Symbol    string
	Direction string
	Price     float64
	Quantity  float64
}

type Balance struct {
	Cash          float64
	RealizedPnL   float64
	UnrealizedPnL float64
	Equity        float64
	Margin        float64
	Positions     map[string]float64
	MarketPrices  map[string]float64
}

type Balance2 struct {
	Cash          float64
	RealizedPnL   float64
	UnrealizedPnL float64
	Equity        float64
	Margin        float64
	Instruments   map[string]Instrument
	MarketPrices  map[string]float64
}

type Instrument struct {
	positionQuantity float64
	positionCost     float64

	TotalNotionalBuy  float64
	TotalNotionalSell float64
	AverageBuyPrice   float64
	AverageSellPrice  float64
	Position          float64
	RealizedPnL       float64
	UnrealizedPnL     float64
}

// var balancex Balance
var balance = 0.0

func depositCash(amount float64) {
	balance += amount
}

func withdrawCash(amount float64) float64 {
	if amount <= balance {
		balance -= amount
		return amount
	} else {
		balance = 0.0
		return balance
	}
}

func calculatePositions(trades []Trade) map[string]float64 {
	var positions = map[string]float64{}

	for _, trade := range trades {
		if trade.Direction == "buy" {
			positions[trade.Symbol] += trade.Quantity
		} else {
			positions[trade.Symbol] -= trade.Quantity
		}
	}

	return positions
}

func calculatePnL(trades []Trade, marketPrices map[string]float64) (float64, float64) {
	var realizedPnL = 0.0
	var unrealizedPnL = 0.0
	var instruments = make(map[string]*Instrument)

	//fmt.Println("DIRECTION\tPRICE\t\tQUANTITY")

	for _, trade := range trades {
		//fmt.Printf("%s\t\t%.2f\t\t%.2f\n", trade.Direction, trade.Price, trade.Quantity)
		instrument, present := instruments[trade.Symbol]

		if !present {
			instrument = &Instrument{}
			instruments[trade.Symbol] = instrument
		}

		if trade.Direction == "buy" {
			if instrument.positionQuantity >= 0 {
				// Opening or increasing a long position
				instrument.positionCost += trade.Price * trade.Quantity
				instrument.positionQuantity += trade.Quantity
			} else {
				entryPrice := instrument.positionCost / -instrument.positionQuantity

				// Closing or decreasing a short position
				if -instrument.positionQuantity >= trade.Quantity {
					instrument.RealizedPnL += (entryPrice - trade.Price) * trade.Quantity
					instrument.positionCost -= entryPrice * trade.Quantity
					instrument.positionQuantity += trade.Quantity
				} else {
					// Opening a long position if not enough
					realizedPnL += (entryPrice - trade.Price) * instrument.positionQuantity
					remainingQuantity := trade.Quantity - -instrument.positionQuantity
					// TODO: instrument.RealizedPnL -= trade.Price * remainingQuantity
					instrument.positionCost = trade.Price * remainingQuantity
					instrument.positionQuantity = remainingQuantity
				}
			}
		} else if trade.Direction == "sell" {
			if instrument.positionQuantity <= 0 {
				// Opening or increasing a short position
				instrument.positionCost += trade.Price * trade.Quantity
				instrument.positionQuantity -= trade.Quantity
			} else {
				entryPrice := instrument.positionCost / instrument.positionQuantity

				// Closing or decreasing a long position
				if instrument.positionQuantity >= trade.Quantity {
					instrument.RealizedPnL += (trade.Price - entryPrice) * trade.Quantity
					instrument.positionCost -= entryPrice * trade.Quantity
					instrument.positionQuantity -= trade.Quantity
				} else {
					// Opening a short position if not enough
					realizedPnL += (trade.Price - entryPrice) * instrument.positionQuantity
					remainingQuantity := trade.Quantity - instrument.positionQuantity
					// TODO: instrument.RealizedPnL -= trade.Price * remainingQuantity
					instrument.positionCost = trade.Price * remainingQuantity
					instrument.positionQuantity = -remainingQuantity
				}
			}
		}
	}

	printSpace()
	printTitle("Instruments")
	for symbol, instrument := range instruments {
		fmt.Printf("%s:  %+v\n", symbol, instrument)
		marketPrice := marketPrices[symbol]

		if instrument.positionQuantity > 0 {
			// Long position
			quantity := instrument.positionQuantity
			entryPrice := instrument.positionCost / quantity
			instrument.UnrealizedPnL = (marketPrice - entryPrice) * quantity
		} else if instrument.positionQuantity < 0 {
			// Short position
			quantity := -instrument.positionQuantity
			entryPrice := instrument.positionCost / quantity
			instrument.UnrealizedPnL = (entryPrice - marketPrice) * quantity
		}

		unrealizedPnL += instrument.UnrealizedPnL
		realizedPnL += instrument.RealizedPnL
	}

	return realizedPnL, unrealizedPnL
}

func calculateNotional(positions map[string]float64, marketPrices map[string]float64) float64 {
	var notional = 0.0

	for symbol, amount := range positions {
		notional += amount * marketPrices[symbol]
	}

	return notional
}

func checkAccount() bool {
	var marketPrices = fetchMarketPrices()
	var trades = fetchTrades()
	var positions = calculatePositions(trades)

	printSpace()
	printTitle("Market prices")
	for symbol, price := range marketPrices {
		printPrice(symbol, price)
	}

	printSpace()
	printTitle("Trades")
	for _, trade := range trades {
		printTrade(&trade)
	}

	printSpace()
	printTitle("Positions")
	for symbol, amount := range positions {
		printPosition(symbol, amount)
	}

	var realizedPnL, unrealizedPnL = calculatePnL(trades, marketPrices)
	var notional = calculateNotional(positions, marketPrices)
	var equity = balance + unrealizedPnL
	var actualMargin = equity / notional
	const requiredMargin = 0.10

	printSpace()
	printTitle("Metrics")
	printPrice("Cash balance", balance)
	printPrice("Notional", notional)
	printPrice("Realized PnL", realizedPnL)
	printPrice("Unrealized PnL", unrealizedPnL)
	printPrice("Equity", equity)
	printPercentage("Actual margin", actualMargin*100)
	printPercentage("Required margin", requiredMargin*100)

	return actualMargin >= requiredMargin
}

func processOrder(order *Order) {
	printOrder(order)

	if order.Direction != "buy" && order.Direction != "sell" {
		fmt.Printf("unknown direction: %s)\n", order.Direction)
		return
	}

	var tradeId = bookTrade(order.Symbol, order.Direction, order.Price, order.Quantity)

	if !checkAccount() {
		printMessage("Cancel trade because of insufficiant margin")
		cancelTrade(tradeId)
	} else {
		printMessage("Trade booked and market price updated")
		setMarketPrice(order.Symbol, order.Price)
	}
}

func consumeOrderQueue() {
	order := popOrder()

	if order != nil {
		printSpace()
		printLine()
		processOrder(order)
	}
}
