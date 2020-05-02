package handler

import (
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

// StoriesHandler is the controller for stories scope
type StoriesHandler struct {
}

// NewStoriesHandler initializes fresh handler
func NewStoriesHandler() *StoriesHandler {
	return &StoriesHandler{}
}

// getQueryString extracts request payload from HTTP request
func getQueryString(r *http.Request) map[string]string {
	return mux.Vars(r)
}

// getParams extracts HTTP POST request payload to a struct
func getParams(r *http.Request, p interface{}) {
	_ = json.NewDecoder(r.Body).Decode(p)
}

// validateParams checks params compliance to validation rules
func validateParams(p interface{}) (errors []error) {
	err := validator.New().Struct(p)
	if err == nil {
		return nil
	}

	if _, ok := err.(*validator.InvalidValidationError); ok {
		errors = append(errors, &InvalidValidationError{})
		return errors
	}

	for _, err := range err.(validator.ValidationErrors) {
		errors = append(errors, &InvalidAttributeError{
			Field:        err.Field(),
			ValidatorTag: err.Tag(),
		})
	}
	return errors
}

// SuccessResponse is the standard JSON API success data-meta response
type SuccessResponse struct {
	Data interface{} `json:"data"`
	Meta interface{} `json:"meta"`
}

// ErrorResponse is the standard JSON API failed errors-meta response
type ErrorResponse struct {
	Errors []ErrorDetail `json:"errors"`
	Meta   interface{}   `json:"meta"`
}

// ErrorDetail represents individual JSON API error detail
type ErrorDetail struct {
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

// respondOK sends success response with HTTP status 200
func respondOK(w http.ResponseWriter, data, meta interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := SuccessResponse{
		Data: data,
		Meta: meta,
	}
	_ = json.NewEncoder(w).Encode(response)
}

// respondCreated sends success response with HTTP status 201
func respondCreated(w http.ResponseWriter, data, meta interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := SuccessResponse{
		Data: data,
		Meta: meta,
	}
	_ = json.NewEncoder(w).Encode(response)
}

// TODO: So many duplicate codes with these error responses
// respondBadRequest sends error response with HTTP status 400
func respondBadRequest(w http.ResponseWriter, errors []error, meta interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	response := ErrorResponse{
		Errors: convertToErrorDetails(errors),
		Meta:   meta,
	}
	_ = json.NewEncoder(w).Encode(response)
}

// Not Found 404
func respondWithNotFound(w http.ResponseWriter, errors []error, meta interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	response := ErrorResponse{
		Errors: convertToErrorDetails(errors),
		Meta:   meta,
	}
	_ = json.NewEncoder(w).Encode(response)
}

// Unprocessable Entity 422
func respondWithUnprocessableEntity(w http.ResponseWriter, errors []error, meta interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnprocessableEntity)
	response := ErrorResponse{
		Errors: convertToErrorDetails(errors),
		Meta:   meta,
	}
	_ = json.NewEncoder(w).Encode(response)
}

// convertToErrorDetails converts standard error to serialized error object
func convertToErrorDetails(errors []error) (errorObjects []ErrorDetail) {
	for _, err := range errors {
		errorObjects = append(errorObjects, ErrorDetail{
			Title:  reflect.TypeOf(err).Elem().Name(),
			Detail: err.Error(),
		})
	}
	return errorObjects
}
