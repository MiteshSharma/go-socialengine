package objectlike

import (
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"helper"
	"middleware"
	"model/objectlike"
	"net/http"
	"strconv"
	"strings"
)

func Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	objectIdStr := r.FormValue("object_id")
	objectType := r.FormValue("object_type")
	userId := context.Get(r, "user_id").(int)

	if helper.IsValidRequest(objectIdStr, objectType) {
		objectId, _ := strconv.Atoi(objectIdStr)

		objectLike := objectlike.Create(userId, objectId, objectType)

		responseJson, responseCode := helper.GetResponseJson(objectLike)
		middleware.Output.Response = responseJson
		middleware.Output.ResponseCode = responseCode
	} else {
		middleware.Output.ResponseCode = http.StatusBadRequest
	}
}

func Read(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	requestType := r.FormValue("request_type")
	objectLikeIdStr := r.FormValue("object_like_id")
	objectIdStr := r.FormValue("object_id")
	objectType := r.FormValue("object_type")
	userId := context.Get(r, "user_id").(int)

	var objectLikes []objectlike.ObjectLike

	if strings.EqualFold(requestType, "user") {
		objectLikes = objectlike.Read(userId)
	} else if strings.EqualFold(requestType, "object") {
		if helper.IsValidRequest(objectIdStr) {
			objectId, _ := strconv.Atoi(objectIdStr)
			objectLikes = objectlike.ReadByObject(objectId, objectType)
		}
	} else if strings.EqualFold(requestType, "id") {
		if helper.IsValidRequest(objectLikeIdStr) {
			objectLikeId, _ := strconv.Atoi(objectLikeIdStr)
			objectLikes = objectlike.ReadById(objectLikeId)
		}
	}
	if objectLikes != nil {
		responseJson, responseCode := helper.GetResponseJson(objectLikes)
		middleware.Output.Response = responseJson
		middleware.Output.ResponseCode = responseCode
	} else {
		middleware.Output.ResponseCode = http.StatusBadRequest
	}
}

func Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	requestType := r.FormValue("request_type")
	objectLikeIdStr := r.FormValue("object_like_id")
	objectIdStr := r.FormValue("object_id")
	objectType := r.FormValue("object_type")
	userId := context.Get(r, "user_id").(int)

	if strings.EqualFold(requestType, "object") {
		if helper.IsValidRequest(objectIdStr) {
			objectId, _ := strconv.Atoi(objectIdStr)
			objectlike.Delete(userId, objectId, objectType)
		}
	} else if strings.EqualFold(requestType, "id") {
		if helper.IsValidRequest(objectLikeIdStr) {
			objectLikeId, _ := strconv.Atoi(objectLikeIdStr)
			objectlike.DeleteById(objectLikeId)
		}
	}
}
