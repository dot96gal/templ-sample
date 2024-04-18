package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dot96gal/templ-sample/handlers"
)

func main() {
	http.HandleFunc("GET /", handlers.IndexHandler)

	fmt.Println("Listening on :3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
