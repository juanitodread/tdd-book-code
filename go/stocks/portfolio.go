package stocks

import (
	"errors"
	"fmt"
)

type Portfolio []Money

func (portfolio Portfolio) Add(money Money) Portfolio {
	portfolio = append(portfolio, money)
	return portfolio
}

func (portfolio Portfolio) Evaluate(bank Bank, currency Currency) (*Money, error) {
	if !currency.exists() {
		return nil, fmt.Errorf("invalid currency: [%s]", currency)
	}

	totalMoney := NewMoney(0, currency)
	failedConversions := make([]string, 0)

	for _, money := range portfolio {
		convertedCurrency, err := bank.Convert(money, currency)
		if err == nil {
			sum, _ := totalMoney.Add(convertedCurrency)
			totalMoney = *sum
		} else {
			failedConversions = append(failedConversions, err.Error())
		}
	}

	if len(failedConversions) > 0 {
		failures := ""

		for _, fail := range failedConversions {
			failures = failures + fail + ", "
		}

		failures = failures[:len(failures)-2]

		return nil, errors.New(failures)
	}

	return &totalMoney, nil
}
