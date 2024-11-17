package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world\n")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting web server on 3000...")

	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
