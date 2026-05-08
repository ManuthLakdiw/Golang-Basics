// package main
// import "fmt"

// func worker(ch chan string) {
//     fmt.Println("Worker: Data යවන්න සූදානම් වෙනවා...")
//     ch <- "Database Result" // මෙතනදී Worker නතර වෙනවා Main එකෙන් මේක ගන්නකම්
//     fmt.Println("Worker: Data යැව්වා. වැඩේ ඉවරයි.")
// }

// func main() {
//     ch := make(chan string)

//     go worker(ch)

//     fmt.Println("Main: මම වෙන වැඩක් කරනවා...")

//     // Main Goroutine එක මෙතනදී නතර වෙනවා අර Background Worker ගෙන් Data එනකම්
//     result := <-ch
//     fmt.Println("Main: මට Data ආවා -", result)
// }

// channel with for loop

// package main

// import "fmt"

// func generateData(ch chan int) {
// 	for i := 1; i <= 3; i++ {
// 		ch <- i // Data යවනවා
// 	}
// 	// Data යවලා ඉවරයි. දැන් Channel එක අනිවාර්යයෙන්ම Close කරන්න ඕනේ.
// 	// නැත්නම් පහළ තියෙන Main ලූප් එක සදාකාලිකව Block වෙලා Deadlock වෙනවා.
// 	close(ch)
// }

// func main() {
// 	ch := make(chan int)

// 	go generateData(ch)

// 	// Channel එක Close වෙනකම්ම මේ ලූප් එක වැඩ කරනවා
// 	for data := range ch {
// 		fmt.Println("Received:", data)
// 	}

// 	fmt.Println("Channel closed. Program finished.")
// }

// buffered channel
package main

import "fmt"

func main() {
	// Capacity 3ක් තියෙන Buffered Channel එකක් හදනවා
	ch := make(chan string, 3)

	// මේවා යවද්දී Receiver කෙනෙක් බලන් ඉන්න ඕනේ නෑ! (Block වෙන්නේ නෑ)
	ch <- "Apple"
	fmt.Println("Apple යැව්වා")

	ch <- "Banana"
	fmt.Println("Banana යැව්වා")

	ch <- "Mango"
	fmt.Println("Mango යැව්වා")

	// 🔴 අනතුරක්:
	// දැන් Capacity එක (3/3) පිරිලා තියෙන්නේ.
	// මම මෙතන ch <- "Orange" කියලා දැම්මොත්, මේ Main Goroutine එක එතනම Block වෙලා Deadlock වෙනවා!
	// (මොකද ගන්න කවුරුත් නෑ, දාන්න ඉඩකුත් නෑ).

	// Data එළියට ගැනීම
	fmt.Println("ගත්තා:", <-ch) // Apple එළියට එනවා. දැන් ඉඩක් හැදුණා (2/3).
	fmt.Println("ගත්තා:", <-ch) // Banana
	fmt.Println("ගත්තා:", <-ch) // Mango
}
