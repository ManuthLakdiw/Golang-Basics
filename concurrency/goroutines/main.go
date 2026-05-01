package main

import (
	"fmt"
	"time"
)

// සාමාන්‍ය Function එකක්
func printNumbers(name string) {
	for i := 1; i <= 3; i++ {
		fmt.Printf("%s: %d\n", name, i)
		time.Sleep(500 * time.Millisecond) // තත්පර බාගයක් නිදාගන්නවා
	}
}

func main() {
	// සාමාන්‍ය විදිහට Call කිරීම (Sequential)
	// මේක ඉවර වෙනකම් ඊළඟ පේළියට යන්නේ නැහැ.
	printNumbers("Normal Task")

	// Goroutines විදිහට Call කිරීම (Concurrent)
	// 'go' දැම්ම ගමන් මේවා Background එකට යනවා,
	// Main function එක ඊළඟ පේළියට ක්ෂණිකවම යනවා.
	go printNumbers("Goroutine A")
	go printNumbers("Goroutine B")

	time.Sleep(2 * time.Second)
}
