package handlers

type HttpImageData struct {
	Id          uint64 `json:"id"`
	Description string `json:"description"`
	Available   bool   `json:"available"`
	Code        int    `json:"code"`
}

type GenericResponse struct {
	Data GenericErrorData `json:"data"`
}

type GenericErrorData struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}
