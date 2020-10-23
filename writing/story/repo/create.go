package repo

import (
	"time"

	"github.com/ianadiwibowo/kisakita/entity"
)

// Create saves the newStory
func (r *StoryRepo) Create(newStory *entity.Story) error {
	createTimestamp(newStory)

	err := r.db.Create(newStory).Error
	if err != nil {
		return err
	}
	return nil
}

func createTimestamp(story *entity.Story) {
	story.CreatedAt = time.Now()
	story.UpdatedAt = story.CreatedAt
}
