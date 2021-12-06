package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"

	"k8s_demo2/generate"
	"k8s_demo2/inbound"
	"k8s_demo2/outbound"

	middleware "github.com/deepmap/oapi-codegen/pkg/chi-middleware"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	outbound.CreditLimit()
}

func main() {

	var port = flag.Int("port", 8082, "Port for test HTTP server")
	flag.Parse()

	swagger, err := generate.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil

	// Create an instance of our handler which satisfies the generated interface
	petStore := inbound.NewPetStore()

	// This is how you set up a basic chi router
	r := chi.NewRouter()

	// Use our validation middleware to check all requests against the
	// OpenAPI schema.
	r.Use(middleware.OapiRequestValidator(swagger))

	// We now register our petStore above as the handler for the interface
	generate.HandlerFromMux(petStore, r)

	s := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("0.0.0.0:%d", *port),
	}

	// And we serve HTTP until the world ends.
	log.Fatal(s.ListenAndServe())
}
