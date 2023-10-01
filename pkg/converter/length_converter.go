package converter

import (
	"fmt"

	"github.com/martinlaursen97/integration-test-exercises/pkg/util"
)

type LengthType int

const INCH_IN_CM = 2.54

type Length struct {
	value  float64
	system LengthType
}

func (l *Length) Convert() (float64, error) {
	if l.value == 0 {
		return 0, fmt.Errorf("cannot convert zero value")
	}
	var convertedLength float64
	if l.system == METRIC {
		convertedLength = l.value * INCH_IN_CM
	} else {
		convertedLength = l.value / INCH_IN_CM
	}
	return util.RoundToTwoDecimals(convertedLength), nil
}

func NewLength(value float64, system LengthType) *Length {
	return &Length{
		value:  util.RoundToTwoDecimals(value),
		system: system,
	}
}
