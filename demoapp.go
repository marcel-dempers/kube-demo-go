package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("Hello! From Webjet")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("You have triggered a log from the go API")
		fmt.Fprintln(w, "Hello! I am the go API! V1")
	})

	http.ListenAndServe(":80", nil)
}