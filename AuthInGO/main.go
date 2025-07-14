package main

import (
	"AuthInGO/app"

	config  "AuthInGO/config" 

)

func main() {
	config.Load() // Load environment variables
	cfg := app.NewConfig()
	app := app.NewApplication(cfg)	
	
	app.Run()

}
