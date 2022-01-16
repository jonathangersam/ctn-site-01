package getImages

import (
	store "ctn01/internal/datastore/imagestore2"
	"ctn01/internal/entities"
	"ctn01/internal/handlers"
	"encoding/json"
	"log"
	"net/http"
)

type response struct {
	Size int                      `json:"size"`
	Data []handlers.HttpImageData `json:"data"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get all images
	images, error := store.GetImages(0, 0, 0, -1)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(error.Error()))
		return
	}

	// reply
	httpImageData := make([]handlers.HttpImageData, len(images))
	for i, image := range images {
		httpImageData[i] = getHttpImageDataFrom(image)
		httpImageData[i].Code = http.StatusOK
	}

	resp := response{
		Size: len(httpImageData),
		Data: httpImageData,
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("json encode of response failed: %s\n", err)
	}
}

func getHttpImageDataFrom(img entities.Image) handlers.HttpImageData {
	return handlers.HttpImageData{
		Id:          img.Id,
		Description: img.Description,
		Available:   img.Available,
	}
}
