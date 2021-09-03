package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

///////////////////////////////////
// 如何避免出现连续的 if err := f(); err != nil {} 的语句
// 这样做的问题是，f1,f2,f3,f4 不管哪个出了一个 error 都会把 4 个 func 都执行一遍
func f(p person) {
	var err error
	checkErr := func(func() error) {
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	checkErr(p.f1)
	checkErr(p.f2)
	checkErr(p.f3)
	checkErr(p.f4)
}

type person struct {
	Age1 int
	Age2 int
	Age3 int
	Age4 int
	Age5 int
}

func (p person) f1() error { return nil }
func (p person) f2() error { return nil }
func (p person) f3() error { return nil }
func (p person) f4() error { return nil }

///////////////////////////////////
// 我们在处理 api 请求的时候，场影是这样：
// 一个请求过来，我需要做一系列的错误或者业务的检查，如果有任何一项检查出现了错误，我就马上 http response 给上游
// 这里以 Gin 框加为例，我们是如何处理这个错误的
type ginHandler func(c *gin.Context)
type ginHandlerWithErr func(c *gin.Context) error

func httpServer() {
	r := gin.Default()

	r.POST("/data/get/list", handler(dataList))
}

func handler(h ginHandlerWithErr) func(c *gin.Context) {
	return func(c *gin.Context) {
		if err := h(c); err != nil {
			// log
			c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		}
	}
}

func dataList(c *gin.Context) error {
	if err := check1(); err != nil {
		return err
	}
	if err := check2(); err != nil {
		return err
	}
	if err := check3(); err != nil {
		return err
	}
	if err := check4(); err != nil {
		return err
	}

	c.JSON(http.StatusOK, "Done")
	return nil
}

func check1() error { return nil }
func check2() error { return nil }
func check3() error { return nil }
func check4() error { return nil }
