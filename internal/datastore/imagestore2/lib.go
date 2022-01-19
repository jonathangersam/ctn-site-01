package imagestore2

import (
	"ctn01/internal/entities"
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"sync"
)

const (
	imageAvailable   = 1
	imageUnavailable = 0
)

var (
	ErrorNotFound         = errors.New("image doesn't exist")
	ErrorImageUnavailable = errors.New("image already taken")
	ErrorNoRowsMatch      = errors.New("no rows matched")

	db          *sql.DB // use singleton global var as it fits the complexity of the application.
	connString  string
	takeImageMu sync.Mutex
)

func init() {
	connString = os.Getenv("DB_CONN1")
	//if err := Connect(); err != nil {
	//	panic(err)
	//}
}

func Connect() {
	var err error
	db, err = sql.Open("mysql", connString)
	if err != nil {
		panic(err)
	}

	if db.Ping() != nil {
		panic(err)
	}
}

func GetImageByID(id uint64) (entities.Image, error) {
	rows, err := db.Query(`SELECT id, filename, description, available, filetype, image_data FROM images1 WHERE id = ? LIMIT 1`, id)
	if err != nil {
		return emptyImage(), err
	}
	defer rows.Close()

	result, err := scanRows(rows)
	if err != nil {
		return emptyImage(), err
	}

	if len(result) == 0 {
		return emptyImage(), ErrorNotFound
	}

	return result[0], nil
}

func GetImages(fromId, toId, afterId uint64, size int) ([]entities.Image, error) {
	if toId == 0 {
		toId = 9999 // TODO: change to SELECT MAX(id) FROM images1
	}

	rows, err := db.Query(`
SELECT id, filename, description, available, filetype, image_data FROM images1
WHERE id >= ? AND id > ? AND id <= ? LIMIT ?`, fromId, afterId, toId, size)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result, err := scanRows(rows)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func InsertImage(img entities.Image) (int, error) {
	// do insert
	rows, err := db.Query(`
INSERT INTO images1 (filename, description, available, filetype, image_data)
VALUES (?, ?, ?, ?, ?) RETURNING id`, img.Filename, img.Description, imageAvailable, img.Filetype, img.Blob)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	// get id
	var id int
	for rows.Next() {
		rows.Scan(&id)
	}

	if err := rows.Err(); err != nil {
		return 0, nil
	}

	return id, nil
}

func TakeImageById(id uint64) (entities.Image, error) {
	// if images not found, error
	img, err := GetImageByID(id)
	if err != nil {
		return emptyImage(), err
	}

	// if image is unavailable, error
	if !img.Available {
		return img, ErrorImageUnavailable
	}

	// mark image as unavailable (already taken)
	takeImageMu.Lock()
	err = doTakeImage(id)
	takeImageMu.Unlock()
	if err != nil {
		log.Printf("doTakeImage(%d) returned error: %s\n", id, err)
		return emptyImage(), ErrorImageUnavailable
	}

	img.Available = false
	return img, nil
}

func doTakeImage(id uint64) error {
	// mark image as unavailable (already taken)
	resp, err := db.Exec(`
UPDATE images1 SET available = ? WHERE id = ? AND available = ?
`, imageUnavailable, id, imageAvailable)
	if err != nil {
		return err
	}

	// if error / no rows matched, return error
	count, err := resp.RowsAffected()
	if err != nil || count == 0 {
		return ErrorImageUnavailable
	}

	return nil
}

func Close() error {
	return db.Close()
}

func scanRows(rows *sql.Rows) ([]entities.Image, error) {
	var result []entities.Image
	for rows.Next() {
		var img entities.Image
		var intAvailable int
		if err := rows.Scan(&img.Id, &img.Filename, &img.Description, &intAvailable, &img.Filetype, &img.Blob); err != nil {
			return nil, err
		}
		img.Available = intAvailable == imageAvailable

		result = append(result, img)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

// emptyImage is a convenience function returning zero value of image struct
func emptyImage() entities.Image {
	return entities.Image{}
}
