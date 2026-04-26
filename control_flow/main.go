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


	// limit කියන variable එක හදන්නෙත්, condition එක check කරන්නෙත් එකම පේළියේ (;) මඟින් වෙන් කර ඇත.
    if limit := 100; limit > 50 {
        fmt.Println("Limit is high:", limit)
    } else {
        fmt.Println("Limit is fine:", limit)
    }

    // fmt.Println(limit) -> මෙහෙම ලිව්වොත් Error එකක් එනවා! 
    // මොකද limit කියන variable එක if එකෙන් එළියේ නැහැ.
	
}

