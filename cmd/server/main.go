package main

import (
	"io"
	"net/http"
	"os"

	"github.com/glauber-silva/farmers-markets/internal/database"
	"github.com/glauber-silva/farmers-markets/internal/markets"
	transportHTTP "github.com/glauber-silva/farmers-markets/internal/transport/http"
	log "github.com/sirupsen/logrus"
)

// App - A struct which contains things related to database
type App struct {
	Name string
}

// Run - set up the application
func (app *App) Run() error {
	log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(
		log.Fields{
			"AppName": app.Name,
		}).Info("Setting up the APP")

	var err error
	db, err := database.NewDatabase()
	if err != nil {
		log.Error(err)
		return err
	}

	err = database.MigrateDB(db)
	if err != nil {
		log.Error(err)
		return err
	}

	marketService := markets.NewService(db)

	handler := transportHTTP.NewHandler(marketService)
	handler.SetupRoutes()
	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		log.Error("Failed to set up server")
		return err
	}
	return nil
}

func main() {
	app := App{
		Name: "Farmers Markets API",
	}

	f, err := os.OpenFile("log.txt", os.O_WRONLY | os.O_APPEND | os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}

	m := io.MultiWriter(os.Stdout, f)
	log.SetOutput(m)

	if err := app.Run(); err != nil {
		log.Error("Error starting up the server")
		log.Fatal(err)
	}
}
