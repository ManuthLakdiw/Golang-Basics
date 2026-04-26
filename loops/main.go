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

	
}
