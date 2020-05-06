package usecase

import (
	"gitlab.com/ianadiwibowo/kisakita/writing"
	"gitlab.com/ianadiwibowo/kisakita/writing/stories/repo"
)

type StoryUsecase struct {
	StoryRepo writing.StoryRepo
}

// NewStoryUsecase initializes a new StoryUsecase instance
func NewStoryUsecase() *StoryUsecase {
	return &StoryUsecase{
		StoryRepo: repo.NewStoryRepo(),
	}
}
