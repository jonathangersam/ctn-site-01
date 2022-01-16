package getImageView

import (
	"ctn01/internal/datastore/imagestore"
	"ctn01/internal/datastore/imagestore/inmem_imagestore"
	"ctn01/internal/entities"
	"ctn01/internal/handlers"
	"html/template"
	"net/http"
	"strconv"
)

var (
	store imagestore.ImageStore
)

type data struct {
	Name         string
	Image        entities.Image
	ImageContent string
}

func init() {
	store, _ = inmem_imagestore.Connect()
}

func Handler() http.HandlerFunc {
	tmpl := template.Must(template.ParseFiles("web/templates/imageViewer.html"))

	return func(w http.ResponseWriter, r *http.Request) {
		// get input
		id := handlers.GetMuxVar(r, "id")
		intId, err := strconv.Atoi(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("id must be numeric"))
			return
		}
		uint64Id := uint64(intId)

		// fetch image
		img, err := store.GetImageByID(uint64Id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("img with given id not found"))
			return
		}

		// display image in html
		d := data{
			Name:         "Jonathan Gersam S. Lopez",
			Image:        img,
			ImageContent: string(img.Blob),
		}

		tmpl.Execute(w, d)
	}
}
