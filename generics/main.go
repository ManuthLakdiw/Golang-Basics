package main

import (
	"fmt"
)

// normal function without generics
func addInt(a int, b int) int {
	return a + b
}
func addFloat(a float64, b float64) float64 {
	return a + b
}

// generic function
// T කියන්නේ ඕනෑම Data Type එකක් (any) වෙන්න පුළුවන්.
// items කියන්නේ ඒ T ජාතියෙන් හදපු Slice එකක්.
func PrintSlice[T any](items []T) {
    for _, item := range items {
        fmt.Println(item)
    }
}

// generic structs
// T කියන්නේ ඕනෑම දෙයක් වෙන්න පුළුවන්.
// Box කියන Struct එක ඇතුළේ Content කියන Variable එකට එන්නේ අන්න ඒ T Type එකයි.
type Box[T any] struct {
    Content T
}


// type constaints
// 1. Constraint එකක් විදිහට Interface එකක් හදනවා.
// මෙයින් කියන්නේ: "Number කෙනෙක් වෙන්න නම් අනිවාර්යයෙන්ම int හෝ float64 වෙන්න ඕනේ"
type Number interface {
    int | float64
}

// 2. Generic Function එක ලිවීම
// [T Number] කියන්නේ: "T කියන්නේ අනිවාර්යයෙන්ම 'Number' එකේ තියෙන ජාතියක් වෙන්න ඕනේ!"
func Add[T Number](a T, b T) T {
    return a + b // දැන් Go Compiler එක දන්නවා මේවා අනිවාර්යයෙන්ම එකතු කරන්න පුළුවන් අංක කියලා.
}

func main() {
	fmt.Println("Sum of 1 and 2:", addInt(1, 2))
	fmt.Println("Sum of 1 and 2:", addFloat(1, 2))

	// cannot do this
	// fmt.Println("Sum of 1 and 2:", addInt(1, 2.5))


	// call generic function
	stringSlice := []string{"Hello", "World"}
	PrintSlice(stringSlice) // Output: Hello, World

	intSlice := []int{1, 2, 3}
	PrintSlice(intSlice)    // Output: 1, 2, 3


	// demo with generic strucs
	// 1. int දාන්න පුළුවන් Box එකක් හදනවා
    intBox := Box[int]{Content: 100}
    
    // 2. string දාන්න පුළුවන් Box එකක් හදනවා
    stringBox := Box[string]{Content: "Secret Message"}

    fmt.Println(intBox.Content)    // 100
    fmt.Println(stringBox.Content) // Secret Message

    // වැරදි දෙයක් දැම්මොත්?
    // intBox.Content = "Hello" -> ERROR! Compile වෙන්නෙම නැහැ. (100% Type Safe!)



	// demo type constraints
	fmt.Println(Add(10, 20))       // Output: 30
    fmt.Println(Add(5.5, 2.2))     // Output: 7.7

    // fmt.Println(Add("A", "B"))  // ERROR! String කියන්නේ Number එකක් නෙවෙයි.

		
}
