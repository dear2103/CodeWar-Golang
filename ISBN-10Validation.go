package kata

import (
	"strconv"
)

func ValidISBN10(isbn string) bool {
	var valid bool
	if len(isbn) == 10 {
		_, err := strconv.Atoi(isbn[:9])
		if err == nil {

			if string(isbn[9]) == "X" || string(isbn[9]) == "x" {
				sum := 100
				for i, j := 0, 1; i < len(isbn)-1 && j < 10; i, j = i+1, j+1 {
					a, _ := strconv.Atoi(string(isbn[i]))
					sum = sum + a*j
				}
				if sum%11 == 0 {
					valid = true
				} else {
					valid = false
				}
			} else {
				sum := 0
				for i, j := 0, 1; i < len(isbn) && j < 11; i, j = i+1, j+1 {
					a, _ := strconv.Atoi(string(isbn[i]))
					sum += a * j
				}
				if sum%11 == 0 {
					valid = true
				} else {
					valid = false
				}
			}
		} else {
			valid = false
		}
	} else {
		valid = false
	}
	return valid
}
