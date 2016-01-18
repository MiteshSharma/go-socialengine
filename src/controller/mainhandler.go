package controller

import (
	"github.com/julienschmidt/httprouter"
    "fmt"
    "net/http"
	)

func MainHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Print("Oh, we finally made our controller run .... ")
}