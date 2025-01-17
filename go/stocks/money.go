package stocks

import (
	"fmt"
)

type Money struct {
	amount   float64
	currency Currency
}

func NewMoney(amount float64, currency Currency) Money {
	if !currency.exists() {
		fmt.Printf("invalid currency: [%s]", currency)
	}

	return Money{
		amount:   amount,
		currency: currency,
	}
}

func NewDollar(amount float64) Money {
	return NewMoney(amount, Usd)
}

func NewEuro(amount float64) Money {
	return NewMoney(amount, Eur)
}

func NewWon(amount float64) Money {
	return NewMoney(amount, Krw)
}

func (money Money) Times(multiplier float64) (*Money, error) {
	if multiplier < 0 {
		return nil, fmt.Errorf("invalid multiplier: Negative value [%.2f]", multiplier)
	}

	return &Money{
		amount:   money.amount * multiplier,
		currency: money.currency,
	}, nil
}

func (money Money) Divide(divisor float64) (*Money, error) {
	if divisor == 0 {
		return nil, fmt.Errorf("invalid divisor: Zero value [0]")
	} else if divisor < 0 {
		return nil, fmt.Errorf("invalid divisor: Negative value [%.2f]", divisor)
	}

	return &Money{
		amount:   money.amount / divisor,
		currency: money.currency,
	}, nil
}

func (money Money) Add(other *Money) (*Money, error) {
	if money.currency != other.currency {
		return nil, fmt.Errorf("incompatible currencies: money1[%s] != money2[%s]", money.currency, other.currency)
	}

	return &Money{
		amount:   money.amount + other.amount,
		currency: money.currency,
	}, nil
}
