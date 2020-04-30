package storywriting

import "gitlab.com/ianadiwibowo/kisakita-stories/entity"

type StoryRepository interface {
	GetByID(storyID int) (story *entity.Story, err error)
	Create(newStory *entity.Story) (err error)
	Update(existingStory *entity.Story) (err error)
	Delete(existingStory *entity.Story) (err error)

	GetByAuthorID(writerID int) (stories []*entity.Story, err error)
	Get10LatestStories() (stories []*entity.Story, err error)
}
