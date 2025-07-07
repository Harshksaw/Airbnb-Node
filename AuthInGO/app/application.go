package app

import ("net/http" 
"time" 
"fmt"
 "os"
  "log")

type Application struct {
	Config Config
}

type Config struct {
	Addr string	

}


func NewConfig(addr string) Config {
	return Config{
		Addr: addr,
	}
}
//constructor in go
func NewApplication(cfg Config) *Application {
	return &Application{
		Config: cfg,
	}
}

func (app *Application) Run()error{
	server := &http.Server{
		Addr: app.Config.Addr, //port
		Handler: nil, //setup a chi router and put it here
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout: 10 * time.Second,

		ErrorLog: log.New(os.Stdout, "", 0),
		

		
	}

	fmt.Println("Server is running on", app.Config.Addr)
	return server.ListenAndServe()

}
