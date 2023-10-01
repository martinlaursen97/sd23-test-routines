package converter

import (
	"github.com/martinlaursen97/integration-test-exercises/pkg/util"
)

type LengthType int

const (
	METRIC = iota
	IMPERIAL
)

const INCH_IN_CM = 2.54

type Length struct {
	value  float64
	system LengthType
}

func (l *Length) Convert() float64 {
	if l.value == 0 {
		return 0
	}

	if l.system == METRIC {
		return l.value / INCH_IN_CM
	}
	return util.RoundToTwoDecimals(l.value * INCH_IN_CM)
}

func NewLength(value float64, system LengthType) *Length {
	return &Length{
		value:  util.RoundToTwoDecimals(value),
		system: system,
	}
}
