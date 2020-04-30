package usecases_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
	"gitlab.com/ianadiwibowo/kisakita-stories/domains/story/mocks"
	"gitlab.com/ianadiwibowo/kisakita-stories/domains/story/usecases"
	"gitlab.com/ianadiwibowo/kisakita-stories/entity"
)

type StoryUsecasesTestSuite struct {
	suite.Suite
	repo    *mocks.StoryRepositories
	usecase *usecases.StoryUsecase
}

func (s *StoryUsecasesTestSuite) SetupTest() {
	s.repo = new(mocks.StoryRepositories)
	s.usecase = usecases.NewStoryUsecase(s.repo)
}

func TestStoryUsecases(t *testing.T) {
	suite.Run(t, new(StoryUsecasesTestSuite))
}

func (s *StoryUsecasesTestSuite) TestCreate() {
	newStory := &entity.Story{}
	s.repo.On("Create", newStory).Return(nil)

	err := s.usecase.Create(newStory)

	s.Nil(err)
	s.repo.AssertCalled(s.T(), "Create", newStory)
}

func (s *StoryUsecasesTestSuite) TestCreate_ButErrorHappened() {
	newStory := &entity.Story{}
	s.repo.On("Create", newStory).Return(errors.New("Some error"))

	err := s.usecase.Create(newStory)

	s.NotNil(err)
	s.repo.AssertCalled(s.T(), "Create", newStory)
}

func (s *StoryUsecasesTestSuite) TestUpdate() {
	newStory := &entity.Story{}
	s.repo.On("Update", newStory).Return(nil)

	err := s.usecase.Update(newStory)

	s.Nil(err)
	s.repo.AssertCalled(s.T(), "Update", newStory)
}

func (s *StoryUsecasesTestSuite) TestUpdate_ButErrorHappened() {
	newStory := &entity.Story{}
	s.repo.On("Update", newStory).Return(errors.New("Some error"))

	err := s.usecase.Update(newStory)

	s.NotNil(err)
	s.repo.AssertCalled(s.T(), "Update", newStory)
}
