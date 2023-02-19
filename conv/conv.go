package conv

/*
  I denne pakken skal alle konverteringfunksjonene
  implementeres. Bruk engelsk.
    FarhenheitToCelsius
    CelsiusToFahrenheit
    KelvinToFarhenheit
    ...
*/
var i5 = 5

// Konverterer Farhenheit til Celsius
func FahrenheitToCelsius(value float64) float64 {
	// Her skal du implementere funksjonen
	// Du skal ikke formattere float64 i denne funksjonen
	// Gjør formattering i main.go med fmt.Printf eller
	// lignende
	return (value - 32) * (float64(i5) / 9)
}

// De andre konverteringsfunksjonene implementere her
// ...
func FahrenheitToKelvin(value float64) float64 {
	return ((value-32)*(float64(i5)/9) + 273.15)
}

func KelvinToCelsius(value float64) float64 {
	return (value - 273.15)
}

func CelsiusToKelvin(value float64) float64 {
	return (value + 273.15)
}

func KelvinToFahrenheit(value float64) float64 {
	return ((value-273.15)*(9/float64(i5)) + 32)
}

func CelsiusToFahrenheit(value float64) float64 {
	return ((value * (9 / float64(i5))) + 32)
}
