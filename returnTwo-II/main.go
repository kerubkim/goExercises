package main

import "fmt"

func main() {
	x := func(num int) (int, bool) {
		if num%2 == 0 {
			return num / 2, true
		} else {
			return num / 2, false
		}
	}
	fmt.Println(x(42))
}
