package api

import (
	"net/http"
	"os"

	"github.com/goji/httpauth"
	"github.com/gorilla/mux"
)

// App is the collection of all hard dependencies
type App struct {
	Router *mux.Router
}

// New initializes the app
func NewApp() *App {
	return &App{
		Router: mux.NewRouter().StrictSlash(true),
	}
}

// SetupRoutes initializes the API endpoint routes
func (app *App) SetupRoutes() {
	app.useBasicAuth()

	c := NewContainer()
	h := c.GetStoryHandler()

	app.route("GET", "/stories/{id}", h.Get)
	app.route("POST", "/stories", h.Create)
	app.route("PATCH", "/stories/{id}", h.Update)
	app.route("DELETE", "/stories/{id}", h.Delete)
	app.route("GET", "/stories/by-authors/{id}", h.GetByAuthorID)
	app.route("GET", "/stories/ten-latests", h.Get10LatestStories)
}

// useBasicAuth add API endpoint authentication using basic auth
func (app *App) useBasicAuth() {
	app.Router.Use(httpauth.SimpleBasicAuth(
		os.Getenv("BASIC_USERNAME"),
		os.Getenv("BASIC_PASSWORD"),
	))
}

// route registers an API endpoint to a handler
func (app *App) route(
	method string,
	path string,
	handler func(http.ResponseWriter, *http.Request),
) {
	app.Router.Methods(method).Path(path).HandlerFunc(handler)
}
