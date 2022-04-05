package stocks

import "errors"

type Portfolio []Money

func (portfolio Portfolio) Add(money Money) Portfolio {
	portfolio = append(portfolio, money)
	return portfolio
}

func (portfolio Portfolio) Evaluate(bank Bank, currency string) (*Money, error) {
	totalMoney := NewMoney(0, currency)
	failedConversions := make([]string, 0)

	for _, money := range portfolio {
		convertedCurrency, err := bank.Convert(money, currency)
		if err == nil {
			totalMoney = *totalMoney.Add(convertedCurrency)
		} else {
			failedConversions = append(failedConversions, err.Error())
		}
	}

	if len(failedConversions) > 0 {
		failures := "["

		for _, fail := range failedConversions {
			failures = failures + fail + ", "
		}

		failures = failures[:len(failures)-2]

		failures = failures + "]"

		return nil, errors.New("Missing exchange rates: " + failures)
	}

	return &totalMoney, nil
}
