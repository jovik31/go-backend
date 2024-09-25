package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	hostname := os.Getenv("HOSTNAME")
	fmt.Fprintf(w, "Hello world from %s!\n", hostname)
}

func main() {
	log.SetOutput(os.Stdout)
	log.Println("Starting server on port :80")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
