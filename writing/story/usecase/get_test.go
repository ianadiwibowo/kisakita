package usecase_test

import (
	"errors"

	"gitlab.com/ianadiwibowo/kisakita/entity"
)

func (s *StoryUsecaseTestSuite) TestGet_WhenHappyFlow() {
	storyID := 7
	story := &entity.Story{ID: storyID}
	s.repo.On("Get", storyID).Return(story, nil)

	result, err := s.usecase.Get(storyID)

	s.Nil(err)
	s.NotNil(result)
	s.Equal(result.ID, storyID)
	s.repo.AssertCalled(s.T(), "Get", storyID)
}

func (s *StoryUsecaseTestSuite) TestGet_WhenErrorHappened() {
	storyID := 7
	s.repo.On("Get", storyID).Return(nil, errors.New("Error"))

	result, err := s.usecase.Get(storyID)

	s.NotNil(err)
	s.Nil(result)
	s.repo.AssertCalled(s.T(), "Get", storyID)
}
