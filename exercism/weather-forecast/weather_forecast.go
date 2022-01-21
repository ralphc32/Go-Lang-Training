// Package weather contains relevant weather details on the location and the condition.
package weather

// CurrentCondition setups current condition of the weather.
var CurrentCondition string

// CurrentLocation setups current location of the said area.
var CurrentLocation string

// Forecast function to retrieve the forecast of the weather in a city location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
