package stats

import (
	"math"
)

func PredictInterv(mean, stdDev float64, input []float64) (float64, float64) {
	zValue := 1.96
	standardErr := float64(stdDev * (math.Sqrt(1 + 1/float64(len(input)))))
	marginErr := float64(zValue * standardErr)
	return (mean - marginErr), (mean + marginErr)
}
