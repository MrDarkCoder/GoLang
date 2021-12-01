package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome")

	var fruitList = []string{"Apple", "Peach"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)
	fmt.Println("Len", len(fruitList))

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println("Fruit List", fruitList)

	fruitList = fruitList[1:3]
	fmt.Println(fruitList)

	// make
	highScores := make([]int, 4)
	highScores[0] = 243
	highScores[1] = 923
	highScores[2] = 473
	highScores[3] = 783
	// highScores[4] = 555 // ou of bound

	highScores = append(highScores, 55, 123)

	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	// removing a value from slices based on index
	var courses = []string{"ReactJs", "JS", "Swift", "Python", "Ruby"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
