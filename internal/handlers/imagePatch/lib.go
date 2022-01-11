package imagePatch

import (
	"ctn01/internal/datastore/imagestore"
	"ctn01/internal/datastore/imagestore/inmem_imagestore"
	"ctn01/internal/handlers"
	"encoding/json"
	"net/http"
	"strconv"
)

var (
	store imagestore.ImageStore
)

type request struct {
	Take bool `json:"take"`
}

type response struct {
	Data httpImageDataWithFile `json:"data"`
}

type httpImageDataWithFile struct {
	handlers.HttpImageData
	File string `json:"file"`
}

func init() {
	store, _ = inmem_imagestore.Connect()
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// get input
	id := handlers.GetMuxVar(r, "id")

	intId, err := strconv.Atoi(id)
	if err != nil {
		handlers.WriteGenericResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	uint64Id := uint64(intId)

	// parse request
	req, err := parseRequest(r)
	if err != nil {
		handlers.WriteGenericResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if !req.Take {
		// nothing to do
		handlers.WriteGenericResponse(w, http.StatusOK, "nothing to do")
		return
	}

	// take image in DB
	img, err := store.TakeImageById(uint64Id)
	if err != nil {
		handlers.WriteGenericResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	// reply success
	status := http.StatusOK
	resp := response{
		Data: httpImageDataWithFile{
			HttpImageData: handlers.HttpImageData{
				Id:          uint64Id,
				Description: img.Description,
				Available:   img.Available,
				Code:        status,
			},
			File: "TBD", // TODO fill this
		},
	}

	handlers.WriteResponse(w, status, resp)
}

func parseRequest(r *http.Request) (request, error) {
	var req request
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}
