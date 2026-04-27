package main	

import "fmt"

// normal function without generics
func addInt(a int, b int) int {
	return a + b
}
func addFloat(a float64, b float64) float64 {
	return a + b
}

// generic function
// T කියන්නේ ඕනෑම Data Type එකක් (any) වෙන්න පුළුවන්.
// items කියන්නේ ඒ T ජාතියෙන් හදපු Slice එකක්.
func PrintSlice[T any](items []T) {
    for _, item := range items {
        fmt.Println(item)
    }
}

func main() {
	fmt.Println("Sum of 1 and 2:", addInt(1, 2))
	fmt.Println("Sum of 1 and 2:", addFloat(1, 2))

	// cannot do this
	// fmt.Println("Sum of 1 and 2:", addInt(1, 2.5))


	// call generic function
	stringSlice := []string{"Hello", "World"}
	PrintSlice(stringSlice) // Output: Hello, World

	intSlice := []int{1, 2, 3}
	PrintSlice(intSlice)    // Output: 1, 2, 3


		
}
