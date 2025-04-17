package common

import (
	"context"
	"net/http"
)

type Response struct {
	Code    int
	Message string
	Data    any
}

func ErrorHandler(err error) (int, any) {
	return http.StatusOK, Response{
		Code:    http.StatusInternalServerError,
		Message: err.Error(),
	}
}
func SuccessHandler(c context.Context, resp any) any {
	return Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    resp,
	}
}
