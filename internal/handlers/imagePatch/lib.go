package imagePatch

import (
	"ctn01/internal/datastore/imagestore"
	"ctn01/internal/datastore/imagestore/inmem_imagestore"
	"ctn01/internal/handlers"
	"encoding/json"
	"net/http"
)

var (
	store imagestore.ImageStore
)

type request struct {
	Take bool `json:"take"`
}

type response struct {
	Data []handlers.HttpImageData `json:"data"`
}

func init() {
	store, _ = inmem_imagestore.Connect()
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// parse request
	_, err := parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		//w.Write()
		return
	}

	// take image in DB

	// if not found, reply w/ error

	// reply
}

func parseRequest(r *http.Request) (*request, error) {
	var req request
	err := json.NewDecoder(r.Body).Decode(&req)
	return &req, err
}
