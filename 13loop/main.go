package main

import "fmt"

func main() {
	fmt.Println("welcome loop")

	// days := []string{"sunday", "tuesday", "wednesday", "friday", "saturday"}
	// fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("Index is : %v, Values is : %v\n", index, day)
	// }

	rougueValue := 1
	for rougueValue < 10 {
		if rougueValue == 2 {
			goto jio
		}
		if rougueValue == 5 {
			rougueValue++
			continue
		}
		fmt.Println("Value is: ", rougueValue)
		rougueValue++
	}

jio:
	fmt.Println("jumping joi")

}
