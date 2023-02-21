package main

import (
	"strconv"
	"syscall/js"
)

var (
	doc js.Value
)

func main() {
	quit := make(chan struct{}, 0)
	doc = js.Global().Get("document")
	buttonEl := doc.Call("getElementById", "button")
	buttonEl.Set("onclick", js.FuncOf(goAdd))
	<-quit
}

func goAdd(js.Value, []js.Value) any {
	aEl := doc.Call("getElementById", "a")
	bEl := doc.Call("getElementById", "b")
	a, _ := strconv.Atoi(aEl.Get("value").String())
	b, _ := strconv.Atoi(bEl.Get("value").String())

	js.Global().Call("alert", a+b)
	return ""
}
