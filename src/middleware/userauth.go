package middleware

import (
	"github.com/gorilla/context"
	"model/user"
	"net/http"
)

type UserAuth struct {
}

func NewUserAuth() *UserAuth {
	return &UserAuth{}
}

func (ua *UserAuth) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// Validate user id before sending user ahead
	r.ParseForm()

	if r.URL.Path[1:] != "user/create" {
		var email = r.FormValue("email")

		var isValid = user.Validate(email)
		if isValid == false {
			Output.ResponseCode = http.StatusUnauthorized
			return
		} else {
			context.Set(r, "user_id", user.Get(email).UserId)
		}
	}
	// If everything goes well, send request to next handler
	next(rw, r)
}
