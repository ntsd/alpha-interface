
run-wasp:
	wasp -c ./wasp-nodes/config.json

wasm:
	make build-wasm
	make deploy-wasm

build-wasm:
	tinygo build -o alphainterface_go.wasm -target wasm ./src/contract/go/main.go

deploy-wasm:
	wasp-cli chain deploy-contract wasmtime alphainterface_0.1 "Alpha Interface" alphainterface_go.wasm -d --verbose
	