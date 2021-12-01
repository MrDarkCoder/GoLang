package main

import "fmt"

func main() {
	fmt.Println("welcome")

	loginCount := 55

	var result string
	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "More User"
	} else {
		result = "puffs.. 10 user"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is Even")
	} else {
		fmt.Println("NUmber is Odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is not less than 10")
	}

	// if err != nil{}

}
