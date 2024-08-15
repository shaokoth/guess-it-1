package stats

import (
	"math"
)

func StandardDev(input []float64, mean float64) float64 {
	sum := 0.0
	for _, val := range input {
		sum += (val - mean) * (val - mean)
	}
	return math.Sqrt(sum / float64(len(input)))
}
