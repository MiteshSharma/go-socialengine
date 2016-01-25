package objectdetail

import (
	"github.com/julienschmidt/httprouter"
	"helper"
	"middleware"
	"model/objectdetail"
	"net/http"
	"strconv"
)

func Read(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	objectIdStr := r.FormValue("object_id")
	objectType := r.FormValue("object_type")

	if helper.IsValidRequest(objectIdStr, objectType) {
		objectId, _ := strconv.Atoi(objectIdStr)
		objectDetail := objectdetail.Read(objectId, objectType)

		responseJson, responseCode := helper.GetResponseJson(objectDetail)
		middleware.Output.Response = responseJson
		middleware.Output.ResponseCode = responseCode
	} else {
		middleware.Output.ResponseCode = http.StatusBadRequest
	}
}
