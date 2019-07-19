package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/api/service2", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from service 2 !!!")
	})

	log.Fatal(http.ListenAndServe(":80", nil))

}
