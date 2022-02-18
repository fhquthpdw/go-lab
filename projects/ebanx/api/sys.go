package api

import (
	"ebanx/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Sys struct {
	Base
}

func (c Sys) Reset(ctx *gin.Context) {
	model.NewSysModel().Reset()
	ctx.String(http.StatusOK, "OK")
}
