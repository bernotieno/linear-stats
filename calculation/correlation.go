package calculation

import "math"

func PearsonCorrelation(y []float64) float64 {
	// Calculate x values (indices)
	x := make([]float64, len(y))
	for i := range y {
		x[i] = float64(i)
	}

	// Calculate means
	meanX := average(x)
	meanY := average(y)

	sumXY, sumX2, sumY2 := 0.0, 0.0, 0.0

	for i, yi := range y {
		xi := float64(i)
		sumXY += (xi - meanX) * (yi - meanY)
		sumX2 += (xi - meanX) * (xi - meanX)
		sumY2 += (yi - meanY) * (yi - meanY)
	}

	// Calculate Pearson correlation coefficient
	numerator := sumXY
	denominator := math.Sqrt(sumX2 * sumY2)

	if denominator == 0 {
		return 0
	}

	return numerator / denominator
}
