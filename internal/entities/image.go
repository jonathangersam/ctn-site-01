package entities

type Image struct {
	Id          string // unique identifier
	Description string
	Available   bool
	Blob        []byte // raw image data
}
