// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

// +build wasm

package main

import (
	alphainterfacecontract "github.com/iotaledger/wasp/alpha-interface/src/contract"
	"github.com/iotaledger/wasp/packages/vm/wasmclient"
)

func main() {
}

//export on_load
func OnLoad() {
	h := &wasmclient.WasmVMHost{}
	h.ConnectWasmHost()
	alphainterfacecontract.OnLoad()
}