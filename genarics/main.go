package main	

import "fmt"

// normal function without generics
func addInt(a int, b int) int {
	return a + b
}
func addFloat(a float64, b float64) float64 {
	return a + b
}

func main() {
	fmt.Println("Sum of 1 and 2:", addInt(1, 2))
	fmt.Println("Sum of 1 and 2:", addFloat(1, 2))

	// cannot do this
	fmt.Println("Sum of 1 and 2:", addInt(1, 2.5))
		
}
