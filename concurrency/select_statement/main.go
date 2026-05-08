package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	time.Sleep(2 * time.Second) // තත්පර 2ක් ගන්නවා
	ch <- "Server 1 රිසල්ට් එක"
}

func server2(ch chan string) {
	time.Sleep(1 * time.Second) // තත්පර 1ක් ගන්නවා (මේක වේගවත්)
	ch <- "Server 2 රිසල්ට් එක"
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go server1(ch1)
	go server2(ch2)

	// මෙතනදී අපි Channels දෙකම දිහා එකපාර බලන් ඉන්නවා
	select {
	case msg1 := <-ch1:
		fmt.Println("ආවේ 1 වෙනි එකෙන්:", msg1)
	case msg2 := <-ch2:
		fmt.Println("ආවේ 2 වෙනි එකෙන්:", msg2) // Server 2 වේගවත් නිසා මේක Print වේවි.
	}
}