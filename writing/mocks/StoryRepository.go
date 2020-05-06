// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	entity "gitlab.com/ianadiwibowo/kisakita/entity"
)

// StoryRepository is an autogenerated mock type for the StoryRepository type
type StoryRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: newStory
func (_m *StoryRepository) Create(newStory *entity.Story) error {
	ret := _m.Called(newStory)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.Story) error); ok {
		r0 = rf(newStory)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: existingStory
func (_m *StoryRepository) Delete(existingStory *entity.Story) error {
	ret := _m.Called(existingStory)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.Story) error); ok {
		r0 = rf(existingStory)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get10LatestStories provides a mock function with given fields:
func (_m *StoryRepository) Get10LatestStories() ([]*entity.Story, error) {
	ret := _m.Called()

	var r0 []*entity.Story
	if rf, ok := ret.Get(0).(func() []*entity.Story); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Story)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByAuthorID provides a mock function with given fields: writerID
func (_m *StoryRepository) GetByAuthorID(writerID int) ([]*entity.Story, error) {
	ret := _m.Called(writerID)

	var r0 []*entity.Story
	if rf, ok := ret.Get(0).(func(int) []*entity.Story); ok {
		r0 = rf(writerID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Story)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(writerID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: storyID
func (_m *StoryRepository) GetByID(storyID int) (*entity.Story, error) {
	ret := _m.Called(storyID)

	var r0 *entity.Story
	if rf, ok := ret.Get(0).(func(int) *entity.Story); ok {
		r0 = rf(storyID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Story)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(storyID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: existingStory
func (_m *StoryRepository) Update(existingStory *entity.Story) error {
	ret := _m.Called(existingStory)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.Story) error); ok {
		r0 = rf(existingStory)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}