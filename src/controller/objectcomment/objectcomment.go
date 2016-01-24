package objectcomment

import (
	"github.com/julienschmidt/httprouter"
	"model/objectcomment"
	"net/http"
	"strconv"
	"helper"
	"middleware"
)

func Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	objectIdStr := ps.ByName("object_id")
	objectType := ps.ByName("object_type")
	parentCommentIdStr := ps.ByName("parent_comment_id")
	comment := ps.ByName("comment")
	userIdStr := ps.ByName("user_id")

	if helper.IsValidRequest(objectIdStr, objectType, userIdStr) {
		objectId, _ := strconv.Atoi(objectIdStr)
		userId, _ := strconv.Atoi(userIdStr)
		parentCommentId, _ := strconv.Atoi(parentCommentIdStr)

		objectComment := objectcomment.Create(userId, objectId, parentCommentId, comment, objectType)

		responseJson, responseCode := helper.GetResponseJson(objectComment)
		middleware.Output.Response = responseJson
		middleware.Output.ResponseCode = responseCode
	} else {
		middleware.Output.ResponseCode = http.StatusBadRequest
	}
}

func Read(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	requestType := ps.ByName("request_type")
	objectLikeIdStr := ps.ByName("object_like_id")
	objectIdStr := ps.ByName("object_id")
	objectType := ps.ByName("object_type")
	userIdStr := ps.ByName("user_id")

	var objectLikes []objectcomment.ObjectComment

	switch requestType {
	case "user":
		if helper.IsValidRequest(userIdStr) {
			userId, _ := strconv.Atoi(userIdStr)
			objectLikes = objectcomment.Read(userId)
		}
		break
	case "object":
		if helper.IsValidRequest(objectIdStr) {
			objectId, _ := strconv.Atoi(objectIdStr)
			objectLikes = objectcomment.ReadByObject(objectId, objectType)
		}
		break
	case "id":
		if helper.IsValidRequest(objectLikeIdStr) {
			objectLikeId, _ := strconv.Atoi(objectIdStr)
			objectLikes = objectcomment.ReadById(objectLikeId)
		}
		break
	default:
		objectLikes = nil
		break
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
	requestType := ps.ByName("request_type")
	objectLikeIdStr := ps.ByName("object_like_id")
	objectIdStr := ps.ByName("object_id")
	objectType := ps.ByName("object_type")
	userIdStr := ps.ByName("user_id")

	switch requestType {
	case "object":
		if helper.IsValidRequest(objectIdStr) {
			objectLikeId, _ := strconv.Atoi(objectLikeIdStr)
			userId, _ := strconv.Atoi(userIdStr)
			objectcomment.Delete(userId, objectLikeId, objectType)
		}
		break
	case "id":
		if helper.IsValidRequest(objectLikeIdStr) {
			objectLikeId, _ := strconv.Atoi(objectIdStr)
			objectcomment.DeleteById(objectLikeId)
		}
		break
	default:
		break
	}
}
