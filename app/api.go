package main

import (
	"net/http"
	"os"

	"github.com/goji/httpauth"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"

	"gitlab.com/ianadiwibowo/kisakita-stories/domains/story/delivery/handler"
)

// App is the collection of required dependencies
type App struct {
	Router *mux.Router
}

func main() {
	loadEnvironmentVariables()

	app := NewApp()
	app.SetupRouter()

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

// SetupRouter initializes the API endpoint routes
func (app *App) SetupRouter() {
	// Authenticate using basic auth, for now
	app.Router.Use(httpauth.SimpleBasicAuth(
		os.Getenv("BASIC_USERNAME"),
		os.Getenv("BASIC_PASSWORD"),
	))

	storiesHandler := handler.NewStoriesHandler()

	app.Router.Methods("GET").Path("/stories/{id}").HandlerFunc(
		storiesHandler.GetByID,
	)
	app.Router.Methods("POST").Path("/stories").HandlerFunc(
		storiesHandler.Create,
	)
	app.Router.Methods("PATCH").Path("/stories/{id}").HandlerFunc(
		storiesHandler.Update,
	)
	app.Router.Methods("DELETE").Path("/stories/{id}").HandlerFunc(
		storiesHandler.Delete,
	)
	app.Router.Methods("GET").Path("/stories/by-authors/{id}").HandlerFunc(
		storiesHandler.GetByAuthorID,
	)
	app.Router.Methods("GET").Path("/stories/ten-latests").HandlerFunc(
		storiesHandler.Get10LatestStories,
	)
}
