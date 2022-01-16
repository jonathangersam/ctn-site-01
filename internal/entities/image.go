package entities

const (
	imgSmileyFacePngBase64 = "iVBORw0KGgoAAAANSUhEUgAAACAAAAAgCAIAAAD8GO2jAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAADsMAAA7DAcdvqGQAAAD0SURBVEhL7Y9RDsMwCEN7/+PuAlmIHUSANE2naR/rk6VGYGz1KF/mHwqOOXScMjUxo1FeiRQeTMjXvAyhqSy8N2QjWEPQFcnhSF7gzrbkOkJhw91siQmdpCAeuAkE3BCSeWf8nSwduHnVbF7FFTLxATKNvixLJ/pwkiEy8QGp9Z5+VyCL4L4hyUEmPsonBfZ2UWCtFwWGNwLxUWQxui/K+uXdSQpkvdlhzXx3fEFF1t2Kx4kUO2FQIy8Qk7nXYyslDhnUSAoq9IUsi25VnI/kBRXas6BUNAemBYB3HRfn4M3IosDBpAZHK/YKbvAULHkKFpTyBiM7FH7ahdgUAAAAAElFTkSuQmCC"
)

var (
	ImageSmileyFacePng = Image{
		Filename:    "smiley_face_32x32.png",
		Description: "a yellow smiley face",
		Blob:        []byte(imgSmileyFacePngBase64),
	}
)

type Image struct {
	Id          uint64 // key
	UID         string // unique identifier (temp)
	Filename    string
	Description string
	Available   bool
	Filetype    string
	Blob        []byte // raw image data
}
