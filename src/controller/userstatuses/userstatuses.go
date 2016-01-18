package userstatuses

import (
	"github.com/julienschmidt/httprouter"
	"model/status/userstatus"
	"net/http"
	"strconv"
	"encoding/json"
	"middleware"
)

func Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	status := r.FormValue("status")
	status_type := r.FormValue("type")
	user_id,_ := strconv.Atoi(r.FormValue("user_id"))

	userStatus := userstatus.Create(user_id, status, status_type)
	
	js, err := json.Marshal(userStatus)
  	if err != nil {
    	http.Error(w, err.Error(), http.StatusInternalServerError)
    	return
  	}
	middleware.Output = string(js[:])
}

func Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userStatusId,_ := strconv.Atoi(r.FormValue("user_status_id"))
	userId,_ := strconv.Atoi(r.FormValue("user_id"))
	userstatus.Delete(userStatusId, userId)
}

func Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userStatusId,_ := strconv.Atoi(r.FormValue("user_status_id"))
	status := r.FormValue("status")
	userstatus.Update(userStatusId, status)
}

func Read(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userId,_ := strconv.Atoi(r.FormValue("user_id"))
	isFull,_ := strconv.Atoi(r.FormValue("is_full"))
	userStatusId,_ := strconv.Atoi(r.FormValue("user_status_id"))
	
	if isFull == 1 {
		userStatuses := userstatus.ReadUserStatus(userId, isFull)
		if !userstatus.IsEmpty(userStatuses) {
			js, err := json.Marshal(userStatuses)
		  	if err != nil {
		    	http.Error(w, err.Error(), http.StatusInternalServerError)
		    	return
		  	}
			middleware.Output = string(js[:])
		}
	} else {
		userStatusDetail := userstatus.Read(userStatusId)
		
		js, err := json.Marshal(userStatusDetail)
	  	if err != nil {
	    	http.Error(w, err.Error(), http.StatusInternalServerError)
	    	return
	  	}
		middleware.Output = string(js[:])
	}
}