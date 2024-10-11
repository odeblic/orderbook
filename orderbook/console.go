package main

import (
	"fmt"
	"strings"
)

const (
	BLACK   = 0
	RED     = 1
	GREEN   = 2
	YELLOW  = 3
	BLUE    = 4
	MAGENTA = 5
	CYAN    = 6
	WHITE   = 7
)

func printSpace() {
	fmt.Println()
}

func printTitle(text string) {
	text = fmt.Sprintf("[\x1b[3%dm%s\x1b[0m]", CYAN, text)
	fmt.Println(text)
}

func printLine() {
	color := WHITE
	text := strings.Repeat("-", 54)
	text = fmt.Sprintf("\x1b[3%dm%s\x1b[0m", color, text)
	fmt.Println(text)
}

func printMessage(text string) {
	text = fmt.Sprintf("> %s", text)
	fmt.Println(text)
}

func printPosition(symbol string, amount float64) {
	var label string

	if amount < 0.0 {
		label = fmt.Sprintf("%s (short):", symbol)
		amount = -amount
	} else if amount > 0.0 {
		label = fmt.Sprintf("%s (long):", symbol)
	} else {
		label = fmt.Sprintf("%s:       ", symbol)
	}

	text := fmt.Sprintf("%s\r\t\t\x1b[3%dm%16.2f\x1b[0m", label, MAGENTA, amount)
	fmt.Println(text)
}

func printPrice(name string, value float64) {
	text := fmt.Sprintf("%s:\r\t\t\x1b[3%dm%16.2f\x1b[0m", name, MAGENTA, value)
	fmt.Println(text)
}

func printPercentage(name string, value float64) {
	text := fmt.Sprintf("%s:\r\t\t\x1b[3%dm%16.2f%%\x1b[0m", name, MAGENTA, value)
	fmt.Println(text)
}

func printOrder(order *Order) {
	direction := fmt.Sprintf("\x1b[3%dm%s\x1b[0m", directionColor(order.Direction), order.Direction)
	price := fmt.Sprintf("\x1b[3%dm%14.4f\x1b[0m", YELLOW, order.Price)
	quantity := fmt.Sprintf("\x1b[3%dm%14.4f\x1b[0m", YELLOW, order.Quantity)
	text := fmt.Sprintf("#%d\t%s\t%s\t%s  %s", order.Id, order.Symbol, direction, price, quantity)
	fmt.Println(text)
}

func printTrade(trade *Trade) {
	direction := fmt.Sprintf("\x1b[3%dm%s\x1b[0m", directionColor(trade.Direction), trade.Direction)
	price := fmt.Sprintf("\x1b[3%dm%14.4f\x1b[0m", YELLOW, trade.Price)
	quantity := fmt.Sprintf("\x1b[3%dm%14.4f\x1b[0m", YELLOW, trade.Quantity)
	text := fmt.Sprintf("#%d\t%s\t%s\t%s  %s", trade.Id, trade.Symbol, direction, price, quantity)
	fmt.Println(text)
}

func directionColor(direction string) int {
	switch direction {
	case "buy":
		return GREEN
	case "sell":
		return RED
	default:
		return WHITE
	}
}
