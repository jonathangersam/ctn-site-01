package imageHandler

import "ctn01/internal/handlers"

type request struct {
	FileName    string `json:"file_name"`
	Description string `json:"description"`
	Data        []byte `json:"data"`
}

type response struct {
	Data handlers.HttpImageData `json:"data"`
}

//type HttpImageData struct {
//	Id          string `json:"id"`
//	Description string `json:"description"`
//	Available   bool   `json:"available"`
//	Code        int    `json:"code"`
//}
