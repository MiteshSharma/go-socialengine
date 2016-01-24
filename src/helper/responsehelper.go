package helper

import (
	"encoding/json"
	"net/http"
)

func GetResponseJson(object interface{}) (response string, responseCode int) {
	js, err := json.Marshal(object)
	if err != nil {
		return emptyJson(), http.StatusInternalServerError
	}
	return string(js[:]), http.StatusOK
}

func emptyJson() string {
	return "{}"
}

func IsValidRequest(inputArgs ...string) bool {
	for _,inputArg := range inputArgs {
		if len(inputArg) == 0 {
			return false
		}
	}
	return true
}
