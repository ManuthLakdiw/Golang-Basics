package main

import (
	"fmt"
	"sync"
	"time"
)

// මෙතනදී අපි WaitGroup එකේ Memory Address එක (Pointer) යවන්න ඕනේ.
// නැත්නම් ඒකේ Copy එකක් හැදිලා Counter එක අවුල් වෙනවා.
func worker(id int, wg *sync.WaitGroup) {
	// Function එක ඉවර වෙලා එළියට යද්දී අනිවාර්යයෙන්ම Counter එකෙන් 1ක් අඩු කරනවා
	defer wg.Done()

	fmt.Printf("Worker %d වැඩ පටන් ගත්තා...\n", id)
	time.Sleep(time.Second) // තත්පරයක් වැඩ කරනවා
	fmt.Printf("Worker %d වැඩ ඉවර කළා!\n", id)
}

func main() {
	var wg sync.WaitGroup // 1. WaitGroup එකක් හදාගන්නවා (Counter එක 0යි)

	// අපි Workers ලා 3 දෙනෙක් යවන්නයි හදන්නේ. ඒ නිසා Counter එක 3ක් කරනවා.
	wg.Add(3)

	// Goroutines 3ක් Start කරනවා.
	// (&wg කියලා Pointer එක යවන්නේ එකම Counter එක හැමෝටම පාවිච්චි කරන්නයි)
	go worker(1, &wg)
	go worker(2, &wg)
	go worker(3, &wg)

	fmt.Println("Main: මම Workers ලා ඉවර වෙනකම් බලන් ඉන්නවා...")

	// 2. Counter එක 0 වෙනකම් Main Thread එක මෙතනින් පහළට යන්නේ නැහැ (Block වෙනවා)
	wg.Wait()

	fmt.Println("Main: ඔක්කොම ඉවරයි, Program එක Close කරනවා!")
}
