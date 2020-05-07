package usecase

import "gitlab.com/ianadiwibowo/kisakita/entity"

// Get retrieves a single story by the storyID
func (u *StoryUsecase) Get(storyID int) (*entity.Story, error) {
	story, err := u.StoryRepo.Get(storyID)
	if err != nil {
		return nil, err
	}
	return story, nil
}
