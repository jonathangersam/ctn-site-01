package healthGet

import (
	"ctn01/internal/handlers"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	handlers.WriteGenericResponse(w, http.StatusOK, "ok")
}
