package main

import (
	"ctn01/internal/datastore/imagestore2"
	"ctn01/internal/handlers/getImage"
	"ctn01/internal/handlers/getImageView"
	"ctn01/internal/handlers/getImages"
	"ctn01/internal/handlers/health"
	"ctn01/internal/handlers/home"
	"ctn01/internal/handlers/patchImage"
	"ctn01/internal/handlers/postImage"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Application starting...")

	imagestore2.Connect()
	defer imagestore2.Close()

	router := mux.NewRouter()

	router.HandleFunc("/health", health.Handler)
	router.HandleFunc("/image/{id:[0-9]+}/view", getImageView.Handler()).Methods("GET")
	router.HandleFunc("/image/{id:[0-9]+}", getImage.Handler).Methods("GET")
	router.HandleFunc("/image/{id:[0-9]+}", patchImage.Handler).Methods("PATCH")
	router.HandleFunc("/image", postImage.Handler).Methods("POST")
	router.HandleFunc("/images", getImages.Handler).Methods("GET")
	router.HandleFunc("/home", home.Handler())
	router.HandleFunc("/", home.Handler())

	fmt.Println("Listening on port 8080")

	log.Fatal(http.ListenAndServe(":8080", router))
}
