package main

import "fmt"

func main() {
	variadicFoo(1, 2)
	fmt.Println("----------")
	variadicFoo(1, 2, 3)
	fmt.Println("----------")
	aSlice := []int{1, 2, 3, 4}
	variadicFoo(aSlice...)
	fmt.Println("----------")
	variadicFoo()
}

func variadicFoo(num ...int) {
	for _, n := range num {
		fmt.Println(n)
	}
}
