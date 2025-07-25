package main

import (
	"fmt"
	"net/http"
)

func mainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello")
}

func main() {
	storage :=
	err := http.ListenAndServe("localhost:8080", http.HandlerFunc(mainPage))
	if err != nil {
		panic(err)
	}
}
