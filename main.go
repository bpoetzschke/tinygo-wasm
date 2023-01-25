package main

import (
	"syscall/js"
)

//https://www.andreagrandi.it/2020/10/23/getting-started-with-tinygo-webassembly/
//https://github.com/justinclift/tinygo_canvas2/blob/8872c737b8ae18420fbc9ed0d8beb62e9b1f89b0/wasm.go

var (
	doc, canvasDiv, ctx       js.Value
	width, height, boxX, boxY int
	stepSize                  = 5
)

func main() {
	doc = js.Global().Get("document")
}

// This function is exported to JavaScript, so can be called using
// exports.add() in JavaScript.
//
//export add
func add(x, y int) int {
	return x + y
}

//export start
func start() {

	width = js.Global().Get("innerWidth").Int()
	height = js.Global().Get("innerHeight").Int()
	boxX = 20
	boxY = 20

	canvasDiv = doc.Call("getElementById", "my-canvas")
	canvasDiv.Call("setAttribute", "width", width)
	canvasDiv.Call("setAttribute", "height", height)
	ctx = canvasDiv.Call("getContext", "2d")

	//canvasBoundingRect := canvasDiv.Call("getBoundingClientRect")
	//
	//width := canvasBoundingRect.Get("width").Int()
	//height := canvasBoundingRect.Get("height").Int()

	ctx.Set("strokeStyle", "black")
	ctx.Set("fillStyle", "black")
	ctx.Set("lineWidth", "1")
	ctx.Call("beginPath")
	ctx.Call("rect", boxX, boxY, 150, 100)
	ctx.Call("stroke")

	js.Global().Call("requestAnimationFrame", js.Global().Get("renderFrame"))
}

//export renderFrame
func renderFrame() {
	ctx.Set("fillStyle", "white")
	ctx.Call("fillRect", 0, 0, width, height)

	boxX += stepSize
	boxY += stepSize

	ctx.Set("strokeStyle", "black")
	ctx.Set("fillStyle", "black")
	ctx.Set("lineWidth", "1")
	ctx.Call("beginPath")
	ctx.Call("fillRect", boxX, boxY, 150, 100)
	ctx.Call("stroke")

	js.Global().Call("requestAnimationFrame", js.Global().Get("renderFrame"))
}
