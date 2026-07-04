// Package weather provide details of weather forecast depending on
// current location and current condition.
package weather

var (
    // CurrentCondition represents CurrentCondition.
	CurrentCondition string
    // CurrentLocation represents CurrentLocation.
	CurrentLocation  string
)

// Forecast take city and condition as string value and return
// string value after concating theme with weather forecast details.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
