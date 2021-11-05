
run-wasp:
	wasp -c ./wasp-nodes/config.json

build:
	tinygo build -o wasm.wasm -target wasm ./src/contract/main.go

deploy:
	wasp-cli chain deploy-contract wasmtime alphainterface "Alpha Interface" tools/cluster/tests/wasm/inccounter_bg.wasm
