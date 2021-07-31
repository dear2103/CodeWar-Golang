package kata

import "fmt"

func Greet(name string) string {
	var s string = fmt.Sprintf("Hello, %v how are you doing today?", name)
	return s
}
