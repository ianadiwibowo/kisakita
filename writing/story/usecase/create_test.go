package usecase_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
	"gitlab.com/ianadiwibowo/kisakita/entity"
	"gitlab.com/ianadiwibowo/kisakita/writing/mocks"
	"gitlab.com/ianadiwibowo/kisakita/writing/story/usecase"
)

type StoryUsecaseTestSuite struct {
	suite.Suite
	repo    *mocks.StoryRepo
	usecase *usecase.StoryUsecase
}

func (s *StoryUsecaseTestSuite) SetupTest() {
	s.repo = new(mocks.StoryRepo)
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

func (s *StoryUsecaseTestSuite) TestCreate_WhenErrorHappened() {
	newStory := &entity.Story{}
	s.repo.On("Create", newStory).Return(errors.New("Some error"))

	err := s.usecase.Create(newStory)

	s.NotNil(err)
	s.repo.AssertCalled(s.T(), "Create", newStory)
}
