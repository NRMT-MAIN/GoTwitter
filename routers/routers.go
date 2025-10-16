package routers

import (
	"GoTwitter/m/controllers"

	"github.com/go-chi/chi/v5"
)

type Router interface {
	Register(r chi.Router)
}

func SetupRouter(userRouter Router) *chi.Mux {
	chiRouter := chi.NewRouter()

	chiRouter.Get("/ping" , controllers.PingHandeler)

	userRouter.Register(chiRouter)

	return chiRouter
}
