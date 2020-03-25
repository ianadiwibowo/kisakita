package story

import "github.com/ianadiwibowo/kisakita/entities"

type StoryRepositories interface {
	GetByID(storyID int) (story *entities.Story, err error)
	Create(newStory *entities.Story) (err error)
	Update(existingStory *entities.Story) (err error)
	Delete(existingStory *entities.Story) (err error)

	GetByAuthorID(writerID int) (stories []*entities.Story, err error)
	Get10LatestStories() (stories []*entities.Story, err error)
}
