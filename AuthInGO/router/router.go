package router

import (
	"AuthInGo/controllers"


	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)
type Router interface{
	Register(r *chi.Router)
}


func SetupRouter() *chi.Mux {

	chirouter := chi.NewRouter()

	chirouter.Use(middleware.Logger)

	chirouter.Get("/ping", controllers.PingHandler)


	UserRouter.Register(chirouter)


	return chirouter


}