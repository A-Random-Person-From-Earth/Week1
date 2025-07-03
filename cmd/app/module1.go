package main

import (
	"fmt"
	"github.com/A-Random-Person-From-Earth/go-camp/internal/greet"
	"net/http"
)

func main() {
	// Hello service that says hi to the world
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		message:= greet.Greet(name)
		_, err := fmt.Fprint(w, message)
if err != nil {
    http.Error(w, "Failed to write response", http.StatusInternalServerError)
    return
}
	})

	fmt.Println("Starting on :8080")
	err := http.ListenAndServe(":8080", nil)
if err != nil {
    fmt.Printf("Server failed to start: %v\n", err)
}
}
