package fil

// test
func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
		return true
	} else {
		return false
	}
}

func ChooseVehicle(option1, option2 string) string {
	if option1 < option2 {
		return option1 + " is clearly the better choice."
	} else {
		return option2 + " is clearly the better choice."
	}
}

func CalculateResellPrice(originalPrice, age float64) float64 {
	if age < 3 {
		originalPrice = originalPrice * 0.8
	} else if age < 10 {
		originalPrice = originalPrice * 0.7
	} else {
		originalPrice = originalPrice * 0.5
	}
	return originalPrice
}
