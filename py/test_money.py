import unittest

from money import Money
from portfolio import Portfolio
from bank import Bank


class TestMoney(unittest.TestCase):
    def setUp(self):
        self.bank = Bank()
        self.bank.add_exchange_rate('EUR', 'USD', 1.2)
        self.bank.add_exchange_rate('USD', 'KRW', 1100)

    def test_multiplication(self):
        ten_euros = Money(10, 'EUR')
        twenty_euros = Money(20, 'EUR')

        self.assertEqual(twenty_euros, ten_euros.times(2))

    def test_division(self):
        original_money = Money(4002, 'KRW')
        actual_money_after_division = original_money.divide(4)
        expected_money_after_division = Money(1000.5, 'KRW')
        
        self.assertEqual(expected_money_after_division, actual_money_after_division)

    def test_addition(self):
        five_dollars = Money(5, 'USD')
        ten_dollars = Money(10, 'USD')
        fifteen_dollars = Money(15, 'USD')
        
        portfolio = Portfolio()
        portfolio.add(five_dollars, ten_dollars)

        self.assertEqual(fifteen_dollars, portfolio.evaluate(self.bank, 'USD'))

    def test_addition_of_dollars_and_euros(self):
        five_dollars = Money(5, 'USD')
        ten_euros = Money(10, 'EUR')
        
        portfolio = Portfolio()
        portfolio.add(five_dollars, ten_euros)
        
        expected_value = Money(17, 'USD')
        actual_value = portfolio.evaluate(self.bank, 'USD')
        
        self.assertEqual(expected_value, actual_value)

    def test_addition_of_dollars_and_wons(self):
        one_dollar = Money(1, 'USD')
        eleven_hundred_won = Money(1100, 'KRW')
        
        portfolio = Portfolio()
        portfolio.add(one_dollar, eleven_hundred_won)
        
        expected_value = Money(2200, 'KRW')
        actual_value = portfolio.evaluate(self.bank, 'KRW')

        self.assertEqual(expected_value, actual_value)


    def test_addition_with_multiple_missing_exchange_rates(self):
        one_dollar = Money(1, 'USD')
        one_euro = Money(1, 'EUR')
        one_won = Money(1, 'KRW')
        
        portfolio = Portfolio()
        portfolio.add(one_dollar, one_euro, one_won)

        with self.assertRaises(Exception) as error:
            portfolio.evaluate(self.bank, 'Kalganid')

        self.assertEqual(
            'Missing exchange rates: [USD->Kalganid, EUR->Kalganid, KRW->Kalganid]', 
            str(error.exception)
        )

    def test_conversion_with_different_rates_between_two_currencies(self):
        ten_euros = Money(10, 'EUR')
        result, missing_exchange = self.bank.convert(ten_euros, 'USD')
        
        self.assertEqual(result, Money(12, 'USD'))
        self.assertIsNone(missing_exchange)

        self.bank.add_exchange_rate('EUR', 'USD', 1.3)
        result, missing_exchange = self.bank.convert(ten_euros, 'USD')
        
        self.assertEqual(result, Money(13, 'USD'))
        self.assertIsNone(missing_exchange)

    def test_conversion_with_missing_exchange_rate(self):
        bank = Bank()
        ten_euros = Money(10, 'EUR')
        result, missing_exchange = self.bank.convert(ten_euros, 'Kalganid')
        
        self.assertIsNone(result)
        self.assertEqual(missing_exchange, 'EUR->Kalganid')

    def test_add_moneys_directly(self):
        self.assertEqual(Money(15, 'USD'), Money(5, 'USD') + Money(10, 'USD'))
        self.assertEqual(Money(15, 'USD'), Money(10, 'USD') + Money(5, 'USD'))
        self.assertEqual(None, Money(5, 'USD') + Money(10, 'EUR'))
        self.assertEqual(None, Money(5, 'USD') + None)


if __name__ == '__main__':
    unittest.main()
