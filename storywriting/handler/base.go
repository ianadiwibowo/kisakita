package handler

import (
	"encoding/json"
	"net/http"
)

// StoriesHandler ...
type StoriesHandler struct {
}

// NewStoriesHandler initializes fresh handler
func NewStoriesHandler() *StoriesHandler {
	return &StoriesHandler{}
}

// SuccessResponse is the standard success data-meta response
type SuccessResponse struct {
	Data interface{} `json:"data"`
	Meta Meta        `json:"meta"`
}

// FailedResponse is the standard failed errors-meta response
type FailedResponse struct {
	Error interface{} `json:"error"`
	Meta  Meta        `json:"meta"`
}

// Message ...
type Message struct {
	Message string `json:"message"`
}

// Meta is used to explain Data
type Meta struct {
	HTTPStatus int `json:"http_status"`
}

// OK 200
func respondWithOK(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")

	httpStatus := http.StatusOK
	w.WriteHeader(httpStatus)

	response := SuccessResponse{
		Data: data,
		Meta: Meta{
			HTTPStatus: httpStatus,
		},
	}

	_ = json.NewEncoder(w).Encode(response)
}

// Created 201
func respondWithCreated(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")

	httpStatus := http.StatusCreated
	w.WriteHeader(httpStatus)

	response := SuccessResponse{
		Data: data,
		Meta: Meta{
			HTTPStatus: httpStatus,
		},
	}

	_ = json.NewEncoder(w).Encode(response)
}

// Bad Request 400
func respondWithBadRequest(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")

	httpStatus := http.StatusBadRequest
	w.WriteHeader(httpStatus)

	response := FailedResponse{
		Error: data,
		Meta: Meta{
			HTTPStatus: httpStatus,
		},
	}

	_ = json.NewEncoder(w).Encode(response)
}

// Not Found 404
func respondWithNotFound(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	httpStatus := http.StatusNotFound
	w.WriteHeader(httpStatus)

	response := FailedResponse{
		Error: Message{
			Message: "Not found",
		},
		Meta: Meta{
			HTTPStatus: httpStatus,
		},
	}

	_ = json.NewEncoder(w).Encode(response)
}

// Unprocessable Entity 422
func respondWithUnprocessableEntity(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	httpStatus := http.StatusUnprocessableEntity
	w.WriteHeader(httpStatus)

	response := FailedResponse{
		Error: Message{
			Message: "Unprocessable entity",
		},
		Meta: Meta{
			HTTPStatus: httpStatus,
		},
	}

	_ = json.NewEncoder(w).Encode(response)
}
