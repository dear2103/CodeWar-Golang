package kata

func BoolToWord(word bool) string {
	var s string
	if word == true {
		s = "Yes"
	} else {
		s = "No"
	}
	return s
}
