from dataclasses import dataclass


@dataclass
class Money:
    amount: float
    currency: str

    def times(self, multiplier: int) -> 'Money':
        return Money(self.amount * multiplier, self.currency)

    def divide(self, divisor: int) -> 'Money':
        return Money(self.amount / divisor, self.currency)

    def __add__(self, money: 'Money') -> 'Money':
        if money and self.currency == money.currency:
            return Money(self.amount + money.amount, self.currency)
        
        return None
