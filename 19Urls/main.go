package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghasdk46543"

func main() {
	fmt.Println("Welcome")
	fmt.Println(myurl)

	// parsing
	result, _ := url.Parse(myurl)
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Hostname())
	// fmt.Println(result.Port())
	// fmt.Println(result.Path)
	// fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("Type Of qparams %T\n", qparams)

	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("Params is: ", val)
	}

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/tutcss",
		RawQuery: "user=hitest",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)

}
