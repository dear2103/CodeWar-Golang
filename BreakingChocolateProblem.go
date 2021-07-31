package kata

func BreakChocolate(n, m int) int {
	// math goes here
	if m > 0 && n > 0 {
		return m*n - 1
	} else {
		return 0
	}
}
