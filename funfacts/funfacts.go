package funfacts

import (
	"fmt"

	"github.com/thevithach/funtemps/conv"
)

type FunFacts struct {
	Sun   []string
	Luna  []string
	Terra []string
}

func GetFunFacts(about string, t string) []string {
	switch about {
	case "sun":
		return Sun(t)
	case "luna":
		return Luna(t)
	case "terra":
		return Terra(t)
	default:
		return []string{}
	}
}

func Sun(t string) []string {
	var temperaturYtre float64
	var temperaturIndre float64
	var unit string
	switch t {
	case "C":
		temperaturYtre = conv.KelvinToCelsius(5778)
		temperaturIndre = conv.KelvinToCelsius(14999726.85)
		unit = "°C"
	case "F":
		temperaturYtre = conv.KelvinToFahrenheit(5778)
		temperaturIndre = conv.KelvinToFahrenheit(14999726.85)
		unit = "°F"
	default:
		temperaturYtre = 5778
		temperaturIndre = 14999726.85
		unit = "K"
	}

	return []string{
		fmt.Sprintf("Temperatur på ytre lag av Solen er %.2f%s\n", temperaturYtre, unit),
		fmt.Sprintf("Temperatur i solens kjerne er %.2f%s\n", temperaturIndre, unit),
	}
}

func Luna(t string) []string {
	var temperaturNatt float64
	var temperaturDagen float64
	var unit string
	switch t {
	case "K":
		temperaturNatt = conv.CelsiusToKelvin(-183)
		temperaturDagen = conv.CelsiusToKelvin(106)
		unit = "°K"
	case "F":
		temperaturNatt = conv.CelsiusToFahrenheit(-183)
		temperaturDagen = conv.CelsiusToFahrenheit(106)
		unit = "°F"
	default:
		temperaturNatt = -183
		temperaturDagen = 106
		unit = "C"
	}
	return []string{
		fmt.Sprintf("Temperatur på Månens overflate om natten %.2f%s\n", temperaturNatt, unit),
		fmt.Sprintf("Temperatur på Månens overflate om dagen er %.2f%s\n", temperaturDagen, unit),
	}
}

func Terra(t string) []string {
	var temperaturMaxOverflate float64
	var temperaturMinOverflate float64
	var unit string
	switch t {
	case "K":
		temperaturMaxOverflate = conv.CelsiusToKelvin(56.7)
		temperaturMinOverflate = conv.CelsiusToKelvin(-89.4)
		unit = "°K"
	case "F":
		temperaturMaxOverflate = conv.CelsiusToFahrenheit(56.7)
		temperaturMinOverflate = conv.CelsiusToFahrenheit(-89.4)
		unit = "°F"
	default:
		temperaturMaxOverflate = 56.7
		temperaturMinOverflate = -89.4
		unit = "C"
	}
	return []string{
		fmt.Sprintf("Temperatur på Månens overflate om natten %.2f%s\n", temperaturMaxOverflate, unit),
		fmt.Sprintf("Temperatur på Månens overflate om dagen er %.2f%s\n", temperaturMinOverflate, unit),
	}
}
