// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

// +build wasm

package main

import (
	"github.com/iotaledger/wasp/packages/vm/wasmclient"

	alphainterfacecontract "github.com/iotaledger/wasp/alphainterface/src/contract/go/contract"
)

func main() {
}

//export on_load
func onLoad() {
	h := &wasmclient.WasmVMHost{}
	h.ConnectWasmHost()
	alphainterfacecontract.OnLoad()
}
