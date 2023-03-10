TINYGO_DIR=$(shell tinygo env TINYGOROOT)

build-go:
	GOOS=js GOARCH=wasm go build -o wasm ./main.go

build:
	tinygo build -o ./out/add.wasm -target wasm ./add.go
	tinygo build -o ./out/go_callback.wasm -target wasm ./go_callback.go
	tinygo build -o ./out/wasm.wasm -target wasm ./main.go
	cp ${TINYGO_DIR}/targets/wasm_exec.js ./out/

run: build
	go run ./server.go