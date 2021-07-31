package kata

func Solution(str string) []string {
	var result []string
	if len(str)%2 == 0 {
		for i := 0; i < len(str); i = i + 2 {
			result = append(result, string(str[i])+string(str[i+1]))
		}
	} else {
		for i := 0; i < len(str)-2; i = i + 2 {
			result = append(result, string(str[i])+string(str[i+1]))
		}
		result = append(result, string(str[len(str)-1])+"_")
	}
	return result
}
