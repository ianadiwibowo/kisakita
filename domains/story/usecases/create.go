package usecases

import (
	"github.com/ianadiwibowo/kisakita/domains/story"
	"github.com/ianadiwibowo/kisakita/entities"
)

type StoryUsecase struct {
	repo story.StoryRepositories
}

// NewStoryUsecase initializes a new StoryUsecase instance
func NewStoryUsecase(repo story.StoryRepositories) *StoryUsecase {
	return &StoryUsecase{
		repo: repo,
	}
}

// Create saves the newStory
func (u *StoryUsecase) Create(newStory *entities.Story) error {
	err := u.repo.Create(newStory)
	if err != nil {
		return err
	}

	return nil
}

// Update edits the existingStory
func (u *StoryUsecase) Update(existingStory *entities.Story) error {
	err := u.repo.Update(existingStory)
	if err != nil {
		return err
	}

	return nil
}