package handler

import (
	"net/http"

	"gitlab.com/ianadiwibowo/kisakita/entity"
)

// Create handles `POST /stories`
func (h *StoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	// Validate request payload
	var params createParams
	getParams(r, &params)
	errs := validateParams(params)
	if errs != nil {
		respondBadRequest(w, errs, nil)
		return
	}

	// Process according to business rules
	err := h.StoryUsecase.Create(
		&entity.Story{
			Title:    params.Title,
			Synopsis: params.Synopsis,
		},
	)
	if err != nil {
		respondUnprocessableEntity(w, []error{err}, nil)
		return
	}

	respondCreated(w)
}

type createParams struct {
	Title    string `validate:"required"`
	Synopsis string `validate:"required"`
}
