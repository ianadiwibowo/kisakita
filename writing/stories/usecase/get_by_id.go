package usecase

import "gitlab.com/ianadiwibowo/kisakita/entity"

// GetByID retrieves a single story by the storyID
func (u *StoryUsecase) GetByID(storyID int) (story *entity.Story, err error) {
	story, err = u.StoryRepo.GetByID(storyID)
	if err != nil {
		return nil, err
	}
	return story, nil
}
