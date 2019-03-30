package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello world")
	})

	log.Println("Serving on localhost:1313")
	err := http.ListenAndServe(":1313", nil)
	log.Fatal(err)
}
