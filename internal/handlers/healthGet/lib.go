package healthGet

import (
	"encoding/json"
	"net/http"
)

type response struct {
	StatusCode int    `json:"status_code"`
	Status     string `json:"status"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	resp := response{
		StatusCode: http.StatusOK,
		Status:     "ok",
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(resp)
}
