package stocks_test

import (
	"tdd/stocks"
	"testing"
)

var bank stocks.Bank

func initExchangeRates() {
	exchangeRateEurToUsd, _ := stocks.NewExchangeRate(stocks.Eur, stocks.Usd, 1.2)
	exchangeRateUsdToKrw, _ := stocks.NewExchangeRate(stocks.Usd, stocks.Krw, 1100)

	bank = stocks.NewBank()
	bank.AddExchangeRate(*exchangeRateEurToUsd)
	bank.AddExchangeRate(*exchangeRateUsdToKrw)
}

func TestNewBank(t *testing.T) {
	bank := stocks.NewBank()

	assertNotNil(t, bank)
}

func TestConversionWithDifferentRatesBetweenTwoCurrencies(t *testing.T) {
	initExchangeRates()
	tenEuros := stocks.NewEuro(10)

	eurosToDollars, err := bank.Convert(tenEuros, stocks.Usd)

	assertNil(t, err)
	assertEquals(t, stocks.NewDollar(12), *eurosToDollars)

	exchangeRate, _ := stocks.NewExchangeRate(stocks.Eur, stocks.Usd, 1.3)
	bank.AddExchangeRate(*exchangeRate)

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
