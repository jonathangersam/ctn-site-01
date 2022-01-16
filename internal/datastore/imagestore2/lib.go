package imagestore2

import (
	"ctn01/internal/entities"
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var (
	ErrorNoRowsMatch = errors.New("no rows matched")

	db         *sql.DB // use singleton global var as it fits the complexity of the application.
	connString string
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
		return entities.Image{}, err
	}
	defer rows.Close()

	result, err := scanRows(rows)
	if err != nil {
		return entities.Image{}, err
	}

	if len(result) == 0 {
		return entities.Image{}, ErrorNoRowsMatch
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
	isAvailable := 1

	// do insert
	rows, err := db.Query(`
INSERT INTO images1 (filename, description, available, filetype, image_data)
VALUES (?, ?, ?, ?, ?) RETURNING id`, img.Filename, img.Description, isAvailable, img.Filetype, img.Blob)
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
	return entities.Image{}, nil
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
		img.Available = intAvailable == 1

		result = append(result, img)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
