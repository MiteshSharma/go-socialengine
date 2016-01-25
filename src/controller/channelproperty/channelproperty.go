package channelproperty


import (
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"helper"
	"middleware"
	"net/http"
	"strconv"
	"model/channelproperty"
)

func Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	channelIdStr := r.FormValue("channel_id")
	name := r.FormValue("name")
	value := r.FormValue("value")
	userId := context.Get(r, "user_id").(int)

	if helper.IsValidRequest(channelIdStr, name) {
		channelId, _ := strconv.Atoi(channelIdStr)
		channelProperty := channelproperty.Create(channelId, userId, name, value)
		responseJson, responseCode := helper.GetResponseJson(channelProperty)
		middleware.Output.Response = responseJson
		middleware.Output.ResponseCode = responseCode
	} else {
		middleware.Output.ResponseCode = http.StatusBadRequest
	}
}

func Read(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	channelIdStr := r.FormValue("channel_id")
	name := r.FormValue("name")

	if helper.IsValidRequest(channelIdStr) {
		channelId, _ := strconv.Atoi(channelIdStr)
		channelProperty := channelproperty.Read(channelId, name)
		responseJson, responseCode := helper.GetResponseJson(channelProperty)
		middleware.Output.Response = responseJson
		middleware.Output.ResponseCode = responseCode
	} else {
		middleware.Output.ResponseCode = http.StatusBadRequest
	}
}

func Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	channelIdStr := r.FormValue("channel_id")
	name := r.FormValue("name")

	if helper.IsValidRequest(channelIdStr) {
		channelId, _ := strconv.Atoi(channelIdStr)
		channelProperty := channelproperty.Read(channelId, name)
		responseJson, responseCode := helper.GetResponseJson(channelProperty)
		middleware.Output.Response = responseJson
		middleware.Output.ResponseCode = responseCode
	} else {
		middleware.Output.ResponseCode = http.StatusBadRequest
	}
}