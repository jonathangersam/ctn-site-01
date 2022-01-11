# ctn-site-01

Assignment solution.

Author: Jonathan Gersam S. Lopez

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
* sample: localhost:8080/image/3/view

GET /
* defaults to home page

## Building
TBD

## Testing
TBD