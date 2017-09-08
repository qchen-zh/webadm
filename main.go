package main

import (
	// "flag"

	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", nil))
}
