package main

import (
	"github.com/joho/godotenv"
	"hundred-go/platform/authenticator"
	"hundred-go/platform/router"
	"log"
	"net/http"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("Failed to load the env vars: %v", err)
		log.Print("Will retrieve info from environment configuration variables.")
	}

	auth, err := authenticator.New()
	if err != nil {
		log.Fatalf("Failed to initialize the authenticator: %v", err)
	}

	rtr := router.New(auth)

	port, ok := os.LookupEnv("PORT")
	if ok == false {
		port = "3000"
	}

	log.Print("Server listening on http://localhost:" + port)
	if err := http.ListenAndServe("0.0.0.0:"+port, rtr); err != nil {
		log.Fatalf("There was an error with the http server: %v", err)
	}
}
