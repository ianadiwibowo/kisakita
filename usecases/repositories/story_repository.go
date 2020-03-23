package repositories

import "github.com/ianadiwibowo/kisakita/entities"

type StoryRepository interface {
	Retrieve(storyID int) (story *entities.Story, retrievalError error)
	Create(newStory *entities.Story) (creationError error)
	Update(existingStory *entities.Story) (updateError error)
	Delete(existingStory *entities.Story) (deletionError error)
	Retrieve10LatestStories() (stories []*entities.Story, retrievalError error)
}
