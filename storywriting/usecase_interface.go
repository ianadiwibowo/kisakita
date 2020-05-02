package storywriting

import "gitlab.com/ianadiwibowo/kisakita/entity"

type StoryUsecase interface {
	GetByID(storyID int) (story *entity.Story, err error)
	Create(newStory *entity.Story) (err error)
	Update(existingStory *entity.Story) (err error)
	Delete(existingStory *entity.Story) (err error)

	GetByAuthorID(writerID int) (stories []*entity.Story, err error)
	Get10LatestStories() (stories []*entity.Story, err error)
}
