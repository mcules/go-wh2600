package models

import "math"

func FahrenheitToCelsius(in float64) float64 {
	return math.Round((in-32)*5/9*10) / 10
}

func MphToKmh(in float64) float64 {
	return math.Round((in * 1.60934) * 10 / 10)
}
