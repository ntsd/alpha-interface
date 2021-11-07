// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package alphainterfacecontract

import "github.com/iotaledger/wasp/packages/vm/wasmlib/go/wasmlib"

const (
	ScName        = "alphainterfacecontract"
	ScDescription = "AlphaInterfaceContract description"
	HScName       = wasmlib.ScHname(0x584a6111)
)

const (
	ParamCountry     = wasmlib.Key("country")
	ParamCropIdx     = wasmlib.Key("cropIdx")
	ParamName        = wasmlib.Key("name")
	ParamOrderIdx    = wasmlib.Key("orderIdx")
	ParamOwner       = wasmlib.Key("owner")
	ParamPositionIdx = wasmlib.Key("positionIdx")
	ParamType        = wasmlib.Key("type")
	ParamYield       = wasmlib.Key("yield")
)

const (
	ResultCrop            = wasmlib.Key("crop")
	ResultCrops           = wasmlib.Key("crops")
	ResultCropsString     = wasmlib.Key("cropsString")
	ResultOrders          = wasmlib.Key("orders")
	ResultOrdersString    = wasmlib.Key("ordersString")
	ResultOwner           = wasmlib.Key("owner")
	ResultPositions       = wasmlib.Key("positions")
	ResultPositionsString = wasmlib.Key("positionsString")
)

const (
	StateCrops     = wasmlib.Key("crops")
	StateOrders    = wasmlib.Key("orders")
	StateOwner     = wasmlib.Key("owner")
	StatePositions = wasmlib.Key("positions")
)

const (
	FuncCloseOrder         = "closeOrder"
	FuncClosePosition      = "closePosition"
	FuncCreateOrder        = "createOrder"
	FuncInit               = "init"
	FuncSetCrop            = "setCrop"
	FuncSetOwner           = "setOwner"
	ViewGetCrop            = "getCrop"
	ViewGetCrops           = "getCrops"
	ViewGetCropsString     = "getCropsString"
	ViewGetOrders          = "getOrders"
	ViewGetOrdersString    = "getOrdersString"
	ViewGetOwner           = "getOwner"
	ViewGetPositions       = "getPositions"
	ViewGetPositionsString = "getPositionsString"
)

const (
	HFuncCloseOrder         = wasmlib.ScHname(0xfc422a52)
	HFuncClosePosition      = wasmlib.ScHname(0x90e90210)
	HFuncCreateOrder        = wasmlib.ScHname(0xe3c7ac26)
	HFuncInit               = wasmlib.ScHname(0x1f44d644)
	HFuncSetCrop            = wasmlib.ScHname(0x321b14e9)
	HFuncSetOwner           = wasmlib.ScHname(0x2a15fe7b)
	HViewGetCrop            = wasmlib.ScHname(0x59d6b0d1)
	HViewGetCrops           = wasmlib.ScHname(0x0f16dbb7)
	HViewGetCropsString     = wasmlib.ScHname(0xb9ae6a38)
	HViewGetOrders          = wasmlib.ScHname(0x53738744)
	HViewGetOrdersString    = wasmlib.ScHname(0x622a0440)
	HViewGetOwner           = wasmlib.ScHname(0x137107a6)
	HViewGetPositions       = wasmlib.ScHname(0x0a766ffb)
	HViewGetPositionsString = wasmlib.ScHname(0xbb02feef)
)
