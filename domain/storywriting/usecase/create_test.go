package usecase_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
	"gitlab.com/ianadiwibowo/kisakita-stories/domain/storywriting/mocks"
	"gitlab.com/ianadiwibowo/kisakita-stories/domain/storywriting/usecase"
	"gitlab.com/ianadiwibowo/kisakita-stories/entity"
)

type StoryUsecaseTestSuite struct {
	suite.Suite
	repo    *mocks.StoryRepository
	usecase *usecase.StoryUsecase
}

func (s *StoryUsecaseTestSuite) SetupTest() {
	s.repo = new(mocks.StoryRepository)
	s.usecase = usecase.NewStoryUsecase(s.repo)
}

func TestStoryUsecase(t *testing.T) {
	suite.Run(t, new(StoryUsecaseTestSuite))
}

func (s *StoryUsecaseTestSuite) TestCreate() {
	newStory := &entity.Story{}
	s.repo.On("Create", newStory).Return(nil)

	err := s.usecase.Create(newStory)

	s.Nil(err)
	s.repo.AssertCalled(s.T(), "Create", newStory)
}

func (s *StoryUsecaseTestSuite) TestCreate_ButErrorHappened() {
	newStory := &entity.Story{}
	s.repo.On("Create", newStory).Return(errors.New("Some error"))

	err := s.usecase.Create(newStory)

	s.NotNil(err)
	s.repo.AssertCalled(s.T(), "Create", newStory)
}

func (s *StoryUsecaseTestSuite) TestUpdate() {
	newStory := &entity.Story{}
	s.repo.On("Update", newStory).Return(nil)

	err := s.usecase.Update(newStory)

	s.Nil(err)
	s.repo.AssertCalled(s.T(), "Update", newStory)
}

func (s *StoryUsecaseTestSuite) TestUpdate_ButErrorHappened() {
	newStory := &entity.Story{}
	s.repo.On("Update", newStory).Return(errors.New("Some error"))

	err := s.usecase.Update(newStory)

	s.NotNil(err)
	s.repo.AssertCalled(s.T(), "Update", newStory)
}
