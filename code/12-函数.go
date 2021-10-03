package main

import "fmt"

func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}
func celsiusToFahrenheit(c float64) float64 {
		var f float64
		f=(c*9.0/5.0)+32
		return f
}
func kelvinToFahrenheit(k float64) float64 {
		cel:=kelvinToCelsius(k)
		res:=celsiusToFahrenheit(cel)
		return res
}
func main() {
	test:=233.0
	celsius:=kelvinToCelsius(test)
	fmt.Println(celsius)
	kelvin := 0.0
	fahrenheit:=kelvinToFahrenheit(kelvin)
	fmt.Println(kelvin, "K is ",fahrenheit , "F")

}
