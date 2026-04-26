package main

import "fmt"

func main() {
	fmt.Println("Control Flow")

	age := 22

    if age >= 18 {
        fmt.Println("You are an adult.")
    } else if age >= 13 {
        fmt.Println("You are a teenager.")
    } else {
        fmt.Println("You are a child.")
    }
}

