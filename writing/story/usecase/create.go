package usecase

import (
	"github.com/ianadiwibowo/kisakita/entity"
)

// Create saves the newStory
func (u *StoryUsecase) Create(newStory *entity.Story) error {
	err := u.StoryRepo.Create(newStory)
	if err != nil {
		return err
	}
	return nil
}
