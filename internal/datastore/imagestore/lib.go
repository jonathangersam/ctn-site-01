package imagestore

import (
	"ctn01/internal/schema"
	"errors"
)

var (
	ErrorImageNotFound     = errors.New("image doesn't exist")
	ErrorConnectionError   = errors.New("image storage connection error")
	ErrorImageExists       = errors.New("image already exists")
	ErrorImageAlreadyTaken = errors.New("image was already taken previously")
)

type ImageStore interface {
	GetImageByID(id int) (*schema.Image, error)
	GetImages(fromId, toId, afterId, size int) ([]*schema.Image, error)
	InsertImage(image schema.Image) error
	TakeImageById(id int) error
}
