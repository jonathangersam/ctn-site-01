package handlers

type HttpImageData struct {
	Id          uint64 `json:"id"`
	Description string `json:"description"`
	Available   bool   `json:"available"`
	Code        int    `json:"code"`
}
