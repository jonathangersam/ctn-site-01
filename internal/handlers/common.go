package handlers

type HttpImageData struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	Available   bool   `json:"available"`
	Code        int    `json:"code"`
}
