package repo

import (
	"github.com/ianadiwibowo/kisakita/entity"
	"github.com/jinzhu/gorm"
)

// Get retrieves a single story by the storyID
func (r *StoryRepo) Get(storyID int) (*entity.Story, error) {
	story := entity.Story{}
	err := r.db.First(&story, storyID).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, &RecordNotFoundError{Entity: story, ID: storyID}
	}
	if err != nil {
		return nil, err
	}
	return &story, nil
}
