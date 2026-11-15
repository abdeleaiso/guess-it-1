package main

import (
	"math"
)

//E(X)(First Raw Moment)
func Variance(Data []float64, av float64) float64 {
	sqrtMean := math.Pow(av, 2)
	var seg float64
	for i := 0; i < len(Data); i++ {
		seg += math.Pow(Data[i], 2)
	}
	SecondRawMoment := seg / float64(len(Data))
	va := SecondRawMoment - sqrtMean
	va = math.Round(va)
	return va
}