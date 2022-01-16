package postImage

import (
	store "ctn01/internal/datastore/imagestore2"
	"ctn01/internal/entities"
	"ctn01/internal/handlers"
	"encoding/json"
	"log"
	"net/http"
)

var (
	insertImage  func(entities.Image) (int, error)
	getImageById func(uint64) (entities.Image, error)
)

type request struct {
	FileName    string `json:"file_name"`
	Description string `json:"description"`
	Data        []byte `json:"data"`
}

type response struct {
	Data handlers.HttpImageData `json:"data"`
}

func init() {
	insertImage = store.InsertImage
	getImageById = store.GetImageByID
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// get request
	req, err := parseRequest(r)
	if err != nil {
		handlers.WriteGenericResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	log.Printf("request received: %+v", req)

	// process request
	newImg, err := saveImage(req)
	if err != nil {
		handlers.WriteGenericResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	// reply
	status := http.StatusOK
	res := response{
		Data: handlers.HttpImageData{
			Id:          newImg.Id,
			Description: newImg.Description,
			Available:   newImg.Available,
			Code:        status,
		},
	}

	handlers.WriteResponse(w, status, res)
}

func parseRequest(r *http.Request) (request, error) {
	defer r.Body.Close()

	var req request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return request{}, err
	}

	return req, nil
}

func saveImage(req request) (entities.Image, error) {
	var img = entities.Image{
		Description: req.Description,
		Available:   true,
		Blob:        req.Data,
	}

	id, err := store.InsertImage(img) // this fn will auto-generate unique Id
	if err != nil {
		return entities.Image{}, err
	}

	log.Printf("last inserted row was: %d\n", id)

	uint64Id := uint64(id)
	return getImageById(uint64Id)
}
