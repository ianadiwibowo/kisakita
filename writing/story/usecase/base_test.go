package usecase_test

import (
	"testing"

	"github.com/ianadiwibowo/kisakita/writing/mocks"
	"github.com/ianadiwibowo/kisakita/writing/story/usecase"
	"github.com/stretchr/testify/suite"
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
