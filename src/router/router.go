package router

import (
	"github.com/julienschmidt/httprouter"
    "controller"
    "controller/userstatuses"
    "controller/userstatuslikes"
    "controller/userstatuscomments"
)

func RouteHandler() *httprouter.Router {
	m := httprouter.New()

	m.GET("/", controller.MainHandler)

	// User status calls
	m.GET("/userstatus/create", userstatuses.Create)
	m.GET("/userstatus/update", userstatuses.Update)
	m.GET("/userstatus/delete", userstatuses.Delete)
	m.GET("/userstatus/read", userstatuses.Read)
	m.GET("/userstatuslike/create", userstatuslikes.Create)
	m.GET("/userstatuslike/delete", userstatuslikes.Delete)
	m.GET("/userstatuscomment/create", userstatuscomments.Create)
	m.GET("/userstatuscomment/update", userstatuscomments.Update)
	m.GET("/userstatuscomment/delete", userstatuscomments.Delete)
	
	return m
}

