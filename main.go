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

// Bruker init (som anbefalt i dokumentasjonen) for å sikre at flagvariablene
// er initialisert.
func init() {

	/*
	   Her er eksempler på hvordan man implementerer parsing av flagg.
	   For eksempel, kommando
	       funtemps -F 0 -out C
	   skal returnere output: 0°F er -17.78°C
	*/

	// Definerer og initialiserer flagg-variablene
	flag.Float64Var(&fahr, "F", 0.0, "temperatur i grader fahrenheit")
	flag.Float64Var(&C, "C", 0.0, "temperatur i grader celsius")
	flag.Float64Var(&K, "K", 0.0, "temperatur i grader kelvin")

	// Du må selv definere flag-variablene for "C" og "K"
	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")
	flag.StringVar(&funfact, "funfacts", "sun", " om sun - Solen, luna - Månen og terra - Jorden")
	// Du må selv definere flag-variabelen for -t flagget, som bestemmer
	// hvilken temperaturskala skal brukes når funfacts skal vises
	flag.StringVar(&t, "t", "C", "temperaturskala for funfacts: C - celsius, F - farhenheit, K- Kelvin")
}

func main() {

	flag.Parse()

	// Sjekker for funfacts, og printer ut funfacts hvis flagget er satt
	if isFlagPassed("funfacts") {
		if funfact == "sun" {
			fmt.Println(funfacts.Sun(t))
		} else if funfact == "luna" {
			fmt.Println(funfacts.Luna(t))
		} else if funfact == "terra" {
			fmt.Println(funfacts.Terra(t))
		}
		//ellers sjekker for temperaturer og printer ut resultatet av konvertering
	} else {
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
