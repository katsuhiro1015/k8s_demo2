package main

import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello Yosistamp")
}

func main() {
	http.HandleFunc("/demo", handler)
	http.ListenAndServe(":8082", nil)
}
