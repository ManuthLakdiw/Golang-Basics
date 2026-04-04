package main

import "fmt"

func main() {
    x := 10
    fmt.Println("Initial x:", x)

    x += 5 // x = 10 + 5
    fmt.Printf("After x += 5: %d\n", x)

    x *= 2 // x = 15 * 2
    fmt.Printf("After x *= 2: %d\n", x)

    // Bitwise Compound Assignment
    // 10 (Binary: 1010) AND 3 (Binary: 0011) = 2 (Binary: 0010)
    y := 10
    y &= 3 
    fmt.Printf("After y &= 3 (Bitwise AND assignment): %d\n", y)
}