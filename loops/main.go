package main

import "fmt"

func main() {
	fmt.Println("Loops")
	var number int = 10;

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}


	// while
	for number > 5 {
		fmt.Println(number)
		number--
	}

	// infinite loop
	for {
		fmt.Printf("my name is %s and i am %v years old\n", "Manuth", 19)
	}
}
