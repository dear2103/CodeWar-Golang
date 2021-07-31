package kata

func RemoveChar(word string) string {
	s := string(word[1 : len(word)-1])
	return s
}
