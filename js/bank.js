import Money from './money.js';

class Bank {
  constructor() {
    this.exchangeRates = new Map();
  }

  addExchangeRate(currencyFrom, currencyTo, rate) {
    const exchange = this._buildExchangeName(currencyFrom, currencyTo);
    this.exchangeRates.set(exchange, rate);
  }

  convert(money, currency) {
    if (money.currency === currency) {
      return new Money(money.amount, money.currency);
    }

    const exchange = this._buildExchangeName(money.currency, currency);
    const rate = this.exchangeRates.get(exchange);

    if (!rate) {
      throw new Error(exchange);
    }
    
    return new Money(money.amount * rate, currency);
  }

  _buildExchangeName(currencyFrom, currencyTo) {
    return `${currencyFrom}->${currencyTo}`;
  }
}

export default Bank;
