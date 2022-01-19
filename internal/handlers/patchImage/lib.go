package patchImage

import (
	store "ctn01/internal/datastore/imagestore2"

	"ctn01/internal/entities"
	"ctn01/internal/handlers"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

var (
	//store imagestore.ImageStore
	takeImageByID func(uint64) (entities.Image, error)
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
	//store, _ = inmem_imagestore.Connect()
	takeImageByID = store.TakeImageById
}

func Handler(w http.ResponseWriter, r *http.Request) {
	//log.Printf("r.Host: %s, r.URL.Path: %s, r.RequestURI: %s\n", r.Host, r.URL.Path, r.RequestURI)

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
	//img, err := store.TakeImageById(uint64Id)
	img, err := takeImageByID(uint64Id)
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
			File: getImageViewURI(r),
		},
	}

	handlers.WriteResponse(w, status, resp)
}

func parseRequest(r *http.Request) (request, error) {
	var req request
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func getImageViewURI(r *http.Request) string {
	return fmt.Sprintf("%s%s/view", r.Host, r.RequestURI)
}
