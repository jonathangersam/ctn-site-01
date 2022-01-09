package inmem_imagestore

import (
	"ctn01/internal/datastore/imagestore"
	"ctn01/internal/schema"
	"sync"
)

type Store struct {
	localStorage map[int]*schema.Image
	muTakeImage  sync.Mutex // protects Take function
}

func NewImageStore() imagestore.ImageStore {
	store := Store{}

	// setup in-memory data store, to quickly get up to speed

	// populate init storage with dummy
	store.localStorage[1] = &schema.Image{
		Id:          1,
		Description: "first image",
		Available:   true,
	}

	store.localStorage[2] = &schema.Image{
		Id:          2,
		Description: "second image",
		Available:   true,
	}

	return &store
}

func (s *Store) GetImageByID(id int) (*schema.Image, error) {
	image, found := s.localStorage[id]
	if !found {
		return nil, imagestore.ErrorImageNotFound
	}

	return image, nil
}

func (s *Store) GetImages(fromId, toId, afterId, size int) ([]*schema.Image, error) {
	var matches []*schema.Image

	// helper fn
	idWithinRange := func(id int) bool {
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

func (s *Store) InsertImage(image schema.Image) error {
	// prevent inserting of new image
	_, imageExists := s.localStorage[image.Id]
	if imageExists {
		return imagestore.ErrorImageExists
	}

	s.localStorage[image.Id] = &image
	return nil
}

func (s *Store) TakeImageById(id int) error {
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
