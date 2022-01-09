package main

import (
	"ctn01/internal/handlers"
	"fmt"
	"log"
	"net/http"
)

func simpleGreet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome!"))
	}
}

func main() {
	fmt.Println("Application starting...")

	http.HandleFunc("/", simpleGreet())
	http.HandleFunc("/image", handlers.ImageHandler())
	http.HandleFunc("/images", handlers.ImagesHandler())

	fmt.Println("Listening on port 8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
