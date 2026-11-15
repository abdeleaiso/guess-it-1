package main

import "math"

func StandardDeviation(va float64) float64 {
	sd := math.Sqrt(va)
	sd = math.Round(sd)
	return sd
}