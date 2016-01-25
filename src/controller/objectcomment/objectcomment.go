package objectcomment

import (
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"helper"
	"middleware"
	"model/objectcomment"
	"net/http"
	"strconv"
	"strings"
)

func Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	objectIdStr := r.FormValue("object_id")
	objectType := r.FormValue("object_type")
	parentCommentIdStr := r.FormValue("parent_comment_id")
	comment := r.FormValue("comment")
	userId := context.Get(r, "user_id").(int)

	if helper.IsValidRequest(objectIdStr, objectType) {
		objectId, _ := strconv.Atoi(objectIdStr)
		parentCommentId, _ := strconv.Atoi(parentCommentIdStr)

		objectComment := objectcomment.Create(userId, objectId, parentCommentId, comment, objectType)

		responseJson, responseCode := helper.GetResponseJson(objectComment)
		middleware.Output.Response = responseJson
		middleware.Output.ResponseCode = responseCode
	} else {
		middleware.Output.ResponseCode = http.StatusBadRequest
	}
}

func Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	commentIdStr := r.FormValue("comment_id")
	comment := r.FormValue("comment")

	if helper.IsValidRequest(commentIdStr) {
		commentId, _ := strconv.Atoi(commentIdStr)
		objectComment := objectcomment.Update(commentId, comment)
		responseJson, responseCode := helper.GetResponseJson(objectComment)
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
	pageIdStr := r.FormValue("page_id")
	pageSizeStr := r.FormValue("page_size")
	userId := context.Get(r, "user_id").(int)

	var objectLikes []objectcomment.ObjectComment

	if strings.EqualFold(requestType, "user") {
		pageId, _ := strconv.Atoi(pageIdStr)
		pageSize, _ := strconv.Atoi(pageSizeStr)
		objectLikes = objectcomment.Read(userId, pageId, pageSize)
	} else if strings.EqualFold(requestType, "object") {
		if helper.IsValidRequest(objectIdStr) {
			objectId, _ := strconv.Atoi(objectIdStr)
			pageId, _ := strconv.Atoi(pageIdStr)
			pageSize, _ := strconv.Atoi(pageSizeStr)
			objectLikes = objectcomment.ReadByObject(objectId, pageId, pageSize, objectType)
		}
	} else if strings.EqualFold(requestType, "id") {
		if helper.IsValidRequest(objectLikeIdStr) {
			objectCommentId, _ := strconv.Atoi(objectLikeIdStr)
			objectLikes = objectcomment.ReadById(objectCommentId)
		}
	} else {
		objectLikes = nil
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
			objectcomment.Delete(userId, objectId, objectType)
		}
	} else if strings.EqualFold(requestType, "id") {
		if helper.IsValidRequest(objectLikeIdStr) {
			objectLikeId, _ := strconv.Atoi(objectLikeIdStr)
			objectcomment.DeleteById(objectLikeId)
		}
	}
}
