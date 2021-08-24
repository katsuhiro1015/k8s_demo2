package main

import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello Yosistamp2")
}

func main() {
	http.HandleFunc("/hello", handler)
	http.ListenAndServe(":80", nil)
}
