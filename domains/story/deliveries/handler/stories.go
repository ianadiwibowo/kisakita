package handler

import (
	"net/http"
)

// StoriesHandler ...
type StoriesHandler struct {
}

// NewStoriesHandler initializes fresh handler
func NewStoriesHandler() *StoriesHandler {
	return &StoriesHandler{}
}

// Show handles GET /stories/{id}
func (h *StoriesHandler) GetByID(w http.ResponseWriter, r *http.Request) {}

// Create handles POST /stories
func (h *StoriesHandler) Create(w http.ResponseWriter, r *http.Request) {}

// Update handles PATCH /stories/{id}
func (h *StoriesHandler) Update(w http.ResponseWriter, r *http.Request) {}

// Delete handles DELETE /stories/{id}
func (h *StoriesHandler) Delete(w http.ResponseWriter, r *http.Request) {}

// GetByAuthorID handles GET /stories/by-authors/{id}
func (h *StoriesHandler) GetByAuthorID(w http.ResponseWriter, r *http.Request) {}

// Get10LatestStories handles GET /stories/ten-latests
func (h *StoriesHandler) Get10LatestStories(w http.ResponseWriter, r *http.Request) {}
