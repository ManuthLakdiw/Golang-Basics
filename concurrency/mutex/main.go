// without mutex - Data Race

package main

import (
	"fmt"
	"sync"
)

var counter int = 0 // හැමෝම පාවිච්චි කරන පොදු Variable එක

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	// Goroutines 1000ක් එකම වෙලාවේ මේකට කඩාපනිනවා!
	counter++ 
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment(&wg)
	}

	wg.Wait()
	// උත්තරේ 1000 එන්න ඕනේ වුණාට, මේක Run කළාම එන්නේ 945, 980 වගේ අමුතු ගණන්!
	fmt.Println("Final Counter:", counter) 
}
