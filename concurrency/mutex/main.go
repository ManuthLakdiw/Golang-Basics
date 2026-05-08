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

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var counter int = 0
// var mu sync.Mutex // 1. Mutex එකක් (යතුරක්) හදාගන්නවා

// func incrementSafe(wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	// 2. යතුරෙන් ලොක් කරනවා. දැන් වෙන කාටවත් මෙතනට එන්න බැහැ!
// 	mu.Lock()

// 	counter++ // මේක තමයි Critical Section එක (දැන් 100% Safe)

// 	// 3. යතුර ආපහු දෙනවා. (බොහෝවිට මේකත් defer mu.Unlock() ලෙස ලියයි)
// 	mu.Unlock()
// }

// func main() {
// 	var wg sync.WaitGroup

// 	for i := 0; i < 1000; i++ {
// 		wg.Add(1)
// 		go incrementSafe(&wg)
// 	}

// 	wg.Wait()
// 	// දැන් අනිවාර්යයෙන්ම කිසිම වරදක් නැතුව 1000 ම එනවා!
// 	fmt.Println("Final Counter:", counter)
// }

package main

import (
	"fmt"
	"sync"
)

var once sync.Once // 1. Once object එක හදාගන්නවා

func initializeSystem() {
	fmt.Println("🚀 System එක Initializing... (මේක Print වෙන්නේ එක පාරයි!)")
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %d වැඩට ආවා.\n", id)
	
	// 2. once.Do() ඇතුළේ අපි අර Function එක කතා කරනවා.
	// Goroutines කීයක් මේ පේළියට ආවත්, initializeSystem() එක Run වෙන්නේ මුලින්ම එන කෙනාට විතරයි.
	once.Do(initializeSystem) 
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	go worker(1, &wg)
	go worker(2, &wg)
	go worker(3, &wg)

	wg.Wait()
}