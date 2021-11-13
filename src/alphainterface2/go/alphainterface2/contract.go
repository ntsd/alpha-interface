// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package alphainterface2

import "github.com/iotaledger/wasp/packages/vm/wasmlib/go/wasmlib"

type CloseOrderCall struct {
	Func   *wasmlib.ScFunc
	Params MutableCloseOrderParams
}

type CreateOrderCall struct {
	Func   *wasmlib.ScFunc
	Params MutableCreateOrderParams
}

type InitCall struct {
	Func   *wasmlib.ScInitFunc
	Params MutableInitParams
}

type SetCropCall struct {
	Func   *wasmlib.ScFunc
	Params MutableSetCropParams
}

type SetOwnerCall struct {
	Func   *wasmlib.ScFunc
	Params MutableSetOwnerParams
}

type SetWalletAmountCall struct {
	Func   *wasmlib.ScFunc
	Params MutableSetWalletAmountParams
}

type ViewGetMyWalletsCall struct {
	Func    *wasmlib.ScFunc
	Results ImmutableViewGetMyWalletsResults
}

type ViewGetOrdersCall struct {
	Func    *wasmlib.ScFunc
	Results ImmutableViewGetOrdersResults
}

type GetCropsCall struct {
	Func    *wasmlib.ScView
	Results ImmutableGetCropsResults
}

type GetOrdersCall struct {
	Func    *wasmlib.ScView
	Results ImmutableGetOrdersResults
}

type GetOwnerCall struct {
	Func    *wasmlib.ScView
	Results ImmutableGetOwnerResults
}

type GetWalletsCall struct {
	Func    *wasmlib.ScView
	Results ImmutableGetWalletsResults
}

type Funcs struct{}

var ScFuncs Funcs

func (sc Funcs) CloseOrder(ctx wasmlib.ScFuncCallContext) *CloseOrderCall {
	f := &CloseOrderCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncCloseOrder)}
	f.Func.SetPtrs(&f.Params.id, nil)
	return f
}

func (sc Funcs) CreateOrder(ctx wasmlib.ScFuncCallContext) *CreateOrderCall {
	f := &CreateOrderCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncCreateOrder)}
	f.Func.SetPtrs(&f.Params.id, nil)
	return f
}

func (sc Funcs) Init(ctx wasmlib.ScFuncCallContext) *InitCall {
	f := &InitCall{Func: wasmlib.NewScInitFunc(ctx, HScName, HFuncInit, keyMap[:], idxMap[:])}
	f.Func.SetPtrs(&f.Params.id, nil)
	return f
}

func (sc Funcs) SetCrop(ctx wasmlib.ScFuncCallContext) *SetCropCall {
	f := &SetCropCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncSetCrop)}
	f.Func.SetPtrs(&f.Params.id, nil)
	return f
}

func (sc Funcs) SetOwner(ctx wasmlib.ScFuncCallContext) *SetOwnerCall {
	f := &SetOwnerCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncSetOwner)}
	f.Func.SetPtrs(&f.Params.id, nil)
	return f
}

func (sc Funcs) SetWalletAmount(ctx wasmlib.ScFuncCallContext) *SetWalletAmountCall {
	f := &SetWalletAmountCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncSetWalletAmount)}
	f.Func.SetPtrs(&f.Params.id, nil)
	return f
}

func (sc Funcs) ViewGetMyWallets(ctx wasmlib.ScFuncCallContext) *ViewGetMyWalletsCall {
	f := &ViewGetMyWalletsCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncViewGetMyWallets)}
	f.Func.SetPtrs(nil, &f.Results.id)
	return f
}

func (sc Funcs) ViewGetOrders(ctx wasmlib.ScFuncCallContext) *ViewGetOrdersCall {
	f := &ViewGetOrdersCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncViewGetOrders)}
	f.Func.SetPtrs(nil, &f.Results.id)
	return f
}

func (sc Funcs) GetCrops(ctx wasmlib.ScViewCallContext) *GetCropsCall {
	f := &GetCropsCall{Func: wasmlib.NewScView(ctx, HScName, HViewGetCrops)}
	f.Func.SetPtrs(nil, &f.Results.id)
	return f
}

func (sc Funcs) GetOrders(ctx wasmlib.ScViewCallContext) *GetOrdersCall {
	f := &GetOrdersCall{Func: wasmlib.NewScView(ctx, HScName, HViewGetOrders)}
	f.Func.SetPtrs(nil, &f.Results.id)
	return f
}

func (sc Funcs) GetOwner(ctx wasmlib.ScViewCallContext) *GetOwnerCall {
	f := &GetOwnerCall{Func: wasmlib.NewScView(ctx, HScName, HViewGetOwner)}
	f.Func.SetPtrs(nil, &f.Results.id)
	return f
}

func (sc Funcs) GetWallets(ctx wasmlib.ScViewCallContext) *GetWalletsCall {
	f := &GetWalletsCall{Func: wasmlib.NewScView(ctx, HScName, HViewGetWallets)}
	f.Func.SetPtrs(nil, &f.Results.id)
	return f
}
