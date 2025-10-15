package routers

import (
	"GoTwitter/m/controllers"

	"github.com/go-chi/chi/v5"
)


func SetupRouter() *chi.Mux {
	chiRouter := chi.NewRouter()

	chiRouter.Get("/ping" , controllers.PingHandeler)
}
