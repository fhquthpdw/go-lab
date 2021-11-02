package main

import (
	"fmt"
	"log"
	"reflect"
	"runtime"

	"github.com/gin-gonic/gin"
)

func main() {
	run()
}

func run() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(AuthMiddleware())

	us1 := u1{}
	r.GET("/user1", decorator(us1.user1))
	n1 := runtime.FuncForPC(reflect.ValueOf(us1.user1).Pointer()).Name() // 读取 us1.user1 的name
	fmt.Printf("from router user1: %s\n", n1)
	//fmt.Println(reflect.TypeOf(us1.user1).In(0))

	us2 := u2{}
	r.GET("/user2/:id", us2.user2)
	n2 := runtime.FuncForPC(reflect.ValueOf(us2.user2).Pointer()).Name()
	fmt.Printf("from router user2: %s\n", n2)

	fmt.Println("+++++++++++ " + H(us2.user2).getName() + " +++++++")

	if err := r.Run("0.0.0.0:8070"); err != nil {
		log.Fatalf(err.Error())
	}
}

// H generate router key, for do permission check
type H gin.HandlerFunc

func (h H) getName() string {
	return runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
}

func decorator(h func(ctx *gin.Context)) func(c *gin.Context) {
	return func(c *gin.Context) {
		h(c)
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("+++ Mid +++")

		f := c.Handler()
		n1 := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
		fmt.Printf("decorator from middleware: %s\n", n1)

		n2 := reflect.TypeOf(f).In(0).Name()
		fmt.Printf("actual from middleware: %s\n", n2)
		//n3 := runtime.FuncForPC(reflect.ValueOf(n2).Pointer()).Name()
		//fmt.Printf("actual from middleware: %s\n", n3)

		fmt.Println("+++ Mid +++")
		c.Next()
	}
}

type u1 struct{}

func (u u1) user1(c *gin.Context) {
	//fmt.Println("************")
	//fmt.Println(c.HandlerName())
	//fmt.Println("************")
}

type u2 struct{}

func (u u2) user2(c *gin.Context) {
	//fmt.Println("************")
	//fmt.Println(c.HandlerName())
	//fmt.Println("************")
}
