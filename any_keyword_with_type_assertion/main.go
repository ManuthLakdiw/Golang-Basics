package main
import "fmt"

// මෙතන data කියන variable එකට ඕනෑම දෙයක් භාරගන්න පුළුවන්
func printSecretBox(data any) {
    fmt.Println("පෙට්ටිය ඇතුළේ තියෙන්නේ:", data)
}

func main() {
    printSecretBox("Hello") // String එකක් යවනවා
    printSecretBox(100)     // int එකක් යවනවා
    printSecretBox(3.14)    // float එකක් යවනවා
}

