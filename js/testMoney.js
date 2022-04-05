import assert from 'assert';
import Money from './money.js';
import Portfolio from './portfolio.js';
import Bank from './bank.js';

class MoneyTest {
  setUp() {
    this.bank = new Bank();
    this.bank.addExchangeRate('EUR', 'USD', 1.2);
    this.bank.addExchangeRate('USD', 'KRW', 1100);
  }

  testMultiplication() {
    const tenEuros = new Money(10, 'EUR');
    const twentyEuros = new Money(20, 'EUR');

    assert.deepStrictEqual(tenEuros.times(2), twentyEuros);
  }

  testDivision() {
    const originalMoney = new Money(4002, 'KRW');
    const actualMoneyAfterDivision = originalMoney.divide(4);
    const expectedMoneyAfterDivision = new Money(1000.5, 'KRW');

    assert.deepStrictEqual(actualMoneyAfterDivision, expectedMoneyAfterDivision);
  }

  testAddition() {
    const fiveDollars = new Money(5, 'USD');
    const tenDollars = new Money(10, 'USD');
    const fifteenDollars = new Money(15, 'USD');
    const portfolio = new Portfolio();
    
    portfolio.add(fiveDollars, tenDollars);

    assert.deepStrictEqual(portfolio.evaluate(this.bank, 'USD'), fifteenDollars);
  }

  testAdditionOfDollarsAndEuros() {
    const fiveDollars = new Money(5, 'USD');
    const tenEuros = new Money(10, 'EUR');
    const portfolio = new Portfolio();
    
    portfolio.add(fiveDollars, tenEuros);
    
    const expectedValue = new Money(17, 'USD');
    
    assert.deepStrictEqual(portfolio.evaluate(this.bank, 'USD'), expectedValue);
  }

  testAdditionOfDollarsAndWons() {
    const oneDollar = new Money(1, 'USD');
    const elevenHundredWon = new Money(1100, 'KRW');
    const portfolio = new Portfolio();
    
    portfolio.add(oneDollar, elevenHundredWon);
    
    const expectedValue = new Money(2200, 'KRW');
    
    assert.deepStrictEqual(portfolio.evaluate(this.bank, 'KRW'), expectedValue);
  }

  testAdditionWithMultipleMissingExchangeRates() {
    const oneDollar = new Money(1, 'USD');
    const oneEuro = new Money(1, 'EUR');
    const oneWon = new Money(1, 'KRW');
    const portfolio = new Portfolio();
    
    portfolio.add(oneDollar, oneEuro, oneWon);
    
    const errorMessage = 'Missing exchange rates: [USD->Kalganid, EUR->Kalganid, KRW->Kalganid]';
    const expectedError = new Error(errorMessage);

    assert.throws(() => portfolio.evaluate(this.bank, 'Kalganid'), expectedError);
  }

  testConversionWithDifferentRatesBetweenTwoCurrencies() {
    const tenEuros = new Money(10, 'EUR');
    
    assert.deepStrictEqual(this.bank.convert(tenEuros, 'USD'), new Money(12, 'USD'));
    
    this.bank.addExchangeRate('EUR', 'USD', 1.3);
    
    assert.deepStrictEqual(this.bank.convert(tenEuros, 'USD'), new Money(13, 'USD'));
  }

  testConversionWithMissingExchangeRate() {
    const bank = new Bank();
    const tenEuros = new Money(10, 'EUR');
    const expectedError = new Error('EUR->Kalganid');
    
    assert.throws(() => bank.convert(tenEuros, 'Kalganid'), expectedError);
  }

  testAdditionWithTestDouble() {
    const moneyCount = 10;
    const moneys = [];


    for (let i = 0; i < moneyCount; i++) {
      moneys.push(new Money(Math.random(Number.MAX_SAFE_INTEGER), 'Does Not Matter'));
    }

    const bank = {
      convert: () => {
        return new Money(Math.PI, 'Kalganid');
      }
    };
    
    const arbitraryResult = new Money(moneyCount * Math.PI, 'Kalganid');
    const portfolio = new Portfolio();
    portfolio.add(...moneys);
    
    assert.deepStrictEqual(portfolio.evaluate(bank, 'Kalganid'), arbitraryResult);
  }

  testAddTwoMoneysInSameCurrency() {
    const fiveKalganid = new Money(5, 'Kalganid');
    const tenKalganid = new Money(10, 'Kalganid');
    const fifteenKalganid = new Money(15, 'Kalganid');
    
    assert.deepStrictEqual(fiveKalganid.add(tenKalganid), fifteenKalganid);
    assert.deepStrictEqual(tenKalganid.add(fiveKalganid), fifteenKalganid);
  }

  testAddTwoMoneysInDifferentCurrencies() {
    const euro = new Money(1, 'EUR');
    const dollar = new Money(1, 'USD');

    assert.throws(() => euro.add(dollar), new Error('Cannot add USD to EUR'));
    assert.throws(() => dollar.add(euro), new Error('Cannot add EUR to USD'));
  }

  runAllTests() {
    const testMethods = this._getAllTestMethods();

    testMethods.forEach(m => {
      console.log(`Running: ${m}`);
      const method = Reflect.get(this, m);

      try {
        this.setUp();
        Reflect.apply(method, this, []);
      } catch (error) {
        if (error instanceof assert.AssertionError) {
          console.log(error);
        } else {
          throw error;
        }
      }
    });
  }

  _randomizeTestOrder(testMethods) {
    for (let i = testMethods.length - 1; i > 0; i--) {
      const j = Math.floor(Math.random() * (i + 1));
      [testMethods[i], testMethods[j]] = [testMethods[j], testMethods[i]];
    }
    return testMethods;
  }

  _getAllTestMethods() {
    const moneyPrototype = MoneyTest.prototype;
    const allProps = Object.getOwnPropertyNames(moneyPrototype);
    const testMethods = allProps.filter(property => {
      return typeof moneyPrototype[property] === 'function' && property.startsWith('test');
    });
    return this._randomizeTestOrder(testMethods);
  }
}

new MoneyTest().runAllTests();
