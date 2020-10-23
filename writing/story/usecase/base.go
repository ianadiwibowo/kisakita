package usecase

import (
	"github.com/ianadiwibowo/kisakita/writing"
)

type StoryUsecase struct {
	StoryRepo writing.StoryRepo
}

// NewStoryUsecase initializes a new StoryUsecase instance
func NewStoryUsecase(storyRepo writing.StoryRepo) *StoryUsecase {
	return &StoryUsecase{
		StoryRepo: storyRepo,
	}
}
