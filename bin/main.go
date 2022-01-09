package main

import (
	"ctn01/internal/handlers/imageHandler"
	"ctn01/internal/handlers/imagesHandler"
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
	http.HandleFunc("/image", imageHandler.GetHandler())
	http.HandleFunc("/images", imagesHandler.GetHandler())

	fmt.Println("Listening on port 8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
