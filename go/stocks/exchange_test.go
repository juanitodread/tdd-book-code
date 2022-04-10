package stocks_test

import (
	"tdd/stocks"
	"testing"
)

func TestNewExchangeRate(t *testing.T) {
	exchangeRate, err := stocks.NewExchangeRate(stocks.Eur, stocks.Krw, 1.3)

	assertNotNil(t, exchangeRate)
	assertNil(t, err)
}

func TestNewExchangeRateInvalidFrom(t *testing.T) {
	exchangeRate, err := stocks.NewExchangeRate("invalid", stocks.Krw, 1.3)

	assertEquals(t, "invalid currency: from = [invalid]", err.Error())
	assertNil(t, exchangeRate)
}

func TestNewExchangeRateInvalidTo(t *testing.T) {
	exchangeRate, err := stocks.NewExchangeRate(stocks.Eur, "invalid", 1.3)

	assertEquals(t, "invalid currency: to = [invalid]", err.Error())
	assertNil(t, exchangeRate)
}

func TestNewExchangeRateInvalidRate(t *testing.T) {
	exchangeRate, err := stocks.NewExchangeRate(stocks.Eur, stocks.Krw, 0)

	assertEquals(t, "invalid rate. Negative or zero: [0.00]", err.Error())
	assertNil(t, exchangeRate)
}

func TestNewExchangeRateSameCurrencies(t *testing.T) {
	exchangeRate, err := stocks.NewExchangeRate(stocks.Usd, stocks.Usd, 2.3)

	assertEquals(t, "invalid currency. Dupplicated currencies: from = [USD], to = [USD]", err.Error())
	assertNil(t, exchangeRate)
}

func TestExchangeRateId(t *testing.T) {
	exchangeRate, err := stocks.NewExchangeRate(stocks.Eur, stocks.Krw, 1.3)
	exchangeRateId := *exchangeRate

	assertEquals(t, "EUR->KRW", exchangeRateId.Id())
	assertNil(t, err)
}
