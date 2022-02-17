package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Sys struct {
	Base
}

func (c Sys) Reset(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "OK")
}
