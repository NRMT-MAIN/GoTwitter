package main

import (
	"GoTwitter/m/app"
	cfgdb "GoTwitter/m/config/db"
)

func main() {
	cfg := app.NewConfig()
	appn := app.NewApplication(*cfg)

	cfgdb.SetupDB()

	appn.Run()
}
