package main

import (
	"fmt"
	"net/http"

	"github.com/glauber-silva/farmers-markets/internal/database"
	transportHTTP "github.com/glauber-silva/farmers-markets/internal/transport/http"
)

// App - A struct which contains things related to database
type App struct {
}

// Run - set up the application
func (app *App) Run() error {
	fmt.Println("Setting up the APP")

	var err error
	_, err = database.NewDatabase()
	if err != nil {
		return err
	}

	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()
	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
		return err
	}
	return nil
}

func main() {
	fmt.Println("Farmers Markets Setup Project")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up the server")
		fmt.Println(err)
	}
}
