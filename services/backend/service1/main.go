package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/api/service1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from service 1 !!!")
	})

	http.HandleFunc("/api/service1/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from service 1 /test !!!")
	})

	log.Fatal(http.ListenAndServe(":80", nil))

}
