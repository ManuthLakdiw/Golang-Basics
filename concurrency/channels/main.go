package main
import "fmt"

func worker(ch chan string) {
    fmt.Println("Worker: Data යවන්න සූදානම් වෙනවා...")
    ch <- "Database Result" // මෙතනදී Worker නතර වෙනවා Main එකෙන් මේක ගන්නකම්
    fmt.Println("Worker: Data යැව්වා. වැඩේ ඉවරයි.") 
}

func main() {
    ch := make(chan string)

    go worker(ch)

    fmt.Println("Main: මම වෙන වැඩක් කරනවා...")
    
    // Main Goroutine එක මෙතනදී නතර වෙනවා අර Background Worker ගෙන් Data එනකම්
    result := <-ch 
    fmt.Println("Main: මට Data ආවා -", result)
}