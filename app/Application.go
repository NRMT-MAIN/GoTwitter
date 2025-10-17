package app

import (
	"GoTwitter/m/config/env"
	"GoTwitter/m/db"
	"fmt"
	"net/http"
	"time"
)

type Config struct {
	Addrs string
}

type Application struct {
	Config Config
	Store db.Storage
}

func NewConfig() *Config {
	port := env.GetString("PORT" , ":8000")

	return &Config{
		Addrs: port,
	}
}

func NewApplication(cfg Config) *Application {
	return &Application{
		Config : cfg ,
		Store : *db.NewStorage() ,
	}
}

func (app *Application) Run() error {
	server := &http.Server{
		Addr: app.Config.Addrs,
		Handler: nil,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Server running on port : " , app.Config.Addrs)
	return server.ListenAndServe()
}
