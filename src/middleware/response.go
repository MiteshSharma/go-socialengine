package middleware

import(
	"net/http"
)
type Response struct {
}

func NewResponse() *Response {
	return &Response{}
}

var Output string
var OutputType string

func (ua *Response) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	Output = "{}"
	OutputType = "application/json"
	next(rw, r)
	rw.Header().Set("Content-Type", OutputType)
  	rw.Write([]byte(Output))
}