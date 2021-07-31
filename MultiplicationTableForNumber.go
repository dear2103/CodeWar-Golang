package kata

import (
	"fmt"
)

func MultiTable(number int) string {
	var s string = fmt.Sprintf("1 * %v = %v\n2 * %v = %v\n3 * %v = %v\n4 * %v = %v\n5 * %v = %v\n6 * %v = %v\n7 * %v = %v\n8 * %v = %v\n9 * %v = %v\n10 * %v = %v", number, number*1, number, number*2, number, number*3, number, number*4, number, number*5, number, number*6, number, number*7, number, number*8, number, number*9, number, number*10)

	return s
}
