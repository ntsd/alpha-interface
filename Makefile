
run-wasp:
	wasp -c ./wasp-nodes/config.json

wasm:
	make build-wasm
	make deploy-wasm

build-wasm:
	tinygo build -o ./src/contract/go/pkg/alphainterface_go.wasm -target wasm ./src/contract/go/main.go

deploy-wasm:
	wasp-cli chain deploy-contract wasmtime alphainterface "Alpha Interface" ./src/contract/go/pkg/alphainterface_go.wasm
