package main

import "fmt"

func main() {
	fmt.Println("Loops")
	var number int = 10

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// while
	for number > 5 {
		fmt.Println(number)
		number--
	}

	// infinite loop
	// for {
	// 	fmt.Printf("my name is %s and i am %v years old\n", "Manuth", 19)
	// }

	// range
	for i, v := range "Manuth" {
		fmt.Println(i, v)
	}

	for v := range 5 {
		fmt.Println(v)
	}

	arr := []int{1, 2, 3, 4, 5}

	for i, v := range arr {
		fmt.Println(i, v)
	}
}
