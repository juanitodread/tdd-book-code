from money import Money
from bank import Bank


class Portfolio:
    def __init__(self) -> None:
        self._moneys = []

    def add(self, *moneys: list) -> None:
        self._moneys.extend(moneys)

    def evaluate(self, bank: Bank, currency: str) -> Money:
        total = Money(0, currency)
        failures = []

        for money in self._moneys:
            converted_currency, error = bank.convert(money, currency)
            if not error:
                total += converted_currency
            else:
                failures.append(error)

        if failures:
            raise Exception(f'Missing exchange rates: [{", ".join(failures)}]')

        return total
