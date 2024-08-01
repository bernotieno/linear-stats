package reader

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"linear-stats/calculation"
)

func InputProcessor(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	var y []float64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number := scanner.Text()
		value, err := strconv.ParseFloat(number, 64)
		if err != nil {
			log.Fatalf("Error parsing float: %v", err)
		}
		y = append(y, value)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Calculate Linear Regression Line and Pearson Correlation Coefficient
	m, c := calculation.LinearRegression(y)
	r := calculation.PearsonCorrelation(y)

	// Print the results
	fmt.Printf("Linear Regression Line: y = %.6fx + %.6f\n", m, c)
	fmt.Printf("Pearson Correlation Coefficient: %.10f\n", r)
}
