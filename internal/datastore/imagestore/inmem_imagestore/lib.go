// Package inmem_imagestore implements an image store solution using in-memory storage.
// Used as fallback and testing solution when DB is not available.
//
// Also used for practical purposes when we just want to run the code without havig to setup
// an actual database.
package inmem_imagestore

import (
	"ctn01/internal/datastore/imagestore"
	"ctn01/internal/entities"
	"sync"

	"github.com/google/uuid"
)

const (
	maxIdFilter = "zz"
)

var (
	localStorage = map[string]*entities.Image{}
	muTakeImage  = sync.Mutex{} // protects Take function
)

func init() {
	// populate storage with dummy
	newId := uuid.New().String()
	localStorage[newId] = &entities.Image{
		Id:          newId,
		Description: "first image",
		Available:   true,
	}

	newId = uuid.New().String()
	localStorage[newId] = &entities.Image{
		Id:          newId,
		Description: "second image",
		Available:   true,
	}
}

type store struct{}

// New returns a singleton
func New() imagestore.ImageStore {
	return &store{}
}

func (s *store) GetImageByID(id string) (*entities.Image, error) {
	image, found := localStorage[id]
	if !found {
		return nil, imagestore.ErrorImageNotFound
	}

	return image, nil
}

// GetImages returns all images that match the search criteria.
//
// Args fromId and toId specify start and end (both inclusive) of retrieval.
// If arg toId is empty, it is treated as max-value.
//
// Images with ID greater than afterId argument will be returned (exclusive).
//
// If size == -1, no limit on number of values returned
func (s *store) GetImages(fromId, toId, afterId string, size int) ([]*entities.Image, error) {
	var matches []*entities.Image

	// helper fn
	upperLimit := toId
	if upperLimit == "" {
		upperLimit = maxIdFilter
	}

	idWithinRange := func(id string) bool {
		return id >= fromId && id > afterId && id <= upperLimit
	}

	for _, image := range localStorage {
		if !idWithinRange(image.Id) {
			continue
		}

		matches = append(matches, image)
	}

	return matches, nil
}

func (s *store) InsertImage(image entities.Image) (*entities.Image, error) {
	// prevent inserting of new image
	image.Id = uuid.New().String()
	_, imageExists := localStorage[image.Id]
	if imageExists {
		return nil, imagestore.ErrorImageExists
	}

	localStorage[image.Id] = &image
	return &image, nil
}

func (s *store) TakeImageById(id string) error {
	// if non existing, error
	image, imageExists := localStorage[id]
	if !imageExists {
		return imagestore.ErrorImageNotFound
	}

	muTakeImage.Lock()
	if !image.Available {
		return imagestore.ErrorImageAlreadyTaken
	}
	image.Available = false
	muTakeImage.Unlock()

	return nil
}
