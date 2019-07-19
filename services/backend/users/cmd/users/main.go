package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// ---------------- Initial set up

//  Environment variables

type envVariables struct {
	port string
}

var config envVariables

func initialSetup() {
	config = envVariables{
		port: getEnv("PORT", "80"),
	}
}

// ---------------- Function helpers

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func ifError(err error, msg string) {
	if err != nil {
		log.Fatal(err)
	}
}

// ---------------- Request handler

func main() {
	initialSetup()
	http.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from users microservice !!!")
	})

	log.Fatal(http.ListenAndServe(":"+config.port, nil))

}
