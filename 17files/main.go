package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome")
	content := "This needs to go in a file - sridhar.in"

	file, err := os.Create("./mysrifile.txt")

	checkNilError(err)

	length, err := io.WriteString(file, content)
	checkNilError(err)
	fmt.Println("Length is: ", length)

	defer file.Close()

	readFile("./mysrifile.txt")

}

func readFile(filename string) {
	dataByte, err := ioutil.ReadFile(filename)
	checkNilError(err)

	fmt.Println("Text data \n", string(dataByte))
	//  [84 104 105 115 32 110 101 101 100 115 32 116 111 32 103 111 32 105
	// 110 32 97 32 102 105 108 101 32 45 32 115 114 105 100 104 97 114 46 105 110]
}
func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
