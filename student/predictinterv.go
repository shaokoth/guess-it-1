package stats

import (
	"math"
)

func PredictInterv(mean, stdDev, input float64) (float64, float64) {
	zValue := 1.96
	standardErr := float64(stdDev * (math.Sqrt(1 + 1/input)))
	marginErr := float64(zValue * standardErr)
	return (mean - marginErr), (mean + marginErr)
}
