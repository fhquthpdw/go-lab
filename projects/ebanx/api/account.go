package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Account struct {
	Base
}

func (c Account) Event(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "event ok!")
}

func (c Account) Balance(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "balance ok!")
}
