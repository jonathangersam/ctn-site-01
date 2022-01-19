# ctn-site-01

Assignment solution.

Author: Jonathan Gersam S. Lopez

## Installation
Backing SQL Storage is required.

1. Ensure DB table "image1" is defined. Below is the initial setup script:

```
-- CREATE TABLE

CREATE TABLE `images1` (
	`id` INT(11) NOT NULL AUTO_INCREMENT,
	`uid` TEXT NULL DEFAULT '',
	`filename` TEXT NULL DEFAULT '',
	`description` TEXT NULL DEFAULT '',
	`available` INT(1) NULL DEFAULT '0',
	`filetype` TEXT NULL DEFAULT '',
	`image_data` BLOB NULL DEFAULT '',
	`ts` TIMESTAMP NOT NULL DEFAULT current_timestamp(),
	PRIMARY KEY (`id`) USING BTREE
)

-- INSERT INIT ROWS

INSERT INTO images2 (filename, description, image_data) 
VALUES ('test1.png', 'first test image', 'iVBORw0KGgoAAAANSUhEUgAAACAAAAAgCAIAAAD8GO2jAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAADsMAAA7DAcdvqGQAAAD0SURBVEhL7Y9RDsMwCEN7/+PuAlmIHUSANE2naR/rk6VGYGz1KF/mHwqOOXScMjUxo1FeiRQeTMjXvAyhqSy8N2QjWEPQFcnhSF7gzrbkOkJhw91siQmdpCAeuAkE3BCSeWf8nSwduHnVbF7FFTLxATKNvixLJ/pwkiEy8QGp9Z5+VyCL4L4hyUEmPsonBfZ2UWCtFwWGNwLxUWQxui/K+uXdSQpkvdlhzXx3fEFF1t2Kx4kUO2FQIy8Qk7nXYyslDhnUSAoq9IUsi25VnI/kBRXas6BUNAemBYB3HRfn4M3IosDBpAZHK/YKbvAULHkKFpTyBiM7FH7ahdgUAAAAAElFTkSuQmCC'),
('test2.png', 'second test image', 'iVBORw0KGgoAAAANSUhEUgAAACAAAAAgCAIAAAD8GO2jAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAADsMAAA7DAcdvqGQAAAD0SURBVEhL7Y9RDsMwCEN7/+PuAlmIHUSANE2naR/rk6VGYGz1KF/mHwqOOXScMjUxo1FeiRQeTMjXvAyhqSy8N2QjWEPQFcnhSF7gzrbkOkJhw91siQmdpCAeuAkE3BCSeWf8nSwduHnVbF7FFTLxATKNvixLJ/pwkiEy8QGp9Z5+VyCL4L4hyUEmPsonBfZ2UWCtFwWGNwLxUWQxui/K+uXdSQpkvdlhzXx3fEFF1t2Kx4kUO2FQIy8Qk7nXYyslDhnUSAoq9IUsi25VnI/kBRXas6BUNAemBYB3HRfn4M3IosDBpAZHK/YKbvAULHkKFpTyBiM7FH7ahdgUAAAAAElFTkSuQmCC');

```

2. Set ENV var `DB_CONN1' with db connection string

* PATTERN: $USERNAME:$PASSWORD@tcp($HOSTNAME:$PORT)/$DBNAME
* SAMPLE: root:CHANGE_ME@tcp(localhost:3306)/MY_DB

## Running in Local machine
1. Run command
```
    go run bin/main.go
```

2. Use web browser to visit `localhost:8080/`

## Building Executable
1. Run the ff:
```
    go build bin/main.go
```

2. Compiled program available at main.exe

## API Endpoints

GET /image/{id}
* returns a single image data in json

GET /images
* returns list of image data in json
* TODO: query params format to parse it

POST /image
* inserts give image data to backend datastore
* request format:

```
    {
        "file_name": <string>,
        "description": <string>,
        "data": <string image data in base64 format. only supports PNG for now>,
    }
```

PATCH /image/{id}
* 'takes' an image out of the datastore if available.
* return error message if image is not available (i.e. previously taken)
* request format:
```
    {
        "take": <boolean>
    }
```

GET /health
* application health-check endpoint

## HTML Endpoints

GET /home
* browser-friendly landing page

GET /image/{id}/view
* displays image data and render of the image itself
* sample: localhost:8080/image/1/view

GET /
* defaults to home page


## Testing
TBD