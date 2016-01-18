package userstatuscomments

import (
	"github.com/julienschmidt/httprouter"
	"model/status/userstatuscomment"
	"net/http"
	"strconv"
)

func Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userStatusId, _ := strconv.Atoi(r.FormValue("user_status_id"))
	userId,_ := strconv.Atoi(r.FormValue("user_id"))
	comment := r.FormValue("comment")

	userstatuscomment.Create(userId, userStatusId, comment)
}

func Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userStatusCommentId, _ := strconv.Atoi(r.FormValue("user_status_comment_id"))
	userId,_ := strconv.Atoi(r.FormValue("user_id"))
	comment := r.FormValue("comment")

	userstatuscomment.Update(userStatusCommentId, userId, comment)
}

func Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userStatusCommentId, _ := strconv.Atoi(r.FormValue("user_status_comment_id"))
	userId,_ := strconv.Atoi(r.FormValue("user_id"))

	userstatuscomment.Delete(userStatusCommentId, userId)
}