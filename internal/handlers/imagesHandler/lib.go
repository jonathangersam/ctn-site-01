package imagesHandler

import (
	"ctn01/internal/datastore/imagestore"
	"ctn01/internal/datastore/imagestore/inmem_imagestore"
	"ctn01/internal/entities"
	"ctn01/internal/handlers"
	"encoding/json"
	"log"
	"net/http"
)

var (
	msgOnlyGetSupported = []byte("only GET method allowed")

	store imagestore.ImageStore
)

func init() {
	store = inmem_imagestore.New()
}

func GetHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// validate method
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write(msgOnlyGetSupported)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		// get all images
		images, error := store.GetImages("", "", "", -1)
		if error != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(error.Error()))
			return
		}

		// reply
		httpImageData := make([]handlers.HttpImageData, len(images))
		for i, image := range images {
			httpImageData[i] = *getHttpImageDataFrom(image)
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
}

//type request struct {
//}

type response struct {
	Size int                      `json:"size"`
	Data []handlers.HttpImageData `json:"data"`
}

func getHttpImageDataFrom(img *entities.Image) *handlers.HttpImageData {
	return &handlers.HttpImageData{
		Id:          img.Id,
		Description: img.Description,
		Available:   img.Available,
	}
}
