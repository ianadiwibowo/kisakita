package handler

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Show handles `GET /stories/{id}`
func (h *StoriesHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	// Validate request payload
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if id == 0 || err != nil {
		respondUnprocessableEntity(w, []error{err}, nil)
		return
	}

	// Process according to business rules
	story, err := h.StoryUsecase.GetByID(id)
	if err != nil {
		respondUnprocessableEntity(w, []error{err}, nil)
		return
	}

	respondOK(w, story, nil)
}
