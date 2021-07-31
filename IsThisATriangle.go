package kata

func IsTriangle(a, b, c int) bool {
	if a+b > c && a+c > b && b+c > a {
		return true
	} else {
		return false
	}
}
