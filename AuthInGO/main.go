package main

import (
	"AuthInGO/app"

)

func main() {
	cfg := app.NewConfig()
	app := app.NewApplication(cfg)	
	
	app.Run()

}
