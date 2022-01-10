package main

import (
	"ctn01/internal/handlers/healthGet"
	"ctn01/internal/handlers/imageGet"
	"ctn01/internal/handlers/imagePatch"
	"ctn01/internal/handlers/imagePost"
	"ctn01/internal/handlers/imagesGet"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func simpleGreet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome!"))
	}
}

func main() {
	fmt.Println("Application starting...")

	router := mux.NewRouter()

	router.HandleFunc("/health", healthGet.Handler)
	router.HandleFunc("/image/{id:[0-9]+}", imageGet.Handler).Methods("GET")
	router.HandleFunc("/image/{id:[0-9]+}", imagePatch.Handler).Methods("PATCH")
	router.HandleFunc("/image", imagePost.Handler).Methods("POST")
	router.HandleFunc("/images", imagesGet.Handler).Methods("GET")
	router.HandleFunc("/", simpleGreet())

	fmt.Println("Listening on port 8080")

	log.Fatal(http.ListenAndServe(":8080", router))
}
