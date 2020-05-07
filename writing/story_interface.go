package writing

import (
	"net/http"

	"gitlab.com/ianadiwibowo/kisakita/entity"
)

type StoryHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	GetByAuthorID(w http.ResponseWriter, r *http.Request)
	Get10LatestStories(w http.ResponseWriter, r *http.Request)
}

type StoryUsecase interface {
	Get(storyID int) (*entity.Story, error)
	Create(newStory *entity.Story) error
	Update(existingStory *entity.Story) error
	Delete(existingStory *entity.Story) error
	GetByAuthorID(writerID int) ([]*entity.Story, error)
	Get10LatestStories() ([]*entity.Story, error)
}

type StoryRepo interface {
	Get(storyID int) (*entity.Story, error)
	Create(newStory *entity.Story) error
	Update(existingStory *entity.Story) error
	Delete(existingStory *entity.Story) error
	GetByAuthorID(writerID int) ([]*entity.Story, error)
	Get10LatestStories() ([]*entity.Story, error)
}
