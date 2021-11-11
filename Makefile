
run-wasp:
	wasp -c ./wasp-nodes/config.json

wasm:
	make build-wasm
	make deploy-wasm

build-wasm:
	tinygo build -o alphainterface2_go.wasm -target wasm ./src/alphainterface2/go/main.go

deploy-wasm:
	wasp-cli chain deploy-contract wasmtime alphainterface "Alpha Interface" alphainterface2_go.wasm -d --verbose
	
deploy-chain:
	wasp-cli chain deploy --committee=0 --quorum=1 --chain=alpha-interface-chain --description="Alpha Interface Chain"
	wasp-cli chain deposit IOTA:10000
