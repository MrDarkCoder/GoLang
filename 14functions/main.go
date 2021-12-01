package main

import "fmt"

func main() {
	fmt.Println("Welcome Function")
	greeter()

	result := adder(3, 5)

	fmt.Println("Result is: ", result)

	proResult, myMsg := proAdder(5, 4, 5, 6, 7)
	fmt.Println("Pro Result: ", proResult)
	fmt.Println("Pro Message: ", myMsg)

}

func adder(valueOne int, valueTwo int) int {
	return valueOne + valueTwo
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, value := range values {
		total += value
	}

	return total, "Hi pro res func"
}

func greeter() {
	fmt.Println("Hello")
}
