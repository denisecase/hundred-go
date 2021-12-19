package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"hundred-go/platform/authenticator"
	"hundred-go/platform/router"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("Failed to load the env vars: %v", err)
		log.Print("Will retrieve from environment configuration variables.")
	}

	auth, err := authenticator.New()
	if err != nil {
		log.Fatalf("Failed to initialize the authenticator: %v", err)
	}

	rtr := router.New(auth)

	log.Print("Server listening on http://localhost:3000/")
	if err := http.ListenAndServe("0.0.0.0:3000", rtr); err != nil {
		log.Fatalf("There was an error with the http server: %v", err)
	}
}
