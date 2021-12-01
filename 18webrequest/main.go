package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Welcome")

	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("REsponse Type: %T\n", res)

	defer res.Body.Close() // to close the connection

	dataByte, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	content := string(dataByte)
	fmt.Println(content)

}
