package user

import (
	"github.com/julienschmidt/httprouter"
	"helper"
	"middleware"
	"model/user"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	email := r.FormValue("email")

	if helper.IsValidRequest(email) {
		user := user.Create(email)

		responseJson, responseCode := helper.GetResponseJson(user)
		middleware.Output.Response = responseJson
		middleware.Output.ResponseCode = responseCode
	} else {
		middleware.Output.ResponseCode = http.StatusBadRequest
	}
}
