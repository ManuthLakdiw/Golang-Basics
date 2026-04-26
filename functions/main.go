package main

import "fmt"

func main() {
	fmt.Println("Functions")

	displayMessage()

	// Parameters / Arguments

	sum := add(5, 10)        // 5 සහ 10 තමයි arguments
	fmt.Println("Sum:", sum) // Output: Sum: 15

	displayMessageWithParam("John") // John තමයි argument

	// Multiple return Values
	a, b := swap("Hello", "World")
	fmt.Println("Swapped:", a, b) // Output: Swapped: World Hello

	// Variadic Functions
	fmt.Println("Sum of 1, 2, 3:", sumAll(1, 2, 3)) // Output: 6
	fmt.Println("Sum of 10, 20:", sumAll(10, 20))   // Output: 30

	// Anonymous Function
	addAnonymous := func(x int, y int) int {
		return x + y
	}
	fmt.Println("Anonymous Sum:", addAnonymous(4, 5)) // Output: 9

	// Closures
	counter := createCounter()
	fmt.Println("Counter:", counter()) // Output: 1
	fmt.Println("Counter:", counter()) // Output: 2
	fmt.Println("Counter:", counter()) // Output: 3

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

//// Multiple return Values

func swap(x string, y string) (string, string) {
	return y, x
}

// variadic functions

func sumAll(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}


// Closures

func createCounter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

