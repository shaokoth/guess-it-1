package stats

import "sort"

func Median(input []float64) float64 {
	// Finds the middle value after rearranging in an ascending or descending order
	sort.Float64s(input)
	n := len(input)
	if n%2 == 1 {
		return float64(input[n/2])
	}
	mid1 := input[n/2]
	mid2 := input[(n+1)/2]
	return float64(mid1+mid2) / 2
}