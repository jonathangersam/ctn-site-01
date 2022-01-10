package entities

type Image struct {
	Id          uint64 // key
	UID         string // unique identifier (temp)
	Description string
	Available   bool
	Blob        []byte // raw image data
}
