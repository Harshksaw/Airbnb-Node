package app

import (
	dbConfig "AuthInGo/config/db"
	envAlias "AuthInGo/config/env"
	dbAlias "AuthInGo/db/repositories"
	"AuthInGo/router"
	"AuthInGo/services"

	"fmt"
	"net/http"
	"time"
)

// Config holds the configuration for the server.
type Config struct {
	Addr string // PORT
}

type Application struct {
	Config Config
}

// Constructor for Config
func NewConfig() Config {

	port := envAlias.GetString("PORT", ":8080")

	return Config{
		Addr: port,
	}
}

// Constructor for Application
func NewApplication(cfg Config) *Application {
	return &Application{
		Config: cfg,
	}
}

func (app *Application) Run() error {

	dbInstance, err := dbConfig.SetupDB()
	if err != nil {
		return fmt.Errorf("failed to setup database: %w", err)
	}

	repo := dbAlias.NewUserRepository(dbInstance)
	services.NewUserService(repo)

	handler := router.SetupRouter()

	server := &http.Server{
		Addr:         app.Config.Addr,
		Handler:      handler,
		ReadTimeout:  10 * time.Second, // Set read timeout to 10 seconds
		WriteTimeout: 10 * time.Second, // Set write timeout to 10 seconds
	}

	fmt.Println("Starting server on", app.Config.Addr)

	return server.ListenAndServe()
}
