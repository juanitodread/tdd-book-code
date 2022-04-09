package stocks_test

import (
	"tdd/stocks"
	"testing"
)

func TestAddSameCurrency(t *testing.T) {
	var portfolio stocks.Portfolio

	fiveDollars := stocks.NewDollar(5)
	tenDollars := stocks.NewDollar(10)
	fiftenDollars := stocks.NewDollar(15)

	portfolio = portfolio.Add(fiveDollars)
	portfolio = portfolio.Add(tenDollars)
	currencyInDollars, err := portfolio.Evaluate(bank, stocks.Usd)

	assertEquals(t, fiftenDollars, *currencyInDollars)
	assertNil(t, err)
}

func TestAdditionOfDollarsAndEuros(t *testing.T) {
	var portfolio stocks.Portfolio

	fiveDollars := stocks.NewDollar(5)
	tenEuros := stocks.NewEuro(10)

	portfolio = portfolio.Add(fiveDollars)
	portfolio = portfolio.Add(tenEuros)

	expectedValue := stocks.NewDollar(17)
	actualValue, err := portfolio.Evaluate(bank, stocks.Usd)

	assertEquals(t, expectedValue, *actualValue)
	assertNil(t, err)
}

func TestAdditionOfDollarsAndWons(t *testing.T) {
	var portfolio stocks.Portfolio

	oneDollar := stocks.NewDollar(1)
	elevenHundredWons := stocks.NewWon(1100)

	portfolio = portfolio.Add(oneDollar)
	portfolio = portfolio.Add(elevenHundredWons)

	expectedValue := stocks.NewWon(2200)
	actualValue, err := portfolio.Evaluate(bank, stocks.Krw)

	assertEquals(t, expectedValue, *actualValue)
	assertNil(t, err)
}

func TestAdditionWithMultipleExchangeRates(t *testing.T) {
	bank = stocks.NewBank()
	bank.AddExchangeRate("EUR", "USD", 1.2)
	bank.AddExchangeRate("USD", "KRW", 1100)
	bank.AddExchangeRate("KRW", "USD", 1.0/1100)

	var portfolio stocks.Portfolio

	oneDollar := stocks.NewDollar(1)
	oneEuro := stocks.NewEuro(1)
	oneWon := stocks.NewWon(1)

	portfolio = portfolio.Add(oneDollar)
	portfolio = portfolio.Add(oneEuro)
	portfolio = portfolio.Add(oneWon)

	expectedValue := stocks.NewDollar(2.200909090909091)
	actualValue, err := portfolio.Evaluate(bank, stocks.Usd)

	assertEquals(t, expectedValue, *actualValue)
	assertNil(t, err)
}

func TestAdditionWithInvalidCurrency(t *testing.T) {
	var portfolio stocks.Portfolio

	oneDollar := stocks.NewDollar(1)
	oneEuro := stocks.NewEuro(1)
	oneWon := stocks.NewWon(1)

	portfolio = portfolio.Add(oneDollar)
	portfolio = portfolio.Add(oneEuro)
	portfolio = portfolio.Add(oneWon)

	expectedErrorMessage := "invalid currency: [mexican-peso]"
	emptyMoney, err := portfolio.Evaluate(bank, "mexican-peso")

	assertNil(t, emptyMoney)
	assertEquals(t, expectedErrorMessage, err.Error())
}

func TestAdditionWithMultipleMissingExchangeRates(t *testing.T) {
	var portfolio stocks.Portfolio

	oneDollar := stocks.NewDollar(1)
	oneEuro := stocks.NewEuro(1)
	oneWon := stocks.NewWon(1)

	portfolio = portfolio.Add(oneDollar)
	portfolio = portfolio.Add(oneEuro)
	portfolio = portfolio.Add(oneWon)

	expectedErrorMessage := "exchange not supported: [USD->EUR], exchange not supported: [KRW->EUR]"
	emptyMoney, err := portfolio.Evaluate(bank, stocks.Eur)

	assertNil(t, emptyMoney)
	assertEquals(t, expectedErrorMessage, err.Error())
}
