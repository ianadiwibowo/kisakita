package repository

import (
	"gitlab.com/ianadiwibowo/kisakita/entity"
)

// GetByID retrieves a single story by the storyID
func (r *StoryRepository) GetByID(storyID int) (
	*entity.Story,
	error,
) {
	story := entity.Story{}
	err := r.db.First(&story, storyID).Error
	if err != nil {
		return nil, err
	}
	return &story, nil
}
