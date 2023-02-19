package main

import (
	"flag"
	"fmt"

	"github.com/thevithach/funtemps/conv"
	"github.com/thevithach/funtemps/funfacts"
)

// Definerer flag-variablene i hoved-"scope"
var fahr float64
var out string
var funfact string
var C float64
var F float64
var K float64
var t string

// Jeg antar at hvis jeg hadde brukt func main() så hadde jeg ikke fått else printene ut?
func init() {

	// Funfacts, sjekker om flagget er satt og henter funskjonen fra funfacts.go og printer ut basert på input av flagg og argumenter av bruker

	if t == "C" || t == "F" && isFlagPassed("funfacts") && funfact == "sun" {
		fmt.Println(funfacts.Sun(t))
	} else if t == "C" || t == "F" && isFlagPassed("funfacts") && funfact == "luna" {
		fmt.Println(funfacts.Luna(t))
	} else if t == "C" || t == "F" && isFlagPassed("funfacts") && funfact == "terra" {
		fmt.Println(funfacts.Terra(t))
	} else {
		fmt.Println("Ugyldig kombinasjon av flagg")
	}

	// beregning fra Fahrenheit til celsius og kelvin
	if out == "C" && isFlagPassed("F") {
		fmt.Printf("%.2f", conv.FahrenheitToCelsius(fahr))
	} else if out == "K" && isFlagPassed("F") {
		fmt.Printf("%.2f", conv.FahrenheitToKelvin(fahr))
	} else if out == "C" && isFlagPassed("K") {
		fmt.Printf("%.2f", conv.KelvinToCelsius(K))
	} else if out == "F" && isFlagPassed("K") {
		fmt.Printf("%.2f", conv.KelvinToFahrenheit(K))
	} else if out == "F" && isFlagPassed("C") {
		fmt.Printf("%.2f", conv.CelsiusToFahrenheit(C))
	} else if out == "K" && isFlagPassed("C") {
		fmt.Printf("%.2f", conv.CelsiusToKelvin(C))
	} else {
		fmt.Println("Ugyldig kombinasjon av flagg")
	}
}

// Eksempel på enkel logikk/

// Funksjonen sjekker om flagget er spesifisert på kommandolinje
// Du trenger ikke å bruke den, men den kan hjelpe med logikken
func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
