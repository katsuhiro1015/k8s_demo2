package main

import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello Yosistamp")
}

func main() {
	http.HandleFunc("/hello", handler)
	http.ListenAndServe(":8081", nil)
}
