package router

import (
	"AuthInGo/controllers"
	"AuthInGo/middlewares"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

)
type Router interface{
	Register(r chi.Router)
}


func SetupRouter(userRouter Router) *chi.Mux {

	chirouter := chi.NewRouter()

	chirouter.Use(middleware.Logger)
	chirouter.Use(middlewares.RateLimitMiddleware)

	chirouter.Get("/ping", controllers.PingHandler)

	userRouter.Register(chirouter)

	return chirouter
}