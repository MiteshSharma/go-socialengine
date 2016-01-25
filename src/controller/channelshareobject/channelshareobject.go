package channelshareobject

import (
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"helper"
	"middleware"
	"model/channelshareobject"
	"net/http"
	"strconv"
	"strings"
)

func Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	channelIdStr := r.FormValue("channel_id")
	objectIdStr := r.FormValue("object_id")
	objectType := r.FormValue("object_type")
	message := r.FormValue("message")
	userId := context.Get(r, "user_id").(int)

	if helper.IsValidRequest(channelIdStr, objectIdStr, objectType) || helper.IsValidRequest(channelIdStr, message) {
		channelId, _ := strconv.Atoi(channelIdStr)
		objectId, _ := strconv.Atoi(objectIdStr)
		channelshareobject := channelshareobject.Create(userId, channelId, objectId, objectType, message)
		responseJson, responseCode := helper.GetResponseJson(channelshareobject)
		middleware.Output.Response = responseJson
		middleware.Output.ResponseCode = responseCode
	} else {
		middleware.Output.ResponseCode = http.StatusBadRequest
	}
}

func Read(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	requestType := r.FormValue("request_type")
	channelIdStr := r.FormValue("channel_id")
	userId := context.Get(r, "user_id").(int)

	var channelshareobjects []channelshareobject.ChannelShareObject

	if strings.EqualFold(requestType, "id") {
		if helper.IsValidRequest(channelIdStr) {
			channelId, _ := strconv.Atoi(channelIdStr)
			channelshareobjects = channelshareobject.Read(channelId, userId)
		}
	} else if strings.EqualFold(requestType, "channel") {
		if helper.IsValidRequest(channelIdStr) {
			channelId, _ := strconv.Atoi(channelIdStr)
			channelshareobjects = channelshareobject.ReadByChannel(channelId)
		}
	} else {
		channelshareobjects = nil
	}
	if channelshareobjects != nil {
		responseJson, responseCode := helper.GetResponseJson(channelshareobjects)
		middleware.Output.Response = responseJson
		middleware.Output.ResponseCode = responseCode
	} else {
		middleware.Output.ResponseCode = http.StatusBadRequest
	}
}

func Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	channelShareObjectIdStr := r.FormValue("channel_share_object_id")
	userId := context.Get(r, "user_id").(int)

	if helper.IsValidRequest(channelShareObjectIdStr) {
		channelShareObjectId, _ := strconv.Atoi(channelShareObjectIdStr)
		channelshareobject.Delete(channelShareObjectId, userId)
	}
}
