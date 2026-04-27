package main	

import "fmt"

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


		
}
