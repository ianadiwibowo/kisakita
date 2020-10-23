package api

import (
	"github.com/ianadiwibowo/kisakita/writing/story/handler"
	"github.com/ianadiwibowo/kisakita/writing/story/repo"
	"github.com/ianadiwibowo/kisakita/writing/story/usecase"
)

// Container is the singleton collection of all internal dependencies
type Container struct {
	StoryHandler *handler.StoryHandler
	storyUsecase *usecase.StoryUsecase
	storyRepo    *repo.StoryRepo
}

// NewContainer creates an empty singleton collection
func NewContainer() *Container {
	return &Container{}
}

// GetStoryHandler gives singleton instance of a StoryHandler
func (c *Container) GetStoryHandler() *handler.StoryHandler {
	if c.StoryHandler == nil {
		c.StoryHandler = handler.NewStoryHandler(c.getStoryUsecase())
	}
	return c.StoryHandler
}

// getStoryUsecase gives singleton instance of a StoryUsecase
func (c *Container) getStoryUsecase() *usecase.StoryUsecase {
	if c.storyUsecase == nil {
		c.storyUsecase = usecase.NewStoryUsecase(c.getStoryRepo())
	}
	return c.storyUsecase
}

// getStoryRepo gives singleton instance of a StoryRepo
func (c *Container) getStoryRepo() *repo.StoryRepo {
	if c.storyRepo == nil {
		c.storyRepo = repo.NewStoryRepo()
	}
	return c.storyRepo
}
