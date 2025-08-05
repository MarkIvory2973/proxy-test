package mathx

import (
	"math"
	"slices"
)

func Sum(input []float64) float64 {
	var output float64
	for _, value := range input {
		output += value
	}

	return output
}

func Mean(input []float64) float64 {
	return Sum(input) / float64(len(input))
}

func isInt(input float64) bool {
	return input == math.Trunc(input)
}

func Percentile(input []float64, k float64) float64 {
	slices.Sort(input)

	a := float64(len(input)) * k
	if isInt(a) {
		return (input[int(a)-1] + input[int(a)]) / 2
	} else {
		return input[int(math.Trunc(a))]
	}
}

func Std(input []float64, eps float64) float64 {
	m := Mean(input)

	var output []float64
	for _, value := range input {
		powValue := math.Pow(value-m, 2)

		output = append(output, powValue)
	}

	return math.Sqrt(Mean(output) + eps)
}

func Norm(input []float64, eps float64) []float64 {
	min := slices.Min(input)
	m := Std(input, eps)

	var output []float64
	for _, value := range input {
		xValue := (value - min) / m

		output = append(output, xValue)
	}

	return output
}
