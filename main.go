package main

import (
	"GoTwitter/m/app"
)

func main() {
	cfg := app.NewConfig()
	appn := app.NewApplication(*cfg)

	appn.Run()
}
