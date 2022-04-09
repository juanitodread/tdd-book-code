package stocks

type Currency string

const (
	Usd Currency = "USD"
	Eur Currency = "EUR"
	Krw Currency = "KRW"
)

func (currency Currency) exists() bool {
	switch currency {
	case Usd:
		return true
	case Eur:
		return true
	case Krw:
		return true
	case "": //@TODO: Empty value shouldnt be allowed
		return true
	default:
		return false
	}
}
