// Package weather provides functions
// to display weather condition from given location.
package weather

// CurrentCondition represents the weather condition to display.
var CurrentCondition string
// CurrentLocation represents the current location of the weather condition to display.
var CurrentLocation string

// Forecast returns a string containing the weather condition of the current location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
