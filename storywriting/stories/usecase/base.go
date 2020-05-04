package usecase

import (
	"gitlab.com/ianadiwibowo/kisakita/storywriting"
	"gitlab.com/ianadiwibowo/kisakita/storywriting/stories/repository"
)

type StoryUsecase struct {
	StoryRepository storywriting.StoryRepository
}

// NewStoryUsecase initializes a new StoryUsecase instance
func NewStoryUsecase() *StoryUsecase {
	return &StoryUsecase{
		StoryRepository: repository.NewStoryRepository(),
	}
}
