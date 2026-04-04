package main

import "fmt"

// Function to increment value using pointer
func increment(ptr *int) {
	*ptr++ 
}


// Function to set a value to zero using pointer
func setZero(ptr *int) {
	*ptr = 0
}

func demo() {
	fmt.Printf("demo func invoked!\n")

}
