// // without mutex - Data Race

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var counter int = 0 // හැමෝම පාවිච්චි කරන පොදු Variable එක

// func increment(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	// Goroutines 1000ක් එකම වෙලාවේ මේකට කඩාපනිනවා!
// 	counter++
// }

// func main() {
// 	var wg sync.WaitGroup

// 	for i := 0; i < 1000; i++ {
// 		wg.Add(1)
// 		go increment(&wg)
// 	}

// 	wg.Wait()
// 	// උත්තරේ 1000 එන්න ඕනේ වුණාට, මේක Run කළාම එන්නේ 945, 980 වගේ අමුතු ගණන්!
// 	fmt.Println("Final Counter:", counter)
// }

// using mutex

package main

import (
	"fmt"
	"sync"
)

var counter int = 0
var mu sync.Mutex // 1. Mutex එකක් (යතුරක්) හදාගන්නවා

func incrementSafe(wg *sync.WaitGroup) {
	defer wg.Done()

	// 2. යතුරෙන් ලොක් කරනවා. දැන් වෙන කාටවත් මෙතනට එන්න බැහැ!
	mu.Lock()

	counter++ // මේක තමයි Critical Section එක (දැන් 100% Safe)

	// 3. යතුර ආපහු දෙනවා. (බොහෝවිට මේකත් defer mu.Unlock() ලෙස ලියයි)
	mu.Unlock()
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go incrementSafe(&wg)
	}

	wg.Wait()
	// දැන් අනිවාර්යයෙන්ම කිසිම වරදක් නැතුව 1000 ම එනවා!
	fmt.Println("Final Counter:", counter)
}
