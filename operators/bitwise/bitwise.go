package main

import "fmt"

func main() {
    // 5 in Binary: 0101
    // 3 in Binary: 0011
    a, b := 5, 3

    fmt.Println("--- Bitwise Bit-by-Bit Operations ---")
    fmt.Printf("5 & 3 (AND): %04b (%d)\n", a&b, a&b) // 0001 (1)
    fmt.Printf("5 | 3 (OR):  %04b (%d)\n", a|b, a|b) // 0111 (7)
    fmt.Printf("5 ^ 3 (XOR): %04b (%d)\n", a^b, a^b) // 0110 (6)

    fmt.Println("\n--- Bitwise Shift Operations ---")
    val := 4 // Binary: 0100
    
    leftShift := val << 1 // 1000 (අගය 8 වේ - 2න් ගුණ විය)
    fmt.Printf("4 << 1 (Left Shift):  %04b (%d)\n", leftShift, leftShift)

    rightShift := val >> 1 // 0010 (අගය 2 වේ - 2න් බෙදුණි)
    fmt.Printf("4 >> 1 (Right Shift): %04b (%d)\n", rightShift, rightShift)
}