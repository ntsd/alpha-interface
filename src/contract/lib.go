// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

//nolint:dupl
package alphainterfacecontract

import "github.com/iotaledger/wasp/packages/vm/wasmlib"

func OnLoad() {
	exports := wasmlib.NewScExports()
	exports.AddFunc(FuncInit, funcInitThunk)
	exports.AddFunc(FuncSetOwner, funcSetOwnerThunk)
	exports.AddView(ViewGetOwner, viewGetOwnerThunk)

	for i, key := range keyMap {
		idxMap[i] = key.KeyID()
	}
}

type InitContext struct {
	Params ImmutableInitParams
	State  MutableAlphaInterfaceContractState
}

func funcInitThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("alphainterfacecontract.funcInit")
	f := &InitContext{
		Params: ImmutableInitParams{
			id: wasmlib.OBJ_ID_PARAMS,
		},
		State: MutableAlphaInterfaceContractState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	funcInit(ctx, f)
	ctx.Log("alphainterfacecontract.funcInit ok")
}

type SetOwnerContext struct {
	Params ImmutableSetOwnerParams
	State  MutableAlphaInterfaceContractState
}

func funcSetOwnerThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("alphainterfacecontract.funcSetOwner")
	// current owner of this smart contract
	access := ctx.State().GetAgentID(wasmlib.Key("owner"))
	ctx.Require(access.Exists(), "access not set: owner")
	ctx.Require(ctx.Caller() == access.Value(), "no permission")

	f := &SetOwnerContext{
		Params: ImmutableSetOwnerParams{
			id: wasmlib.OBJ_ID_PARAMS,
		},
		State: MutableAlphaInterfaceContractState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	ctx.Require(f.Params.Owner().Exists(), "missing mandatory owner")
	funcSetOwner(ctx, f)
	ctx.Log("alphainterfacecontract.funcSetOwner ok")
}

type GetOwnerContext struct {
	Results MutableGetOwnerResults
	State   ImmutableAlphaInterfaceContractState
}

func viewGetOwnerThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("alphainterfacecontract.viewGetOwner")
	f := &GetOwnerContext{
		Results: MutableGetOwnerResults{
			id: wasmlib.OBJ_ID_RESULTS,
		},
		State: ImmutableAlphaInterfaceContractState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	viewGetOwner(ctx, f)
	ctx.Log("alphainterfacecontract.viewGetOwner ok")
}