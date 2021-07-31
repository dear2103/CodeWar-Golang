package kata

import (
	"fmt"
	"strings"
)

func AbbrevName(name string) string {
	//your code here
	s := strings.Split(name, " ")
	a1 := string(name[0])
	a2 := string(name[len(name)-len(s[1])])
	sf := fmt.Sprintf("%v.%v", strings.ToUpper(a1), strings.ToUpper(a2))
	return sf
}
