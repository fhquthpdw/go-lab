package main

import (
	"strconv"
	"syscall/js"
)

func main() {
	c := make(chan struct{}, 0)
	println("hello wasm")
	registerCallbacks()
	<-c
}

func sum(this js.Value, args []js.Value) interface{} {
	var sum int
	for _, val := range args {
		sum += val.Int()
	}
	println(sum)
	return sum
}

func registerCallbacks() {
	global := js.Global()
	document := global.Get("document")

	getElementId := func(id string) js.Value {
		return document.Call("getElementById", id)
	}
	aValue := getElementId("aValue")
	bValue := getElementId("bValue")
	cValue := getElementId("cValue")
	sumValue := getElementId("sum")

	sumButton := getElementId("sumButton")
	runButton := getElementId("runButton")

	onRun := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		println("button on click")
		return ""
	})
	onSum := js.FuncOf(func(this js.Value, args []js.Value) interface {} {
		a, _ := strconv.Atoi(aValue.Get("value").String())
		b, _ := strconv.Atoi(bValue.Get("value").String())
		c, _ := strconv.Atoi(cValue.Get("value").String())
		sumValue.Set("value", js.ValueOf(a + b + c))
		return ""
	})

	global.Set("sum", js.FuncOf(sum))
	sumButton.Call("addEventListener", "click", onSum)
	runButton.Call("addEventListener", "click", onRun)
}
