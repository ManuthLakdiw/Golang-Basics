package main

import "fmt"

func main() {
	fmt.Println("Functions")

	displayMessage()

		// Parameters / Arguments

	sum := add(5, 10) // 5 සහ 10 තමයි arguments
	fmt.Println("Sum:", sum)     // Output: Sum: 15

	displayMessageWithParam("John") // John තමයි argument


}

func displayMessage() {
	fmt.Println("Hello, welcome to Functions!")
}

func add(x int, y int) int {
	return x + y
}

func displayMessageWithParam(name string) {
	fmt.Println("Hello,", name, "Welcome to Functions!")
}