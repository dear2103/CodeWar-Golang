package kata

import (
	"strconv"
	"strings"
)

func Is_valid_ip(ip string) bool {
	ip4 := strings.Split(ip, ".")
	var valid bool
	if len(ip4) == 4 {
		for i := 0; i < len(ip4); i++ {
			a, err := strconv.Atoi(string(ip4[i]))
			if a < 0 || a > 255 || err != nil || string(ip4[i]) != strconv.Itoa(a) {
				valid = false
				break
			} else {
				valid = true
			}
		}
	} else {
		valid = false
	}
	return valid
}
