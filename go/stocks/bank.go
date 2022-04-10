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

func (bank Bank) AddExchangeRate(exchangeRate ExchangeRate) {
	bank.exchangeRates[exchangeRate.Id()] = exchangeRate.rate
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
