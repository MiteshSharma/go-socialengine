package channeldetail

import (
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"helper"
	"middleware"
	"model/channel"
	"model/channeldetail"
	"net/http"
	"strconv"
	"strings"
)

func Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	channelIdStr := r.FormValue("channel_id")
	userId := context.Get(r, "user_id").(int)

	if helper.IsValidRequest(channelIdStr) {
		channelId, _ := strconv.Atoi(channelIdStr)
		channelDetail := channeldetail.Create(userId, channelId)
		if channelDetail != (channeldetail.ChannelDetail{}) {
			channel.Update(channelId, 1)
			responseJson, responseCode := helper.GetResponseJson(channelDetail)
			middleware.Output.Response = responseJson
			middleware.Output.ResponseCode = responseCode
		}
	} else {
		middleware.Output.ResponseCode = http.StatusBadRequest
	}
}

func Read(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	requestType := r.FormValue("request_type")
	channelIdStr := r.FormValue("channel_id")
	userId := context.Get(r, "user_id").(int)

	var channelDetails []channeldetail.ChannelDetail

	if strings.EqualFold(requestType, "id") {
		channelDetails = channeldetail.Read(userId)
	} else if strings.EqualFold(requestType, "channel") {
		if helper.IsValidRequest(channelIdStr) {
			channelId, _ := strconv.Atoi(channelIdStr)
			channelDetails = channeldetail.ReadByChannel(channelId)
		}
	} else {
		channelDetails = nil
	}
	if channelDetails != nil {
		responseJson, responseCode := helper.GetResponseJson(channelDetails)
		middleware.Output.Response = responseJson
		middleware.Output.ResponseCode = responseCode
	} else {
		middleware.Output.ResponseCode = http.StatusBadRequest
	}
}

func Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	channelIdStr := r.FormValue("channel_id")
	userId := context.Get(r, "user_id").(int)

	if helper.IsValidRequest(channelIdStr) {
		channelId, _ := strconv.Atoi(channelIdStr)
		channels := channel.ReadById(channelId)
		if channels[0].OwnerId == userId {
			channeldetail.DeleteByChannel(channelId)
			channel.DeleteById(channelId, userId)
		} else {
			isDelete := channeldetail.Delete(userId, channelId)
			if isDelete {
				channel.Update(channelId, -1)
			}
		}
	}
}
