package imagestore

import (
	"ctn01/internal/entities"
	"errors"
)

var (
	ErrorImageNotFound     = errors.New("image doesn't exist")
	ErrorConnectionError   = errors.New("image storage connection error")
	ErrorImageExists       = errors.New("image already exists")
	ErrorImageAlreadyTaken = errors.New("image was already taken previously")
)

type ImageStore interface {
	GetImageByID(id uint64) (entities.Image, error)
	GetImages(fromId, toId, afterId uint64, size int) ([]entities.Image, error)
	InsertImage(image entities.Image) (entities.Image, error) // return generated image since ID is auto-gen
	TakeImageById(id uint64) (entities.Image, error)
}
