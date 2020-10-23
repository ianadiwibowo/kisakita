package usecase

import (
	"github.com/ianadiwibowo/kisakita/entity"
)

// GetByAuthorID retrieves all stories that the writerID involved
func (u *StoryUsecase) GetByAuthorID(writerID int) ([]*entity.Story, error) {
	return nil, nil
}
