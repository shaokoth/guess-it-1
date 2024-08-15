package stats

func Average(input []float64) float64 {
	sum := 0.0
	for _, val := range input {
		sum += val
	}
	return sum / float64(len(input))
}
