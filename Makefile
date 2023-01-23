TINYGO_DIR=$(shell tinygo env TINYGOROOT)

build:
	tinygo build -o ./out/wasm.wasm -target wasm ./main.go
	cp ${TINYGO_DIR}/targets/wasm_exec.js ./out/