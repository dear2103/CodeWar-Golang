package kata

func QuarterOf(month int) int {
	// your code here
	var quarter int
	switch month {
	case 1, 2, 3:
		quarter = 1
	case 4, 5, 6:
		quarter = 2
	case 7, 8, 9:
		quarter = 3
	case 10, 11, 12:
		quarter = 4
	}
	return quarter
}
