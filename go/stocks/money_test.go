package stocks_test

import (
	"tdd/stocks"
	"testing"
)

func TestNewDollar(t *testing.T) {
	fourDollars := stocks.NewDollar(4)

	expectedResult := stocks.NewMoney(4, stocks.Usd)

	assertEquals(t, expectedResult, fourDollars)
}

func TestNewEuro(t *testing.T) {
	fourEuros := stocks.NewEuro(4)

	expectedResult := stocks.NewMoney(4, stocks.Eur)

	assertEquals(t, expectedResult, fourEuros)
}

func TestNewWon(t *testing.T) {
	fourWons := stocks.NewWon(4)

	expectedResult := stocks.NewMoney(4, stocks.Krw)

	assertEquals(t, expectedResult, fourWons)
}

func TestMultiplication(t *testing.T) {
	tenDollars := stocks.NewDollar(10.0)

	twentyDollars, err := tenDollars.Times(2)

	assertEquals(t, stocks.NewDollar(20), *twentyDollars)
	assertNil(t, err)
}

func TestMultiplicationByZero(t *testing.T) {
	tenDollars := stocks.NewDollar(10.0)

	zeroDollars, err := tenDollars.Times(0)

	assertEquals(t, stocks.NewDollar(0), *zeroDollars)
	assertNil(t, err)
}

func TestMultiplicationByNegativeNumber(t *testing.T) {
	tenDollars := stocks.NewDollar(10.0)

	invalidMoney, err := tenDollars.Times(-2)

	assertEquals(t, "invalid multiplier: Negative value [-2.00]", err.Error())
	assertNil(t, invalidMoney)
}

func TestDivision(t *testing.T) {
	tenDollars := stocks.NewDollar(10)

	fiveDollars, err := tenDollars.Divide(2)

	assertEquals(t, stocks.NewDollar(5), *fiveDollars)
	assertNil(t, err)
}

func TestDivisionByZero(t *testing.T) {
	tenDollars := stocks.NewDollar(10)

	invalidMoney, err := tenDollars.Divide(0)

	assertEquals(t, "invalid divisor: Zero value [0]", err.Error())
	assertNil(t, invalidMoney)
}

func TestDivisionByNegativeNumber(t *testing.T) {
	tenDollars := stocks.NewDollar(10)

	invalidMoney, err := tenDollars.Divide(-2)

	assertEquals(t, "invalid divisor: Negative value [-2.00]", err.Error())
	assertNil(t, invalidMoney)
}

func TestAdd(t *testing.T) {
	twoDollars := stocks.NewDollar(2)
	fourDollars := stocks.NewDollar(4)

	sixDollars, err := twoDollars.Add(&fourDollars)

	assertEquals(t, stocks.NewDollar(6), *sixDollars)
	assertNil(t, err)
}

func TestAddDifferentCurrencies(t *testing.T) {
	twoDollars := stocks.NewDollar(2)
	fourEuros := stocks.NewEuro(4)

	invalidMoney, err := twoDollars.Add(&fourEuros)

	assertEquals(t, "incompatible currencies: money1[USD] != money2[EUR]", err.Error())
	assertNil(t, invalidMoney)
}
