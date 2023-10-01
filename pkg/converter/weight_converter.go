package converter

import "github.com/martinlaursen97/integration-test-exercises/pkg/util"

type WeightType int

const (
	METRIC = iota
	IMPERIAL
)

const LB_IN_KG = 0.453

type Weight struct {
	value  float64
	system WeightType
}

func (w *Weight) Convert() float64 {
	var convertedWeight float64
	if w.system == METRIC {
		convertedWeight = w.value / LB_IN_KG
	} else {
		convertedWeight = w.value * LB_IN_KG
	}
	return util.RoundToTwoDecimals(convertedWeight)
}

func NewWeight(value float64, system WeightType) *Weight {
	return &Weight{
		value:  util.RoundToTwoDecimals(value),
		system: system,
	}
}
