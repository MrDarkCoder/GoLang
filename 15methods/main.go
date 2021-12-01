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

	sridhar.GetStatus()
	sridhar.NewMail()
	fmt.Println(sridhar)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is User Active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@dev.go"
	fmt.Println("Email of this user is: ", u.Email)
}
