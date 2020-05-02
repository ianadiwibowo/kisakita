package usecase

import (
	"gitlab.com/ianadiwibowo/kisakita/entity"
	"gitlab.com/ianadiwibowo/kisakita/storywriting"
)

type StoryUsecase struct {
	repo storywriting.StoryRepository
}

// NewStoryUsecase initializes a new StoryUsecase instance
func NewStoryUsecase(repo storywriting.StoryRepository) *StoryUsecase {
	return &StoryUsecase{
		repo: repo,
	}
}

// Create saves the newStory
func (u *StoryUsecase) Create(newStory *entity.Story) error {
	err := u.repo.Create(newStory)
	if err != nil {
		return err
	}

	return nil
}

// Update edits the existingStory
func (u *StoryUsecase) Update(existingStory *entity.Story) error {
	err := u.repo.Update(existingStory)
	if err != nil {
		return err
	}

	return nil
}

// Delete removes the existingStory
func (u *StoryUsecase) Delete(existingStory *entity.Story) error {
	return nil
}

// GetByAuthorID retrieves all stories that the writerID involved
func (u *StoryUsecase) GetByAuthorID(writerID int) (stories []*entity.Story, err error) {
	return nil, nil
}

// Get10LatestStories retrieve 10 newest stories ordered by latest update
func Get10LatestStories() (stories []*entity.Story, err error) {
	return nil, nil
}
