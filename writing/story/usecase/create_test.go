package usecase_test

import (
	"errors"

	"gitlab.com/ianadiwibowo/kisakita/entity"
)

func (s *StoryUsecaseTestSuite) TestCreate() {
	newStory := &entity.Story{}
	s.repo.On("Create", newStory).Return(nil)

	err := s.usecase.Create(newStory)

	s.Nil(err)
	s.repo.AssertCalled(s.T(), "Create", newStory)
}

func (s *StoryUsecaseTestSuite) TestCreate_WhenErrorHappened() {
	newStory := &entity.Story{}
	s.repo.On("Create", newStory).Return(errors.New("Error"))

	err := s.usecase.Create(newStory)

	s.NotNil(err)
	s.repo.AssertCalled(s.T(), "Create", newStory)
}
