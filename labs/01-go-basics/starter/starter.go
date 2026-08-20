package starter

// SumPositive returns the sum of values greater than zero.
func SumPositive(values []int) int {

	var sum int
	for _, v := range values {
		if v > 0 {
			sum += v
		}
	}
	return sum
}
