package repository

import (
	"gitlab.com/ianadiwibowo/kisakita/entity"
)

// Update edits the updatedStory
func (r *StoryRepository) Update(
	updatedStory *entity.Story,
) (err error) {
	// Retrieve first
	var story entity.Story
	err = r.db.First(&story, updatedStory.ID).Error
	if err != nil {
		return err
	}

	// Update
	story = *updatedStory
	err = r.db.Save(&story).Error
	if err != nil {
		return err
	}

	return nil
}
