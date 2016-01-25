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
	// Channel calls
	m.GET("/channel/create", channel.Create)
	m.GET("/channel/read", channel.Read)
	m.GET("/channel/delete", channel.Delete)
	// Channel detail calls
	m.GET("/channeldetail/create", channeldetail.Create)
	m.GET("/channeldetail/read", channeldetail.Read)
	m.GET("/channeldetail/delete", channeldetail.Delete)
	// Channel property calls
	m.GET("/channelproperty/create", channelproperty.Create)
	m.GET("/channelproperty/update", channelproperty.Update)
	m.GET("/channelproperty/read", channelproperty.Read)
	// Channel share object calls
	m.GET("/channelshareobject/create", channelshareobject.Create)
	m.GET("/channelshareobject/read", channelshareobject.Read)
	m.GET("/channelshareobject/delete", channelshareobject.Delete)
	return m
}
