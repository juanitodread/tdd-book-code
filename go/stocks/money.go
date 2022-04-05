package stocks

type Money struct {
	amount   float64
	currency string
}

func NewMoney(amount float64, currency string) Money {
	return Money{
		amount:   amount,
		currency: currency,
	}
}

func (money Money) Times(multiplier float64) Money {
	return Money{
		amount:   money.amount * multiplier,
		currency: money.currency,
	}
}

func (money Money) Divide(divisor float64) Money {
	return Money{
		amount:   money.amount / divisor,
		currency: money.currency,
	}
}

func (money Money) Add(other *Money) *Money {
	var result Money

	if money.currency != other.currency {
		return nil
	}

	result = Money{
		amount:   money.amount + other.amount,
		currency: money.currency,
	}

	return &result
}
