package main

import "fmt"

func main() {
	var person string
	fmt.Print("Enter a name : ") // ask client for a number
	fmt.Scan(&person)               // save the input number into the variable "num"
	fmt.Println("Hello ",person)
}
