package main

func Average(Data []float64) float64 {
	var sum float64
	for _, v := range Data {
		sum += v
	}
	l := float64(len(Data))
	av := sum / l
	return av
}