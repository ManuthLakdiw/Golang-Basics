package main

import "fmt"

func main() {
	var num int = 10
	fmt.Printf("number is: %d\n", num)

	var ptr *int

	ptr = &num

	fmt.Printf("address of num is: %d\n", &num)

	fmt.Printf("address of ptr is: %d\n", ptr)

	fmt.Printf("de-referencing: %d\n", *ptr)

	// change value using pointer

	*ptr = 20

	fmt.Printf("number is: %d\n", num)


	// pointers memory address
	fmt.Printf("address of ptr is: %d\n", &ptr)

}
