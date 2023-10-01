package converter

import (
	"fmt"

	"github.com/martinlaursen97/integration-test-exercises/pkg/util"
)

type TemperatureType int

const (
	CELCIUS = iota
	FAHRENHEIT
	KELVIN
)

const (
	ONE_CELCIUS       = 1
	CELCIUS_TO_KELVIN = 273.15
)

type Temperature struct {
	value  float64
	system TemperatureType
}

func (t *Temperature) Convert(target TemperatureType) (float64, error) {
	if t.system == target {
		return 0, fmt.Errorf("cannot convert between the same systems")
	}

	var convertedTemperature float64

	switch t.system {
	case CELCIUS:
		if target == FAHRENHEIT {
			convertedTemperature = (t.value * 9 / 5) + 32
		} else if target == KELVIN {
			convertedTemperature = t.value + CELCIUS_TO_KELVIN
		}
	case FAHRENHEIT:
		if target == CELCIUS {
			convertedTemperature = (t.value - 32) * 5 / 9
		} else if target == KELVIN {
			convertedTemperature = (t.value-32)*5/9 + CELCIUS_TO_KELVIN
		}
	case KELVIN:
		if target == CELCIUS {
			convertedTemperature = t.value - CELCIUS_TO_KELVIN
		} else if target == FAHRENHEIT {
			convertedTemperature = (t.value-CELCIUS_TO_KELVIN)*9/5 + 32
		}
	default:
		return 0, fmt.Errorf("invalid temperature system")
	}
	return util.RoundToTwoDecimals(convertedTemperature), nil
}

func NewTemperature(value float64, system TemperatureType) *Temperature {
	return &Temperature{
		value:  value,
		system: system,
	}
}
