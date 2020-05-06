package usecase

import (
	"gitlab.com/ianadiwibowo/kisakita/writing"
	"gitlab.com/ianadiwibowo/kisakita/writing/stories/repository"
)

type StoryUsecase struct {
	StoryRepository writing.StoryRepository
}

// NewStoryUsecase initializes a new StoryUsecase instance
func NewStoryUsecase() *StoryUsecase {
	return &StoryUsecase{
		StoryRepository: repository.NewStoryRepository(),
	}
}
