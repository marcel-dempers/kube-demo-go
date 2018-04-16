package main

import (
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func main() {

	fmt.Println("Hello! From Webjet")
	
	router := httprouter.New()

	router.GET("/status", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
		fmt.Fprintln(w, "OK")
	})

	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
		fmt.Fprintln(w, "Hello demoapp! Version: V1")
	})

	http.ListenAndServe(":80", router)
}