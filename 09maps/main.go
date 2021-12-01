package main

import "fmt"

func main() {
	fmt.Println("welcome")

	languages := make(map[string]string)
	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of Languages: ", languages)
	fmt.Println("JS shorts for: ", languages["JS"])

	// remove
	delete(languages, "RB")
	fmt.Println(languages)

	// loops
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

}
