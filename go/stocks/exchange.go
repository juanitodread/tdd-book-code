package stocks

import "fmt"

type ExchangeRate struct {
	from Currency
	to   Currency
	rate float64
}

func NewExchangeRate(from Currency, to Currency, rate float64) (*ExchangeRate, error) {
	if !from.exists() {
		return nil, fmt.Errorf("invalid currency: from = [%s]", from)
	}
	if !to.exists() {
		return nil, fmt.Errorf("invalid currency: to = [%s]", to)
	}
	if rate <= 0 {
		return nil, fmt.Errorf("invalid rate. Negative or zero: [%.2f]", rate)
	}
	if from == to {
		return nil, fmt.Errorf("invalid currency. Dupplicated currencies: from = [%s], to = [%s]", from, to)
	}

	return &ExchangeRate{
		from: from,
		to:   to,
		rate: rate,
	}, nil
}

func (eR ExchangeRate) Id() string {
	return fmt.Sprintf("%s->%s", eR.from, eR.to)
}
