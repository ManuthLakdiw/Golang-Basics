package main

import "fmt"

func main() {
	fmt.Printf("slice.go running....\n")

	// decalre and initalizing a slice
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("numbers are: %v\n", numbers)
	fmt.Printf("length is: %d\n", len(numbers))
	fmt.Printf("capacity is: %d\n", cap(numbers))

	// update value
	numbers[0] = 10
	fmt.Printf("numbers are: %v\n", numbers)

	// append
	numbers = append(numbers, 6)
	fmt.Printf("numbers are: %v\n", numbers)
	fmt.Printf("length is: %d\n", len(numbers))
	fmt.Printf("capacity is: %d\n", cap(numbers))

	// remove items using append in slice
	numbers = append(numbers[:0], numbers[1:]...)
	fmt.Printf("numbers are: %v\n", numbers)
	fmt.Printf("length is: %d\n", len(numbers))
	fmt.Printf("capacity is: %d\n", cap(numbers))

	// remove 6 in this slice
	fmt.Printf("-------------before remove 6-------------\n")
	fmt.Printf("numbers are: %v\n", numbers)
	fmt.Printf("length is: %d\n", len(numbers))
	fmt.Printf("capacity is: %d\n", cap(numbers))

	numbers = append(numbers[:4], numbers[6:]...)
	fmt.Printf("-------------after remove 6-------------\n")
	fmt.Printf("numbers are: %v\n", numbers)
	fmt.Printf("length is: %d\n", len(numbers))
	fmt.Printf("capacity is: %d\n", cap(numbers))

	// creating slice using make function
	

}
