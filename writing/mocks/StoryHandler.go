// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// StoryHandler is an autogenerated mock type for the StoryHandler type
type StoryHandler struct {
	mock.Mock
}

// Create provides a mock function with given fields: w, r
func (_m *StoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// Delete provides a mock function with given fields: w, r
func (_m *StoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// Get provides a mock function with given fields: w, r
func (_m *StoryHandler) Get(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// Get10LatestStories provides a mock function with given fields: w, r
func (_m *StoryHandler) Get10LatestStories(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// GetByAuthorID provides a mock function with given fields: w, r
func (_m *StoryHandler) GetByAuthorID(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// Update provides a mock function with given fields: w, r
func (_m *StoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}
