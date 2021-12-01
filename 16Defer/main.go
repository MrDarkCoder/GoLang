package main

import "fmt"

func main() {

	defer fmt.Println("Hello world, i'm last")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Welcome")
	fmt.Println("yes i'm second")
	myDefer()

}

// [world, one, two] => LIFO -> two, one, world

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

// [world, one, two, 0, 1, 2, 3, 4] => 4, 3, 2, 1, 0, two, one, world
