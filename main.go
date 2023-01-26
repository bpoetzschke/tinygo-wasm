package main

import (
	"fmt"
	"syscall/js"
)

//https://www.andreagrandi.it/2020/10/23/getting-started-with-tinygo-webassembly/
//https://github.com/justinclift/tinygo_canvas2/blob/8872c737b8ae18420fbc9ed0d8beb62e9b1f89b0/wasm.go

var (
	doc, canvasDiv, ctx                                                    js.Value
	width, height, boxX, boxY, boxWidth, boxHeight, xDirection, yDirection int
	stepSize                                                               = 1
	colorMap                                                               = []string{
		"black",
		"indianred",
		"lightblue",
		"lightgreen",
	}
	colorIdx = 0
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

	println(fmt.Sprintf("width: %d, height: %d", width, height))

	xDirection = 1
	yDirection = 1
	boxX = 20
	boxY = 20
	boxHeight = 100
	boxWidth = 150

	canvasDiv = doc.Call("getElementById", "my-canvas")
	canvasDiv.Call("setAttribute", "width", width)
	canvasDiv.Call("setAttribute", "height", height)
	ctx = canvasDiv.Call("getContext", "2d")

	ctx.Set("strokeStyle", "black")
	ctx.Set("fillStyle", "black")
	ctx.Set("lineWidth", "1")
	ctx.Call("beginPath")
	ctx.Call("rect", boxX, boxY, boxWidth, boxHeight)
	ctx.Call("stroke")

	js.Global().Call("requestAnimationFrame", js.Global().Get("renderFrame"))
}

//export renderFrame
func renderFrame() {
	ctx.Set("fillStyle", "white")
	ctx.Call("fillRect", 0, 0, width, height)

	boxX += stepSize * xDirection
	boxY += stepSize * yDirection

	if boxY+boxHeight > height {
		boxY = height - boxHeight
		yDirection = -1
	}

	if boxY < 0 {
		boxY = 0
		yDirection = 1
	}

	if boxX+boxWidth > width {
		boxX = width - boxWidth
		xDirection = -1
	}

	if boxX < 0 {
		boxX = 0
		xDirection = 1
	}

	color := colorMap[colorIdx/30]
	ctx.Set("strokeStyle", color)
	ctx.Set("fillStyle", color)
	ctx.Set("lineWidth", "1")
	ctx.Call("beginPath")
	ctx.Call("fillRect", boxX, boxY, 150, 100)
	ctx.Call("stroke")

	colorIdx++
	if colorIdx > (len(colorMap)*30)-1 {
		colorIdx = 0
	}

	js.Global().Call("requestAnimationFrame", js.Global().Get("renderFrame"))
}
