package imageGet

import (
	"ctn01/internal/datastore/imagestore"
	"ctn01/internal/datastore/imagestore/inmem_imagestore"
	"ctn01/internal/handlers"
	"net/http"
	"strconv"
)

const (
	errorIdNotNumeric = "search id must be numeric"
)

var (
	store imagestore.ImageStore
)

type response struct {
	Data handlers.HttpImageData `json:"data"`
}

func init() {
	store, _ = inmem_imagestore.Connect()
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// get input
	searchId := handlers.GetMuxVar(r, "id")

	intSearchId, err := strconv.Atoi(searchId)
	if err != nil {
		handlers.WriteGenericResponse(w, http.StatusBadRequest, errorIdNotNumeric)
		return
	}

	uint64SearchId := uint64(intSearchId)

	// fetch the image
	img, err := store.GetImageByID(uint64SearchId)
	if err != nil {
		handlers.WriteGenericResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	// reply success
	status := http.StatusOK
	resp := response{
		Data: handlers.HttpImageData{
			Id:          img.Id,
			Description: img.Description,
			Available:   img.Available,
			Code:        status,
		},
	}

	handlers.WriteResponse(w, status, resp)
}
