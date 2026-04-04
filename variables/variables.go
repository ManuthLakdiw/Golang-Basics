package main

import "fmt"

func main() {

	///////// int ////////
	var firstNumber = 10
	var secondNumber int = 20

	fmt.Println(firstNumber + secondNumber) // 30
	fmt.Println(firstNumber - secondNumber) // -10
	fmt.Println(firstNumber * secondNumber) // 200
	fmt.Println(firstNumber / secondNumber) // 0
	fmt.Println(firstNumber % secondNumber) // 10
	fmt.Printf("Type: %T\n", firstNumber)   // Type: int

	///////// string ////////
	var userName = "manuthlakdiw"
	var password string = "ManuthLakdiw0000"

	fmt.Println(userName)              // manuthlakdiw
	fmt.Println(password)              // ManuthLakdiw0000
	fmt.Printf("Type: %T\n", userName) // Type: string
	fmt.Printf("Type: %T\n", password) // Type: string

	///////// float ////////
	var pi float32 = 3.14
	var radius float64 = 10

	fmt.Printf("Type: %T\n", radius) // Type: float64
	fmt.Printf("Type: %T\n", pi)     // Type: float32

	fmt.Println(pi, radius)

	///////// bool ////////
	var isStudent = true
	var isTeacher = false

	fmt.Printf("Type: %T\n", isStudent) // Type: bool
	fmt.Printf("Type: %T\n", isTeacher) // Type: bool

	fmt.Println(isStudent, isTeacher)

	////// conv. 1

	var subj1, subj2 string = "Maths", "Science"
	var subj3, subj4 string = "Physics", "Chemistry"

	fmt.Println(subj1, subj2, subj3, subj4)

	var fName, lName, age = "Manuth", "Lakdiw", 24

	fmt.Println(fName, lName, age)

	////// conv. 2
	lang1 := "Java"
	lang2 := "Python"
	marks := 90
	isPassed := false

	fmt.Println(lang1, lang2, marks, isPassed)

	// constants

	const name = "Manuth"

	fmt.Println(name)

	// name = "can't change" // variables/variables.go:72:2: cannot assign to name (neither addressable nor a map index expression)

	fmt.Println(name)

}
