package util

/*
NOTE:i made this code according to section 15
*/
const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
)


func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CAD:
		return true
	}
	return false
}