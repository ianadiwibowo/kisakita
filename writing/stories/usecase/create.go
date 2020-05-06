package usecase

import (
	"gitlab.com/ianadiwibowo/kisakita/entity"
)

// Create saves the newStory
func (u *StoryUsecase) Create(newStory *entity.Story) (err error) {
	err = u.StoryRepository.Create(newStory)
	if err != nil {
		return err
	}
	return nil
}
