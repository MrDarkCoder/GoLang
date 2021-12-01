package main

import "fmt"

// public
const LoginToken string = "sdghfjavjhnasd"

func main(){
	var username string = "sridhar"
	fmt.Println(username)
	fmt.Printf("Variable is Of Type: %T \n", username)
	
	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is Of Type: %T \n", isLoggedIn)
	
	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is Of Type: %T \n", smallVal)
	
	var smallFloat float32 = 255.00897654
	fmt.Println(smallFloat)
	fmt.Printf("Variable is Of Type: %T \n", smallFloat)

	// default values and aliases
	var anotherVariable int
	fmt.Println(anotherVariable) // 0
	fmt.Printf("Variable is Of Type: %T \n", anotherVariable)

	// implicit type - lexer will do the type!
	var website = "sridhar.github.io"
	fmt.Println(website)

	// no var style
	numberOfUser := 90000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is Of Type: %T \n", LoginToken)
}
