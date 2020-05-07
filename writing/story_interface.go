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
	Get(storyID int) (story *entity.Story, err error)
	Create(newStory *entity.Story) (err error)
	Update(existingStory *entity.Story) (err error)
	Delete(existingStory *entity.Story) (err error)
	GetByAuthorID(writerID int) (stories []*entity.Story, err error)
	Get10LatestStories() (stories []*entity.Story, err error)
}

type StoryRepo interface {
	Get(storyID int) (story *entity.Story, err error)
	Create(newStory *entity.Story) (err error)
	Update(existingStory *entity.Story) (err error)
	Delete(existingStory *entity.Story) (err error)
	GetByAuthorID(writerID int) (stories []*entity.Story, err error)
	Get10LatestStories() (stories []*entity.Story, err error)
}
