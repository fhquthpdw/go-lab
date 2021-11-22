package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

///////////////////////////////////
// 如何避免出现连续的 if err := f(); err != nil {} 的语句
func f(p person) error {
	var err error
	checkErr := func(f func() error) {
		// 任何一个 check 出错后，后面的 check 就不会再检查
		if err != nil {
			return
		}
		err = f()
	}

	checkErr(p.f1)
	checkErr(p.f2)
	checkErr(p.f3)
	checkErr(p.f4)

	return err
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
type ginHandlerWithErr func(c *gin.Context) error

func httpServer() {
	r := gin.Default()

	r.POST("/data/get/list1", handler(dataList1))
	r.POST("/data/get/list2", handler(dataList2))
}

func handler(h ginHandlerWithErr) func(c *gin.Context) {
	return func(c *gin.Context) {
		if err := h(c); err != nil {
			// log
			c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		}
	}
}

// 一般都会这么写
func dataList1(c *gin.Context) (err error) {
	if err = check1(); err != nil {
		return err
	}
	if err = check2(); err != nil {
		return err
	}
	if err = check3(); err != nil {
		return err
	}
	if err = check4(); err != nil {
		return err
	}

	c.JSON(http.StatusOK, "Done")
	return nil
}

// 还可以这么写
func dataList2(c *gin.Context) (err error) {
	checkErr := func(f func() error) {
		if err != nil {
			return
		}
		err = f()
	}

	checkErr(check1)
	checkErr(check2)
	checkErr(check3)
	checkErr(check4)

	c.JSON(http.StatusOK, "Done")
	return
}

func check1() error { return nil }
func check2() error { return nil }
func check3() error { return nil }
func check4() error { return nil }
