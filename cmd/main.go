package main

import (
	"net/http"
	"os"

	"github.com/ianadiwibowo/kisakita/cmd/api"
	"github.com/subosito/gotenv"
)

func main() {
	loadEnvironmentVariables()

	app := api.NewApp()
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
