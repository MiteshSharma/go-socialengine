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

func emptyJson() {
	return "{}"
}

func IsValidRequest(inputArgs ...string) bool {
	for inputArg := range inputArgs {
		if inputArg == nil {
			return false
		}
	}
	return true
}
