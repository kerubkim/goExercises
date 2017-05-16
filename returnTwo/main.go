package main

import "fmt"

func main() {
	fmt.Println(divide(42))
	h, even := solution(42)
	fmt.Println("Divided by = ", h)
	fmt.Println("Even or False - ", even)
}

func divide(num int) (int, bool) {
	var b bool
	if num%2 == 0 {
		b = true
	}
	return num / 2, b
}

// Second solution
func solution(num int) (int, bool) {
	return num / 2, num%2 == 0
}
