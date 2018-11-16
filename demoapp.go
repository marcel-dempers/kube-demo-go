package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func StatusHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "OK: path" + r.URL.Path)
}

func RootHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	fmt.Fprintf(w, "PATH: %v\n", vars["id"])
    fmt.Fprintln(w, "Hello demoapp! Version: V1")
}

func main() {

	fmt.Println("Hello! From Webjet")
	
	r := mux.NewRouter()

	r.HandleFunc("/status", StatusHandler)
	r.HandleFunc("/", RootHandler)

	http.ListenAndServe(":80", r)
}