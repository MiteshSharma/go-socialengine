package userstatuslikes

import (
	"github.com/julienschmidt/httprouter"
	"model/status/userstatuslike"
	"net/http"
	"strconv"
)

func Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userStatusId, _ := strconv.Atoi(r.FormValue("user_status_id"))
	userId,_ := strconv.Atoi(r.FormValue("user_id"))

	userstatuslike.Create(userId, userStatusId)
}

func Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userStatusLikeId, _ := strconv.Atoi(r.FormValue("user_status_like_id"))
	userId,_ := strconv.Atoi(r.FormValue("user_id"))

	userstatuslike.Delete(userStatusLikeId, userId)
}