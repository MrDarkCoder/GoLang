package main

import "fmt"

func main() {
	fmt.Println("Welcome")

	// var ptr *int
	// fmt.Println("value, ", ptr) // nil

	mynumber := 23
	var ptrNumber = &mynumber
	fmt.Println("value ptr, ", ptrNumber)  //address
	fmt.Println("value ptr, ", *ptrNumber) //value

	*ptrNumber = *ptrNumber * 2
	fmt.Println("New Value, ", mynumber) // 46

}
