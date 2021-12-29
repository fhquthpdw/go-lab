package main

import (
	"fmt"
	"reflect"
)

type Phone interface {
	Call()
	SendMsg(s string) bool
}

type Android struct {
	Name string
}

func (a Android) Call() {
	fmt.Println("call")
}
func (a Android) SendMsg(s string) bool {
	fmt.Println(s)
	return true
}

type Ios struct{}

func (a Ios) Call() {
	fmt.Println("call")
}
func (a Ios) SendMsg(s string) bool {
	fmt.Println(s)
	return true
}

func main() {
	var p Phone = Android{}

	v := reflect.ValueOf(p)
	fmt.Println(v.Kind())
	fmt.Println(v.String())
	fmt.Println(v.Elem())
	fmt.Println()

	//
	t := reflect.TypeOf(p)
	fmt.Println(t.Kind())
	fmt.Println(t.String())
}
