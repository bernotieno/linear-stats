package calculation

func average(nums []float64) float64 {
	total := 0.0
	lenNum := len(nums)
	for _, val := range nums {
		total += val
	}
	mean := total / float64(lenNum)
	return mean
}
