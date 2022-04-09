package stocks_test

import (
	"tdd/stocks"
	"testing"
)

var bank stocks.Bank

func initExchangeRates() {
	bank = stocks.NewBank()
	bank.AddExchangeRate("EUR", "USD", 1.2)
	bank.AddExchangeRate("USD", "KRW", 1100)
}

func TestNewBank(t *testing.T) {
	bank := stocks.NewBank()

	assertNotNil(t, bank)
}

func TestAddExchangeRate(t *testing.T) {
	bank := stocks.NewBank()

	err := bank.AddExchangeRate(stocks.Eur, stocks.Usd, 1.3)

	assertNil(t, err)
}

func TestAddExchangeRateInvalidFrom(t *testing.T) {
	bank := stocks.NewBank()

	err := bank.AddExchangeRate("invalid", stocks.Usd, 1.3)

	assertEquals(t, "invalid currency: from = [invalid]", err.Error())
}

func TestAddExchangeRateInvalidTo(t *testing.T) {
	bank := stocks.NewBank()

	err := bank.AddExchangeRate(stocks.Eur, "invalid", 0.7)

	assertEquals(t, "invalid currency: to = [invalid]", err.Error())
}

func TestConversionWithDifferentRatesBetweenTwoCurrencies(t *testing.T) {
	initExchangeRates()
	tenEuros := stocks.NewEuro(10)

	eurosToDollars, err := bank.Convert(tenEuros, stocks.Usd)

	assertNil(t, err)
	assertEquals(t, stocks.NewDollar(12), *eurosToDollars)

	bank.AddExchangeRate(stocks.Eur, stocks.Usd, 1.3)

	eurosToDollars, err = bank.Convert(tenEuros, stocks.Usd)

	assertNil(t, err)
	assertEquals(t, stocks.NewDollar(13), *eurosToDollars)
}

func TestConversionWithMissingExchangeRate(t *testing.T) {
	initExchangeRates()
	tenEuros := stocks.NewEuro(10)

	eurosToWon, err := bank.Convert(tenEuros, stocks.Krw)

	assertNil(t, eurosToWon)
	assertEquals(t, "exchange not supported: [EUR->KRW]", err.Error())
}

func TestConversionWithInvalidCurrency(t *testing.T) {
	initExchangeRates()
	tenEuros := stocks.NewEuro(10)

	eurosToKalganid, err := bank.Convert(tenEuros, "Kalganid")

	assertNil(t, eurosToKalganid)
	assertEquals(t, "invalid currency: to = [Kalganid]", err.Error())
}
