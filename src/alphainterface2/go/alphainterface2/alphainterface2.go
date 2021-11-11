// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package alphainterface2

import "github.com/iotaledger/wasp/packages/vm/wasmlib/go/wasmlib"

func funcCloseOrder(ctx wasmlib.ScFuncContext, f *CloseOrderContext) {
}

func funcCreateOrder(ctx wasmlib.ScFuncContext, f *CreateOrderContext) {
}

func funcInit(ctx wasmlib.ScFuncContext, f *InitContext) {
    if f.Params.Owner().Exists() {
        f.State.Owner().SetValue(f.Params.Owner().Value())
        return
    }
    f.State.Owner().SetValue(ctx.ContractCreator())
}

func funcSetCrop(ctx wasmlib.ScFuncContext, f *SetCropContext) {
}

func funcSetOwner(ctx wasmlib.ScFuncContext, f *SetOwnerContext) {
}

func funcViewMyWallets(ctx wasmlib.ScFuncContext, f *ViewMyWalletsContext) {
}

func viewGetCrop(ctx wasmlib.ScViewContext, f *GetCropContext) {
}

func viewGetCrops(ctx wasmlib.ScViewContext, f *GetCropsContext) {
}

func viewGetCropsString(ctx wasmlib.ScViewContext, f *GetCropsStringContext) {
}

func viewGetOrders(ctx wasmlib.ScViewContext, f *GetOrdersContext) {
}

func viewGetOrdersString(ctx wasmlib.ScViewContext, f *GetOrdersStringContext) {
}

func viewGetOwner(ctx wasmlib.ScViewContext, f *GetOwnerContext) {
}

func viewViewWallets(ctx wasmlib.ScViewContext, f *ViewWalletsContext) {
}
