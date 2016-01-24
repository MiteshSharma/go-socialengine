package controller

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func MainHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Print("Oh, we finally made our controller run .... ")
}
