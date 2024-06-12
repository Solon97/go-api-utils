package rest

import (
	"fmt"
	"net/http"
)

type Response struct {
	Body       any `json:"body"`
	StatusCode int `json:"status"`
}

type EndpointFunc func(w http.ResponseWriter, r *http.Request) (Response, error)

func HandleResponse(endpointFunc EndpointFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response, err := endpointFunc(w, r)
		if err != nil {
			response.Body, response.StatusCode = handleErrorResponse(err)
		}

		w.WriteHeader(response.StatusCode)
		if response.Body != nil {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(fmt.Sprintf("%v", response.Body)))
		}
	})
}
