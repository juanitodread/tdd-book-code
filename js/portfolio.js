import Money from './money.js';

class Portfolio {
  constructor() {
    this.moneys = [];
  }

  add(...moneys) {
    this.moneys = this.moneys.concat(moneys);
  }

  evaluate(bank, currency) {
    const failures = [];

    const total = this.moneys.reduce((sum, money) => {
      try {
        const convertedMoney = bank.convert(money, currency);
        return sum.add(convertedMoney);
      } catch (error) {
        failures.push(error.message);
        return sum;
      }
    }, new Money(0, currency));

    if (failures.length) {
      throw new Error(`Missing exchange rates: [${failures.join(', ')}]`);
    }

    return total;
  }
}

export default Portfolio;
