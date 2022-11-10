package main

import (
	"fmt"

	"github.com/arturfil/gorilla_soccer/internal/db"
	"github.com/arturfil/gorilla_soccer/internal/field"
	"github.com/arturfil/gorilla_soccer/internal/group"
	transportHttp "github.com/arturfil/gorilla_soccer/internal/transport/http"
)

// Run - is going to be responsible for the
// instantiation and startup of our go application
func Run() error {
	fmt.Println("Starting up our application")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to db")
		return err
	}
	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate database")
		return err
	}

	// new services will be added here
	groupService := group.NewService(db)
	fieldService := field.NewService(db)

	// I will pass all the services here in the transporhttp obj
	groupHttpHandler := transportHttp.NewHandler(groupService, fieldService)

	if err := groupHttpHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func main() {
	fmt.Println("API Main Function")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
