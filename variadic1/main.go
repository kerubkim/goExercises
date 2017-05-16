package main

import "fmt"

func main() {
	fmt.Println(variadic(1, 25, 3, 4, 5, 6, 7, 8, 9, 10))
}
func variadic(num ...int) int {
	var max int
	for _, n := range num {
		if n > max {
			max = n
		}
	}
	return max
}
