package main

import (
	"log"

	"github.com/martinlaursen97/integration-test-exercises/pkg/converter"
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

	grade := converter.NewGrade(converter.DK)
	convertedGrade, _ := grade.Convert(12, converter.US)
	log.Println(convertedGrade)

}
