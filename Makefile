
run-wasp:
	wasp -c ./wasp-nodes/config.json

build-wasm:
	tinygo build -o alphainterface.wasm -target wasm ./src/contract/wasmmain/main.go

deploy-wasm:
	wasp-cli chain deploy-contract wasmtime alphainterface "Alpha Interface" alphainterface.wasm
