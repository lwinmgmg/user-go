package middlewares

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type DefResp struct {
	Code    int            `json:"code,omitempty"`
	Message string         `json:"message,omitempty"`
	Data    map[string]any `json:"data,omitempty"`
}

type PanicResponse struct {
	Response   DefResp
	HttpStatus int
	Error      error
}

func RecoveryMiddleware(ctx *gin.Context, err any) {
	switch err := err.(type) {
	case PanicResponse:
		val, _ := os.LookupEnv("GIN_MODE")
		if val != "release" && err.Error != nil {
			err.Response.Data["error"] = err.Error.Error()
		}
		ctx.AbortWithStatusJSON(err.HttpStatus, err.Response)
		return
	default:
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, DefResp{
			Code:    500,
			Message: fmt.Sprintf("Unknown error, %v", err),
		})
	}
}

func NewPanic(httpStatus, code int, mesg string, err error, data ...map[string]any) PanicResponse {
	respData := map[string]any{}
	if len(data) > 0 {
		respData = data[0]
	}
	return PanicResponse{
		HttpStatus: httpStatus,
		Response: DefResp{
			Code:    code,
			Message: mesg,
			Data:    respData,
		},
		Error: err,
	}
}
