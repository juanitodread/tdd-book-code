package stocks

import (
	"fmt"
)

type Bank struct {
	exchangeRates map[string]float64
}

func NewBank() Bank {
	return Bank{
		exchangeRates: make(map[string]float64),
	}
}

//@TODO: Make ExchangeRate struct
func (bank Bank) AddExchangeRate(from Currency, to Currency, rate float64) error {
	if !from.exists() {
		return fmt.Errorf("invalid currency: from = [%s]", from)
	}
	if !to.exists() {
		return fmt.Errorf("invalid currency: to = [%s]", to)
	}

	exchange := buildExchangeName(from, to)
	bank.exchangeRates[exchange] = rate

	return nil
}

func (bank Bank) Convert(money Money, to Currency) (*Money, error) {
	var moneyResult Money

	if !to.exists() {
		return nil, fmt.Errorf("invalid currency: to = [%s]", to)
	}

	if money.currency == to {
		moneyResult = NewMoney(money.amount, money.currency)
		return &moneyResult, nil
	}

	exchange := buildExchangeName(money.currency, to)
	rate, isValidExchange := bank.exchangeRates[exchange]

	if !isValidExchange {
		return nil, fmt.Errorf("exchange not supported: [%s]", exchange)
	}

	moneyResult = NewMoney(money.amount*rate, to)
	return &moneyResult, nil
}

func buildExchangeName(from Currency, to Currency) string {
	return fmt.Sprintf("%s->%s", from, to)
}
