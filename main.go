package main

import (
	"fmt"
	"math"
	"os"

	"guess-it/stats"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	fileName := os.Args[1]

	if fileName == "" {
		return
	}
	file := stats.ReadFile(fileName)

	// MedianVale := math.Round(stats.Median(file))
	Mean := math.Round(stats.Average(file))
	StandardDeviation := math.Round(stats.StandardDev(file, Mean))
	Lower, Upper := stats.PredictInterv(Mean, StandardDeviation, float64(len(file)))
	fmt.Println(math.Round(Lower), math.Round(Upper))
}
