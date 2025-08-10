package main

import (
	"AuthInGo/app"

	config  "AuthInGo/config/env" 
	

)

func main() {
	config.Load() // Load environment variables
	cfg := app.NewConfig()
	app := app.NewApplication(cfg)	
	
	app.Run()

}
