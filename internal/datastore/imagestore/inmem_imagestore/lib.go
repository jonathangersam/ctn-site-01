// Package inmem_imagestore implements an image store solution using in-memory storage.
// Used as fallback and testing solution when DB is not available.
//
// Also used for practical purposes when we just want to run the code without havig to setup
// an actual database.
package inmem_imagestore

import (
	"ctn01/internal/datastore/imagestore"
	"ctn01/internal/entities"
	"sort"
	"sync"

	"github.com/google/uuid"
)

const (
	maxIdFilter uint64 = 18446744073709551615
)

var (
	localStorage = map[uint64]entities.Image{}
	muStore      = sync.Mutex{} // protects Take function
)

func init() {
	// populate storage with dummy
	uid := uuid.New().String()
	localStorage[1] = entities.Image{
		Id:          1,
		UID:         uid,
		Description: "first image",
		Available:   true,
	}

	uid = uuid.New().String()
	localStorage[2] = entities.Image{
		Id:          2,
		Description: "second image",
		Available:   true,
	}

	uid = uuid.New().String()
	img := entities.ImageSmileyFacePng
	img.Id = 3
	img.UID = uid
	img.Available = true
	localStorage[3] = img
}

type store struct{}

// Connect returns a singleton
func Connect() (imagestore.ImageStore, error) {
	return &store{}, nil
}

func (s *store) GetImageByID(id uint64) (entities.Image, error) {
	image, found := localStorage[id]
	if !found {
		return entities.Image{}, imagestore.ErrorImageNotFound
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
func (s *store) GetImages(fromId, toId, afterId uint64, size int) ([]entities.Image, error) {
	var matches []entities.Image

	// helper fn
	upperLimit := toId
	if upperLimit == 0 {
		upperLimit = maxIdFilter
	}

	idWithinRange := func(id uint64) bool {
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

func getMaxId() uint64 {
	// generate sorted IDs
	keys := make([]uint64, 0, len(localStorage))
	for k := range localStorage {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	// return last ID
	return keys[len(keys)-1]
}

func (s *store) InsertImage(image entities.Image) (entities.Image, error) {
	// prevent inserting of new image
	muStore.Lock()

	newId := getMaxId() + 1
	image.Id = newId

	//_, imageExists := localStorage[image.Id]
	//if imageExists {
	//	return nil, imagestore.ErrorImageExists
	//}

	localStorage[image.Id] = image
	muStore.Unlock()
	return image, nil
}

func (s *store) TakeImageById(id uint64) (entities.Image, error) {
	// return error if not found
	muStore.Lock()
	image, imageExists := localStorage[id]
	if !imageExists {
		muStore.Unlock()
		return entities.Image{}, imagestore.ErrorImageNotFound
	}

	// return error if unavailable
	if !image.Available {
		muStore.Unlock()
		return entities.Image{}, imagestore.ErrorImageAlreadyTaken
	}

	image.Available = false
	localStorage[id] = image
	muStore.Unlock()

	return image, nil
}
