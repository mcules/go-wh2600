package models

import "math"

func FahrenheitToCelsius(in float32) float32 {
	return float32(math.Round((float64(in)-32)*5/9*10) / 10)
}

func MphToKmh(in float32) float32 {
	return float32(math.Round((float64(in) * 1.60934) * 10 / 10))
}
