package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"Reactjs Bootcamp", 299, "sridhar.go.dev", "abc123", []string{"Web-dev", "JS"}},
		{"MERN Bootcamp", 199, "sridhar.go.dev", "abc1234", []string{"Full-Stack", "JS"}},
		{"MEAN Bootcamp", 399, "sridhar.go.dev", "abc1235", nil},
	}
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)

}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	
        {
                "coursename": "Reactjs Bootcamp",
                "price": 299,
                "website": "sridhar.go.dev",
                "tags": ["Web-dev", "JS"]
        }
	
	`)

	var lcoCourse course
	checkVaild := json.Valid(jsonDataFromWeb)

	if checkVaild {
		fmt.Println("Json Was Valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	// key value pair
	var myData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myData)
	fmt.Printf("%#v\n", myData)

	for k, v := range myData {
		fmt.Printf(" %v : %v, type : %T\n", k, v, v)
	}
}
