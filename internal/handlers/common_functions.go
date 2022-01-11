package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func GetMuxVar(r *http.Request, key string) string {
	vars := mux.Vars(r)
	return vars[key]
}

func SetContentTypeToJSON(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func WriteGenericResponse(w http.ResponseWriter, code int, desc string) {
	resp := GenericResponse{
		Data: GenericErrorData{
			Code:        code,
			Description: desc,
		},
	}
	WriteResponse(w, code, resp)
}

func WriteResponse(w http.ResponseWriter, code int, v interface{}) {
	SetContentTypeToJSON(w)
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Println(err)
	}
}
