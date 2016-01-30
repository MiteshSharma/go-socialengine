package channel

import (
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"helper"
	"encoding/json"
	"middleware"
	"model/channel"
	"model/channeldetail"
	"net/http"
	"strconv"
	"strings"
)

func Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var ch channel.Channel
	if err := json.Unmarshal([]byte(r.FormValue("channel")), &ch); err != nil {
		middleware.Output.ResponseCode = http.StatusBadRequest
	}
	userId := context.Get(r, "user_id").(int)

	channel := channel.Create(ch.OwnerId, ch.Name, ch.Description)
	channeldetail.Create(userId, channel.Id)

	responseJson, responseCode := helper.GetResponseJson(channel)
	middleware.Output.Response = responseJson
	middleware.Output.ResponseCode = responseCode
}

func Read(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	requestType := r.FormValue("request_type")
	channelIdStr := r.FormValue("channel_id")
	userId := context.Get(r, "user_id").(int)

	var channels []channel.Channel

	if strings.EqualFold(requestType, "owner") {
		channels = channel.Read(userId)
	} else if strings.EqualFold(requestType, "id") {
		if helper.IsValidRequest(channelIdStr) {
			channelId, _ := strconv.Atoi(channelIdStr)
			channels = channel.ReadById(channelId)
		}
	} else {
		channels = nil
	}
	if channels != nil {
		responseJson, responseCode := helper.GetResponseJson(channels)
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
		channel.DeleteById(channelId, userId)
	}
}
