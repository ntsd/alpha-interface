// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package alphainterfacecontract

import "github.com/iotaledger/wasp/packages/vm/wasmlib/go/wasmlib"

const (
	IdxParamAmount     = 0
	IdxParamCountry    = 1
	IdxParamCropID     = 2
	IdxParamLeverage   = 3
	IdxParamName       = 4
	IdxParamOrderID    = 5
	IdxParamOwner      = 6
	IdxParamPositionID = 7
	IdxParamType       = 8
	IdxParamYield      = 9
	IdxResultCrop      = 10
	IdxResultCrops     = 11
	IdxResultOrders    = 12
	IdxResultOwner     = 13
	IdxResultPositions = 14
	IdxStateCrops      = 15
	IdxStateOrders     = 16
	IdxStateOwner      = 17
	IdxStatePositions  = 18
)

const keyMapLen = 19

var keyMap = [keyMapLen]wasmlib.Key{
	ParamAmount,
	ParamCountry,
	ParamCropID,
	ParamLeverage,
	ParamName,
	ParamOrderID,
	ParamOwner,
	ParamPositionID,
	ParamType,
	ParamYield,
	ResultCrop,
	ResultCrops,
	ResultOrders,
	ResultOwner,
	ResultPositions,
	StateCrops,
	StateOrders,
	StateOwner,
	StatePositions,
}

var idxMap [keyMapLen]wasmlib.Key32
