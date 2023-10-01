package main

import (
	"fmt"
	"math"
)

func main() {
	// length := pkg.NewLength(1, pkg.IMPERIAL)
	// converted_length := length.Convert()
	// log.Println(converted_length)

	// weight := pkg.NewWeight(1, pkg.METRIC)
	// converted_weight := weight.Convert()
	// log.Println(converted_weight)

	// temperature := pkg.NewTemperature(1, pkg.CELCIUS)
	// converted_temperature := temperature.Convert(pkg.FAHRENHEIT)
	// log.Println(converted_temperature)

	// // currency := pkg.NewCurrency(pkg.USD)
	// // targetCurrency := pkg.NewCurrency(pkg.DKK)
	// // convertedCurrency, err := currency.Convert(1, targetCurrency)

	// // if err != nil {
	// // 	log.Println(err.Error())
	// // 	return
	// // }

	// // log.Println(convertedCurrency)

	fmt.Println(RoundToTwoDecimals(0.39370078740157477))

}

func RoundToTwoDecimals(num float64) float64 {
	return math.Round(num*100) / 100
}
