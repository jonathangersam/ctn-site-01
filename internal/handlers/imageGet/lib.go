package imageGet

import (
	"ctn01/internal/datastore/imagestore"
	"ctn01/internal/datastore/imagestore/inmem_imagestore"
	"ctn01/internal/handlers"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const (
	errorIdNotNumeric = "search id must be numeric"
)

var (
	store imagestore.ImageStore
)

//type request struct {
//}

type response struct {
	Data    handlers.HttpImageData `json:"data"`
	Error   bool                   `json:"error"`
	Message string                 `json:"message"`
}

func init() {
	store, _ = inmem_imagestore.Connect()
}

func Handler(w http.ResponseWriter, r *http.Request) {
	responder := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")

	// get image by id from db
	vars := mux.Vars(r)
	searchId := vars["id"]

	// validate input
	intSearchId, err := strconv.Atoi(searchId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		responder.Encode(response{
			Data: handlers.HttpImageData{
				Code: http.StatusBadRequest,
			},
			Error:   true,
			Message: errorIdNotNumeric,
		})
		return
	}

	uint64SearchId := uint64(intSearchId)

	// fetch the image
	img, err := store.GetImageByID(uint64SearchId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		responder.Encode(response{
			Data: handlers.HttpImageData{
				Code: http.StatusInternalServerError,
			},
			Error:   true,
			Message: err.Error(),
		})
		return
	}

	// reply success
	resp := response{
		Data: handlers.HttpImageData{
			Id:          img.Id,
			Description: img.Description,
			Available:   img.Available,
			Code:        http.StatusOK,
		},
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("json encoding of response failed: %s\n", err)
	}
}
