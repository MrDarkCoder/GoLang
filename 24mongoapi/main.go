package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MrDarkCoder/mongoapi/router"
)

func main() {
	fmt.Println("Welcome MDB")
	fmt.Println("Server is Getting start... ")

	r := router.Router()
	log.Fatal(http.ListenAndServe(":5000", r))
	fmt.Println("Listening at port 5000")
}
