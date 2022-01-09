package handlers

import "net/http"

func ImageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

type ImageEndpointRequest struct {
}

type ImageEndpointResponse struct {
}
