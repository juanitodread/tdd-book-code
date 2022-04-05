from typing import Tuple
from money import Money


class Bank:
    def __init__(self) -> None:
        self._exchange_rates = {}

    def add_exchange_rate(self, currency_from: str, currency_to: str, rate: float) -> None:
        exchange = self._build_exchange_name(currency_from, currency_to)
        self._exchange_rates[exchange] = rate

    def convert(self, money: Money, currency: str) -> Tuple:
        if money.currency == currency:
            return Money(money.amount, currency), None

        exchange = self._build_exchange_name(money.currency, currency)
        if exchange in self._exchange_rates:
            return (Money(money.amount * self._exchange_rates[exchange], currency), None)

        return (None, exchange)

    def _build_exchange_name(self, currency_from: str, currency_to: str) -> str:
        return f'{currency_from}->{currency_to}'
