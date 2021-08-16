package main

import "fmt"

// App - A struct which contains things related to database
type App struct {
}

// Run - set up the application
func (app *App) Run() error {
	fmt.Println("Setting up the APP")
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
