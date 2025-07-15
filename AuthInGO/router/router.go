package router

import (
	"AuthInGO/controllers"


	"github.com/go-chi/chi/v5"
)
type Router interface{
	Register(r *chi.Router)
}


func SetupRouter() *chi.Mux {

	chirouter := chi.NewRouter()
	chirouter.Get("/ping", controllers.PingHandler)


	UserRouter.Register(chirouter)


	return chirouter


}