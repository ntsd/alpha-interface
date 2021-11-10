// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package alphainterface2

import "github.com/iotaledger/wasp/packages/vm/wasmlib/go/wasmlib"

const (
	ScName        = "alphainterface2"
	ScDescription = "alphainterface2 description"
	HScName       = wasmlib.ScHname(0x426940a3)
)

const (
	ParamAmount   = wasmlib.Key("amount")
	ParamCountry  = wasmlib.Key("country")
	ParamCropIdx  = wasmlib.Key("cropIdx")
	ParamName     = wasmlib.Key("name")
	ParamOrderIdx = wasmlib.Key("orderIdx")
	ParamOwner    = wasmlib.Key("owner")
	ParamPrice    = wasmlib.Key("price")
	ParamType     = wasmlib.Key("type")
)

const (
	ResultCrop         = wasmlib.Key("crop")
	ResultCrops        = wasmlib.Key("crops")
	ResultCropsString  = wasmlib.Key("cropsString")
	ResultOrders       = wasmlib.Key("orders")
	ResultOrdersString = wasmlib.Key("ordersString")
	ResultOwner        = wasmlib.Key("owner")
	ResultWallets      = wasmlib.Key("wallets")
)

const (
	StateCrops   = wasmlib.Key("crops")
	StateOrders  = wasmlib.Key("orders")
	StateOwner   = wasmlib.Key("owner")
	StateWallets = wasmlib.Key("wallets")
)

const (
	FuncCloseOrder      = "closeOrder"
	FuncCreateOrder     = "createOrder"
	FuncInit            = "init"
	FuncSetCrop         = "setCrop"
	FuncSetOwner        = "setOwner"
	FuncViewMyWallets   = "viewMyWallets"
	ViewGetCrop         = "getCrop"
	ViewGetCrops        = "getCrops"
	ViewGetCropsString  = "getCropsString"
	ViewGetOrders       = "getOrders"
	ViewGetOrdersString = "getOrdersString"
	ViewGetOwner        = "getOwner"
	ViewViewWallets     = "viewWallets"
)

const (
	HFuncCloseOrder      = wasmlib.ScHname(0xfc422a52)
	HFuncCreateOrder     = wasmlib.ScHname(0xe3c7ac26)
	HFuncInit            = wasmlib.ScHname(0x1f44d644)
	HFuncSetCrop         = wasmlib.ScHname(0x321b14e9)
	HFuncSetOwner        = wasmlib.ScHname(0x2a15fe7b)
	HFuncViewMyWallets   = wasmlib.ScHname(0x0c2f77f3)
	HViewGetCrop         = wasmlib.ScHname(0x59d6b0d1)
	HViewGetCrops        = wasmlib.ScHname(0x0f16dbb7)
	HViewGetCropsString  = wasmlib.ScHname(0xb9ae6a38)
	HViewGetOrders       = wasmlib.ScHname(0x53738744)
	HViewGetOrdersString = wasmlib.ScHname(0x622a0440)
	HViewGetOwner        = wasmlib.ScHname(0x137107a6)
	HViewViewWallets     = wasmlib.ScHname(0x8c75313a)
)
