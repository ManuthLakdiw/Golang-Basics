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

package main

import "fmt"

func generateData(ch chan int) {
	for i := 1; i <= 3; i++ {
		ch <- i // Data යවනවා
	}
	// Data යවලා ඉවරයි. දැන් Channel එක අනිවාර්යයෙන්ම Close කරන්න ඕනේ.
	// නැත්නම් පහළ තියෙන Main ලූප් එක සදාකාලිකව Block වෙලා Deadlock වෙනවා.
	close(ch)
}

func main() {
	ch := make(chan int)

	go generateData(ch)

	// Channel එක Close වෙනකම්ම මේ ලූප් එක වැඩ කරනවා
	for data := range ch {
		fmt.Println("Received:", data)
	}

	fmt.Println("Channel closed. Program finished.")
}
