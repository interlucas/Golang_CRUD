package rest_err

import "net/http"

type RestErr struct {
	Message string `json:"message"`
	Err string `json:"error"`
	Code int64 `json:"code"`
	Causes []Causes `json:"causes"`
}
type Causes struct {
	Field string `json:"field"`
	Message string `json:"message"`
}
func (r*RestErr) Error() string{
	return r.Message
}

func NewRestErr(message, err string, code int64, causes []Causes)*RestErr{
	return &RestErr{
		Message: message,
		Err: err,
		Code: code,
		Causes: causes,

	}
}
type RestErr struct {
	Message string `json:"message"`
	Err string `json:"error"`
	Code int `json:"code"`
	Causes []Causes `json:"causes"`
}
