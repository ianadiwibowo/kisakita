package main

import (
	"net/http"
	"os"

	"github.com/goji/httpauth"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"

	"gitlab.com/ianadiwibowo/kisakita/storywriting/stories/handler"
)

// App is the collection of required dependencies
type App struct {
	Router *mux.Router
}

func main() {
	loadEnvironmentVariables()

	app := NewApp()
	app.SetupRoutes()

	_ = http.ListenAndServe(":8080", app.Router)
}

// Load system configurations, get them via: os.Getenv("{ENV_KEY_NAME}")
func loadEnvironmentVariables() {
	if os.Getenv("APP_ENV") != "production" {
		err := gotenv.Load()
		if err != nil {
			panic(err)
		}
	}
}

// New initializes the app
func NewApp() *App {
	return &App{
		Router: mux.NewRouter().StrictSlash(true),
	}
}

// SetupRoutes initializes the API endpoint routes
func (app *App) SetupRoutes() {
	// Authenticate using basic auth, for now
	app.Router.Use(httpauth.SimpleBasicAuth(
		os.Getenv("BASIC_USERNAME"),
		os.Getenv("BASIC_PASSWORD"),
	))

	storywritingHandler := handler.NewStoriesHandler()

	app.Router.Methods("GET").Path("/stories/{id}").HandlerFunc(
		storywritingHandler.GetByID,
	)
	app.Router.Methods("POST").Path("/stories").HandlerFunc(
		storywritingHandler.Create,
	)
	app.Router.Methods("PATCH").Path("/stories/{id}").HandlerFunc(
		storywritingHandler.Update,
	)
	app.Router.Methods("DELETE").Path("/stories/{id}").HandlerFunc(
		storywritingHandler.Delete,
	)
	app.Router.Methods("GET").Path("/stories/by-authors/{id}").HandlerFunc(
		storywritingHandler.GetByAuthorID,
	)
	app.Router.Methods("GET").Path("/stories/ten-latests").HandlerFunc(
		storywritingHandler.Get10LatestStories,
	)
}
