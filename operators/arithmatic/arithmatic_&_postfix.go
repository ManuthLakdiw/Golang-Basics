package main

import "fmt"

func main() {
    a := 10
    b := 3

    fmt.Println("--- Arithmetic Operations ---")
    fmt.Println("Addition (10 + 3):", a+b)
    fmt.Println("Subtraction (10 - 3):", a-b)
    fmt.Println("Multiplication (10 * 3):", a*b)
    fmt.Println("Division (10 / 3):", a/b)     // Result is 3 (integer division)
    fmt.Println("Modulus (10 % 3):", a%b)      // Remainder is 1

    fmt.Println("\n--- Postfix Increment/Decrement ---")
    num := 5
    num++ // num = num + 1
    fmt.Println("After num++:", num)
    
    num-- // num = num - 1
    fmt.Println("After num--:", num)

    // වැරදි ක්‍රම (මේවා uncomment කළොත් error එකක් එයි):
    // y := num++   // ERROR: syntax error: unexpected ++, expecting comma or )
    // ++num        // ERROR: syntax error: unexpected ++
}