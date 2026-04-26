package main

import "fmt"

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

func main() {
	fmt.Println("Structs......")

	manuth := Person{
		FirstName: "Manuth",
		LastName:  "Lakdiw",
		Age:       20,
	}

	fmt.Println(manuth)
	fmt.Println("First Name: ", manuth.FirstName)
	fmt.Println("Last Name: ", manuth.LastName)
	fmt.Println("Age: ", manuth.Age)

	job := struct {
		JobTitle string
		Salary   int
	}{
		JobTitle: "Software Engineer",
		Salary:   100000,
	}

	fmt.Println(job)

}
