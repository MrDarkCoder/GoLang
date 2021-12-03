package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels")

	myChanel := make(chan int, 2)
	wg := &sync.WaitGroup{}
	// mut := &sync.Mutex{}

	// myChanel <- 5
	// fmt.Println(<-myChanel) deadlock

	wg.Add(2)
	// <- receive only
	go func(ch <-chan int, wg *sync.WaitGroup) {
		value, isChannelOpen := <-myChanel

		fmt.Println(isChannelOpen)
		fmt.Println(value)

		// fmt.Println(<-myChanel)
		// fmt.Println(<-myChanel) // intro to two listeners
		wg.Done()
	}(myChanel, wg)
	// send only <-
	go func(ch chan<- int, wg *sync.WaitGroup) {
		// myChanel <- 5
		// myChanel <- 6 //deadlock
		close(myChanel)
		wg.Done()
	}(myChanel, wg)

	wg.Wait()

}
