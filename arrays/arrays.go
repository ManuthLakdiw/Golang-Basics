package main

import "fmt"

func main() {

	var isPassed [5]bool
	fmt.Printf("isPassed are: %v\n", isPassed)
	fmt.Printf("length is: %d\n", len(isPassed))
	isPassed[0] = true
	isPassed[1] = true
	isPassed[2] = true
	fmt.Printf("isPassed are: %v\n", isPassed)

	
	


	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("values are: %v\n", numbers)

	var colours [5]string
	colours[0] = "Red"
	// colors[1] = "Green"
	// colors[2] = "Blue"
	// colors[3] = "Yellow"
	colours[4] = "Purple"

	fmt.Printf("colors are: %v\n", colours)


	// doesnt know initial lenght

	var languages = [...]string{"Sinhala", "English", "Tamil"}
	fmt.Printf("languages are: %v\n", languages)
	fmt.Printf("length is: %d\n", len(languages))





	//////// multi diamentional arrays
	var newArray [4][2]int
	fmt.Printf("newArray is: %v\n", newArray)

	newArray[0][0] = 1
	newArray[0][1] = 2
	newArray[1][0] = 3
	newArray[1][1] = 4
	newArray[2][0] = 5
	newArray[2][1] = 6
	newArray[3][0] = 7
	newArray[3][1] = 8
	fmt.Printf("newArray is: %v\n", newArray)
	

}
