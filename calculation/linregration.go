package calculation



func LinearRegression(y []float64) (m, c float64) {

	// Calculate x values (indices)
	x := make([]float64, len(y))
	for i := range y {
		x[i] = float64(i)
	}

	// Calculate means
	meanX := average(x)
	meanY := average(y)

	sumXY, sumX2 := 0.0, 0.0

	for i, yi := range y {
		xi := float64(i)
		sumXY += (xi - meanX) * (yi - meanY)
		sumX2 += (xi - meanX) * (xi - meanX)
	}

	// Calculate slope (m) and intercept (c)
	m = sumXY / sumX2
	c = meanY - m*meanX
	return
}