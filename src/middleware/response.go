package middleware

import (
	"net/http"
)

type Response struct {
}

func NewResponse() *Response {
	return &Response{}
}

type OutputObject struct {
	Response     string
	OutputType   string
	ResponseCode int
}

var Output OutputObject

func (ua *Response) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	Output = OutputObject{}
	Output.Response = "{}"
	Output.OutputType = "application/json"
	Output.ResponseCode = http.StatusOK
	next(rw, r)
	rw.Header().Set("Content-Type", Output.OutputType)
	rw.Write([]byte(Output.Response))
	rw.WriteHeader(Output.ResponseCode)
}
