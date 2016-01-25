package router

import (
	"controller"
	"controller/objectcomment"
	"controller/objectdetail"
	"controller/objectlike"
	"controller/user"
	"github.com/julienschmidt/httprouter"
)

func RouteHandler() *httprouter.Router {
	m := httprouter.New()

	m.GET("/", controller.MainHandler)

	// User
	m.POST("/user/create", user.Create)
	// Object details
	m.GET("/objectdetail/read", objectdetail.Read)
	// Object like calls
	m.POST("/objectlike/create", objectlike.Create)
	m.GET("/objectlike/read", objectlike.Read)
	m.DELETE("/objectlike/delete", objectlike.Delete)
	// Object comment calls
	m.POST("/objectcomment/create", objectcomment.Create)
	m.PUT("/objectcomment/update", objectcomment.Update)
	m.GET("/objectcomment/read", objectcomment.Read)
	m.DELETE("/objectcomment/delete", objectcomment.Delete)

	return m
}
