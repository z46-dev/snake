package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting server...")
	http.Handle("/", http.FileServer(http.Dir("./public")))
	fmt.Println("Listening on port 3000 (HTTP)...")

	http.ListenAndServe("127.0.0.1:3000", nil)
}
