package main

import (
	"fmt"
	"sync"
)

func main() {
	// Sync
	fmt.Println("do struff")
	wg := sync.WaitGroup{}
	wg.Add(1)
	go doStruff(&wg)
	fmt.Println("THanks")
	wg.Wait()

	// Channels
	fmt.Println("do struff")
	c := make(chan bool)
	go doStruff2(c)
	<-c
	wg.Wait()
	fmt.Println("THanks")

}

func doStruff(wg *sync.WaitGroup) {
	fmt.Println("Hello. I do stuff")
	wg.Done()
}
func doStruff2(c chan bool) {
	c <- true
	fmt.Println("Hello. I do stuff")
}
