package helper

import (
	"strings"
)

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
	Data    interface{} `json:"data"`
}

type EmptyObj struct {
}

func BuildResponse(status bool, message string, errors interface{}, data interface{}) Response {
	r := Response{
		Status:  true,
		Message: message,
		Error:   nil,
		Data:    data,
	}
	return r
}

func BuildErrorResponse(message string, err string, data interface{}) Response {
	splitedError := strings.Split(err, "\n")

	r := Response{
		Status:  false,
		Message: message,
		Error:   splitedError,
		Data:    data,
	}

	return r
}
