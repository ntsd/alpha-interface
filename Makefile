
run-wasp:
	wasp -c ./wasp-nodes/config.json

wasm:
	make build-wasm
	make deploy-wasm

build-wasm:
	tinygo build -o alphainterface_go.wasm -target wasm ./src/contract/go/main.go

deploy-wasm:
	wasp-cli chain deploy-contract wasmtime alphainterface "Alpha Interface" alphainterface_go.wasm -d --verbose
	
deploy-chain:
	wasp-cli chain deploy --committee=0 --quorum=1 --chain=alpha-interface-chain --description="Alpha Interface Chain"
	wasp-cli chain deposit IOTA:10000
