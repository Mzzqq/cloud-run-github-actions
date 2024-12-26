package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, Cloud Run!")
	})

	port := "8080"
	fmt.Println("Listening on port", port)
	http.ListenAndServe(":"+port, nil)
}
