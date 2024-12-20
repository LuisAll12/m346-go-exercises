package main
import(
	"fmt"
)
// TODO: implement the function convertCelsiusToFahrenheit
func convertCelsiusToFahrenheit(celsius float64) float64 {
	var fahrenheit float64 = celsius * 9 / 5 + 32
	return fahrenheit
}
// TODO: implement the function convertFahrenheitToCelsius
func convertFahrenheitToCelsius(fahrenheit float64) float64 {
	var celsius float64 = (fahrenheit - 32) * 5 / 9
	return celsius
}
func main() {
	// TODO: call the function convertCelsiusToFahrenheit
	var resultFahrenheit float64 
	resultFahrenheit = convertCelsiusToFahrenheit(28)
	fmt.Println("20°C in Fahrenheit: ", resultFahrenheit)
	// TODO: call the function convertFahrenheitToCelsius
	var resultCelsius float64  
	resultCelsius = convertFahrenheitToCelsius(62)
	fmt.Println("68°F in Celsius: ", resultCelsius)
}
