package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Base struct {
	//
}

func responseFail(ctx *gin.Context, err error) {
	ctx.AbortWithStatusJSON(http.StatusOK, "")
}

type jsonResponse struct {
	Status string        `json:"status,omitempty"` // success | fail
	Data   interface{}   `json:"data,omitempty"`
	Error  responseError `json:"error,omitempty"`
}
type responseError struct {
	Message string `json:"message,omitempty"`
}

type PortalCtx struct {
	*gin.Context
}
