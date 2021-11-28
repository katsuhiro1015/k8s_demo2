package main

import (
	"net/http"

	"k8s_demo2/output"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	output.CreditLimit()
}

func main() {
	http.HandleFunc("/demo", handler)
	http.ListenAndServe(":8082", nil)
}
