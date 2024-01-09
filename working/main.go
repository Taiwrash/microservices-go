package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello World!")
		data, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Could not read the request body", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "Hello %s\n", data)
	})

	http.HandleFunc("/bye", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Goodbye World!")
	})
	http.ListenAndServe(":9090", nil)
}
