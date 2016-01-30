package router

import (
	"controller"
	"controller/objectcomment"
	"controller/objectdetail"
	"controller/objectlike"
	"controller/user"
	"controller/channel"
	"controller/channeldetail"
	"controller/channelproperty"
	"controller/channelshareobject"
	"github.com/julienschmidt/httprouter"
)

func RouteHandler() *httprouter.Router {
	m := httprouter.New()

	m.GET("/", controller.MainHandler)

	// User
	m.POST("/user", user.Create)
	// Object details
	m.GET("/objectdetail", objectdetail.Read)
	// Object like calls
	m.POST("/objectlike", objectlike.Create)
	m.GET("/objectlike", objectlike.Read)
	m.DELETE("/objectlike", objectlike.Delete)
	// Object comment calls
	m.POST("/objectcomment", objectcomment.Create)
	m.PUT("/objectcomment", objectcomment.Update)
	m.GET("/objectcomment", objectcomment.Read)
	m.DELETE("/objectcomment", objectcomment.Delete)
	// Channel calls
	m.POST("/channel", channel.Create)
	m.GET("/channel", channel.Read)
	m.DELETE("/channel", channel.Delete)
	// Channel detail calls
	m.POST("/channeldetail", channeldetail.Create)
	m.GET("/channeldetail", channeldetail.Read)
	m.DELETE("/channeldetail", channeldetail.Delete)
	// Channel property calls
	m.POST("/channelproperty", channelproperty.Create)
	m.PUT("/channelproperty", channelproperty.Update)
	m.GET("/channelproperty", channelproperty.Read)
	// Channel share object calls
	m.POST("/channelshareobject", channelshareobject.Create)
	m.GET("/channelshareobject", channelshareobject.Read)
	m.DELETE("/channelshareobject", channelshareobject.Delete)
	return m
}
