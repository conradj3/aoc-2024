package shared

// Abs computes the absolute value of an integer.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
