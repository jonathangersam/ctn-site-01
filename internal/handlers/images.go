package handlers

import "net/http"

func ImagesHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

type ImagesEndpointRequest struct {
}

type ImagesEndpointResponse struct {
}
