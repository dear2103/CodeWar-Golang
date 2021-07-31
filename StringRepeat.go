package kata

func RepeatStr(repetitions int, value string) string {
	s := ""
	for i := 1; i <= repetitions; i++ {
		s += value
	}
	return s
}
