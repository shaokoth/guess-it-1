package stats

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Opens and reads the file input.
func ReadFile(fileName string) []float64 {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("error opening file")
	}
	defer file.Close()
	var result []float64

	// Scans the file input line by line
	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		count++
		values := strings.Split(line, "\n")
		if line == "" {
			continue
		}
		// converts the string to float64
		number, err := strconv.ParseFloat(values[0.0], 64)
		if err != nil {
			fmt.Println("Error parsing data")
			return nil
		} else {
			result = append(result, number)
		}
		
	}
	if len(result) == 0 {
		fmt.Println("File is empty")
		os.Exit(1)
	}
	for scanner.Scan() {
		if err != nil {
			fmt.Println("err")
		} else {
			return nil
		}
	}
	return result
}
