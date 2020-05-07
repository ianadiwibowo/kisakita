package repo

import (
	"gitlab.com/ianadiwibowo/kisakita/entity"
)

// GetByAuthorID retrieves all stories that the writerID involved
func (r *StoryRepo) GetByAuthorID(writerID int) ([]*entity.Story, error) {
	return nil, nil
}
