package main

import (
	"config"
	"github.com/codegangsta/negroni"
	"logger"
	"middleware"
	"model"
	"router"
)

func initModules() {
	model.Init()
}

func main() {

	initModules()
	
	defer model.Db.Close()

	r := router.RouteHandler()
	
	port := config.Port
	logger.Get().SetLogLevel(3)
	logger.Get().Debug("Port is : " + port)
	n := negroni.Classic()
	n.Use(middleware.NewResponse())
	n.Use(middleware.NewUserAuth())
	n.UseHandler(r)
	n.Run(port)
}