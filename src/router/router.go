package router

import (
	"github.com/julienschmidt/httprouter"
    "controller"
    "controller/objectdetail"
    "controller/objectlike"
    "controller/objectcomment"
)

func RouteHandler() *httprouter.Router {
	m := httprouter.New()

	m.GET("/", controller.MainHandler)

	// Object details
	m.GET("/objectdetail/read", objectdetail.Read)
	// Object like calls
	m.GET("/objectlike/create", objectlike.Create)
	m.GET("/objectlike/read", objectlike.Read)
	m.GET("/objectlike/delete", objectlike.Delete)
	// Object comment calls
	m.GET("/objectcomment/create", objectcomment.Create)
	m.GET("/objectcomment/read", objectcomment.Read)
	m.GET("/objectcomment/delete", objectcomment.Delete)
	
	return m
}

