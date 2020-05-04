package repository

import (
	"gitlab.com/ianadiwibowo/kisakita/entity"
)

// GetByAuthorID retrieves all stories that the writerID involved
func (r *StoryRepository) GetByAuthorID(writerID int) (stories []*entity.Story, err error) {
	return nil, nil
}
