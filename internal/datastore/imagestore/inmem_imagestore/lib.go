package inmem_imagestore

import (
	"ctn01/internal/datastore/imagestore"
	"ctn01/internal/entities"
	"sync"

	"github.com/google/uuid"
)

type Store struct {
	localStorage map[string]*entities.Image
	muTakeImage  sync.Mutex // protects Take function
}

func New() imagestore.ImageStore {
	store := Store{
		localStorage: map[string]*entities.Image{},
	}

	// setup in-memory data store, to quickly get up to speed

	// populate init storage with dummy
	newId := uuid.New().String()
	store.localStorage[newId] = &entities.Image{
		Id:          newId,
		Description: "first image",
		Available:   true,
	}

	newId = uuid.New().String()
	store.localStorage[newId] = &entities.Image{
		Id:          newId,
		Description: "second image",
		Available:   true,
	}

	return &store
}

func (s *Store) GetImageByID(id string) (*entities.Image, error) {
	image, found := s.localStorage[id]
	if !found {
		return nil, imagestore.ErrorImageNotFound
	}

	return image, nil
}

func (s *Store) GetImages(fromId, toId, afterId string, size int) ([]*entities.Image, error) {
	var matches []*entities.Image

	// helper fn
	idWithinRange := func(id string) bool {
		return id >= fromId && id > afterId && id <= toId
	}

	for _, image := range s.localStorage {
		if !idWithinRange(image.Id) {
			continue
		}

		matches = append(matches, image)
	}

	return matches, nil
}

func (s *Store) InsertImage(image entities.Image) (*entities.Image, error) {
	// prevent inserting of new image
	image.Id = uuid.New().String() // TODO: use something else
	_, imageExists := s.localStorage[image.Id]
	if imageExists {
		return nil, imagestore.ErrorImageExists
	}

	s.localStorage[image.Id] = &image
	return &image, nil
}

func (s *Store) TakeImageById(id string) error {
	// if non existing, error
	image, imageExists := s.localStorage[id]
	if !imageExists {
		return imagestore.ErrorImageNotFound
	}

	s.muTakeImage.Lock()
	if !image.Available {
		return imagestore.ErrorImageAlreadyTaken
	}
	image.Available = false
	s.muTakeImage.Unlock()

	return nil
}
