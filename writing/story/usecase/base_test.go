package usecase_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
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
