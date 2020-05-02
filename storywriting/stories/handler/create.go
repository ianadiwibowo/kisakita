package handler

import (
	"fmt"
	"net/http"
)

// Create handles POST /stories
func (h *StoriesHandler) Create(w http.ResponseWriter, r *http.Request) {
	// Validate request payload
	var p createParams
	getParams(r, &p)
	errs := validateParams(p)
	if errs != nil {
		respondBadRequest(w, errs, nil)
		return
	}

	// Process according to business rules
	fmt.Println("Process according to business rules")
}

type createParams struct {
	Title    string `validate:"required"`
	Synopsis string `validate:"required"`
}
