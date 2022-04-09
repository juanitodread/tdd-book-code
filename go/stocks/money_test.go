package stocks_test

import (
	"reflect"
	"tdd/stocks"
	"testing"
)

func TestNewMoneyDefault(t *testing.T) {
	money := stocks.NewMoney(0.0, "")

	expectedResult := stocks.Money{}

	assertEquals(t, expectedResult, money)
}

func TestMultiplication(t *testing.T) {
	tenDollars := stocks.NewMoney(10.0, "USD")

	twentyDollars, err := tenDollars.Times(2)

	assertEquals(t, stocks.NewMoney(20, "USD"), *twentyDollars)
	assertNil(t, err)
}

func TestMultiplicationByZero(t *testing.T) {
	tenDollars := stocks.NewMoney(10.0, "USD")

	zeroDollars, err := tenDollars.Times(0)

	assertEquals(t, stocks.NewMoney(0, "USD"), *zeroDollars)
	assertNil(t, err)
}

func TestMultiplicationByNegative(t *testing.T) {
	tenDollars := stocks.NewMoney(10.0, "USD")

	invalidMoney, err := tenDollars.Times(-2)

	assertEquals(t, "invalid multiplier: Negative value [-2.00]", err.Error())
	assertNil(t, invalidMoney)
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
