package main

//https://www.andreagrandi.it/2020/10/23/getting-started-with-tinygo-webassembly/
//https://github.com/justinclift/tinygo_canvas2/blob/8872c737b8ae18420fbc9ed0d8beb62e9b1f89b0/wasm.go

func main() {

}

// This function is exported to JavaScript, so can be called using
// exports.add() in JavaScript.
//
//export add
func add(x, y int) int {
	return x + y
}
