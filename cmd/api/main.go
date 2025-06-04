package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	fmt.Println("Listening on port 8080")
	fmt.Println("GET http://127.0.0.1:8080")

	http.ListenAndServe(":8080", nil)
}
