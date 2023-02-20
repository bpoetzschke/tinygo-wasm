package main

import (
	"fmt"
	"strconv"
	"syscall/js"
)

func main() {
}

//export do
func do() {
	doc := js.Global().Get("document")
	aEl := doc.Call("getElementById", "a")
	bEl := doc.Call("getElementById", "b")
	a, _ := strconv.Atoi(aEl.Get("value").String())
	b, _ := strconv.Atoi(bEl.Get("value").String())

	print("foo")
	js.Global().Call("alert", fmt.Sprintf("%d", a+b))
}
