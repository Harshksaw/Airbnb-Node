package app

import (
	"AuthInGO/config"
	db "AuthInGO/db/repositories"
	"AuthInGO/router"
	"AuthInGO/services"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type Application struct {
	Config Config
	Store db.Storage
}

type Config struct {
	Addr string	

}


func NewConfig() Config {

	port := config.GetString("PORT", ":8080")

	return Config{
		Addr: port,
	}
}
//constructor in go
func NewApplication(cfg Config) *Application {
	return &Application{
		Config: cfg,
		Store: *db.NewStorage(),
	}
}

func (app *Application) Run()error{
	// Initialize the router

	ur := db.NewUserRepository()
	us := services.NewUserService(ur)
	uc := controllers.NewUserController(us)
	userRouter := router.NewUserRouter(uc)

	server := &http.Server{
		Addr: app.Config.Addr, //port
		Handler: router.SetupRouter(userRouter), //setup a chi router and put it here
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout: 10 * time.Second,

		ErrorLog: log.New(os.Stdout, "", 0),
		

		
	}

	fmt.Println("Server is running on", app.Config.Addr)
	return server.ListenAndServe()

}
