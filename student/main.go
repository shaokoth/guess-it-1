package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"

	"guess-it/stats"
)

func main() {
	if len(os.Args) != 1 {
		fmt.Println("Usage should be: go run .")
		return
	}

	values := []float64{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		data := scanner.Text()
		num, err := strconv.ParseFloat(data, 64)
		if err != nil {
			fmt.Println("Error parsing data")
			return
		} else {
			values = append(values, num)
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		Mean := stats.Average(values)
		StandardDeviation := stats.StandardDev(values, Mean)
		Lower, Upper := stats.PredictInterv(Mean, StandardDeviation, values)
		fmt.Printf("%d %d\n", int(math.Round(Lower)), int(math.Round(Upper)))
	}

}
