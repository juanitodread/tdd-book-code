package stocks

import "errors"

type Bank struct {
	exchangeRates map[string]float64
}

func (bank Bank) AddExchangeRate(currencyFrom string, currencyTo string, rate float64) {
	exchange := buildExchangeName(currencyFrom, currencyTo)
	bank.exchangeRates[exchange] = rate
}

func (bank Bank) Convert(money Money, currencyTo string) (*Money, error) {
	var moneyResult Money
	if money.currency == currencyTo {
		moneyResult = NewMoney(money.amount, money.currency)
		return &moneyResult, nil
	}

	exchange := buildExchangeName(money.currency, currencyTo)
	rate, isValidExchange := bank.exchangeRates[exchange]

	if !isValidExchange {
		return nil, errors.New(exchange)
	}

	moneyResult = NewMoney(money.amount*rate, currencyTo)
	return &moneyResult, nil
}

func NewBank() Bank {
	return Bank{
		exchangeRates: make(map[string]float64),
	}
}

func buildExchangeName(currencyFrom string, currencyTo string) string {
	return currencyFrom + "->" + currencyTo
}
