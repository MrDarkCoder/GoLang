package main

import "fmt"

func main() {
	fmt.Println("Welcome")

	// no inheritance
	// no super or parent

	sridhar := User{"Sridhar", "sri@go.dev", true, 20}
	fmt.Println(sridhar)
	fmt.Printf("Sridhar detail are: %+v\n", sridhar)
	fmt.Printf("Name is: %v\n", sridhar.Name)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
