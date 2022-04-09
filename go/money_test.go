package main

import (
	"reflect"
	"tdd/stocks"
	"testing"
)

var bank stocks.Bank

func initExchangeRates() {
	bank = stocks.NewBank()
	bank.AddExchangeRate("EUR", "USD", 1.2)
	bank.AddExchangeRate("USD", "KRW", 1100)
}

func TestAddition(t *testing.T) {
	initExchangeRates()
	var portfolio stocks.Portfolio

	fiveDollars := stocks.NewMoney(5, "USD")
	tenDollars := stocks.NewMoney(10, "USD")
	fiftenDollars := stocks.NewMoney(15, "USD")

	portfolio = portfolio.Add(fiveDollars)
	portfolio = portfolio.Add(tenDollars)
	currencyInDollars, err := portfolio.Evaluate(bank, "USD")

	assertEquals(t, fiftenDollars, *currencyInDollars)
	assertNil(t, err)
}

func TestAdditionOfDollarsAndEuros(t *testing.T) {
	initExchangeRates()
	var portfolio stocks.Portfolio

	fiveDollars := stocks.NewMoney(5, "USD")
	tenEuros := stocks.NewMoney(10, "EUR")

	portfolio = portfolio.Add(fiveDollars)
	portfolio = portfolio.Add(tenEuros)

	expectedValue := stocks.NewMoney(17, "USD")
	actualValue, err := portfolio.Evaluate(bank, "USD")

	assertEquals(t, expectedValue, *actualValue)
	assertNil(t, err)
}

func TestAdditionOfDollarsAndWons(t *testing.T) {
	initExchangeRates()
	var portfolio stocks.Portfolio

	oneDollar := stocks.NewMoney(1, "USD")
	elevenHundredWons := stocks.NewMoney(1100, "KRW")

	portfolio = portfolio.Add(oneDollar)
	portfolio = portfolio.Add(elevenHundredWons)

	expectedValue := stocks.NewMoney(2200, "KRW")
	actualValue, err := portfolio.Evaluate(bank, "KRW")

	assertEquals(t, expectedValue, *actualValue)
	assertNil(t, err)
}

func TestAdditionWithMultipleMissingExchangeRates(t *testing.T) {
	initExchangeRates()
	var portfolio stocks.Portfolio

	oneDollar := stocks.NewMoney(1, "USD")
	oneEuro := stocks.NewMoney(1, "EUR")
	oneWon := stocks.NewMoney(1, "KRW")

	portfolio = portfolio.Add(oneDollar)
	portfolio = portfolio.Add(oneEuro)
	portfolio = portfolio.Add(oneWon)

	expectedErrorMessage := "Missing exchange rates: [USD->Kalganid, EUR->Kalganid, KRW->Kalganid]"
	emptyMoney, err := portfolio.Evaluate(bank, "Kalganid")

	assertNil(t, emptyMoney)
	assertEquals(t, expectedErrorMessage, err.Error())
}

func TestAddTwoMoneysInSameCurrency(t *testing.T) {
	fiveEuros := stocks.NewMoney(5, "EUR")
	tenEuros := stocks.NewMoney(10, "EUR")
	expectedValue := stocks.NewMoney(15, "EUR")

	actualValue := fiveEuros.Add(&tenEuros)
	assertEquals(t, expectedValue, *actualValue)

	actualValue = tenEuros.Add(&fiveEuros)
	assertEquals(t, expectedValue, *actualValue)
}

func TestAddTwoMoneysInDiffernetCurrencies(t *testing.T) {
	euro := stocks.NewMoney(1, "EUR")
	dollar := stocks.NewMoney(1, "USD")

	assertNil(t, dollar.Add(&euro))
	assertNil(t, euro.Add(&dollar))
}

func TestConversionWithDifferentRatesBetweenTwoCurrencies(t *testing.T) {
	initExchangeRates()
	tenEuros := stocks.NewMoney(10, "EUR")

	eurosToDollars, err := bank.Convert(tenEuros, "USD")

	assertNil(t, err)
	assertEquals(t, stocks.NewMoney(12, "USD"), *eurosToDollars)

	bank.AddExchangeRate("EUR", "USD", 1.3)

	eurosToDollars, err = bank.Convert(tenEuros, "USD")

	assertNil(t, err)
	assertEquals(t, stocks.NewMoney(13, "USD"), *eurosToDollars)
}

func TestConversionWithMissingExchangeRate(t *testing.T) {
	initExchangeRates()
	tenEuros := stocks.NewMoney(10, "EUR")

	eurosToKalganid, err := bank.Convert(tenEuros, "Kalganid")

	assertNil(t, eurosToKalganid)
	assertEquals(t, "EUR->Kalganid", err.Error())
}

func assertEquals(t *testing.T, expectedResult interface{}, actualResult interface{}) {
	if expectedResult != actualResult {
		t.Errorf("Error: Expected [%+v], Got [%+v]", expectedResult, actualResult)
	}
}

func assertNil(t *testing.T, actualValue interface{}) {
	if actualValue != nil && !reflect.ValueOf(actualValue).IsNil() {
		t.Errorf("Error: Expected to be nil, Got [%+v]", actualValue)
	}
}
