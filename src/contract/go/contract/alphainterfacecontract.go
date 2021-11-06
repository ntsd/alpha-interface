// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package alphainterfacecontract

import (
	"fmt"
	"time"

	"github.com/iotaledger/wasp/packages/vm/wasmlib/go/wasmlib"
)

func funcCloseOrder(ctx wasmlib.ScFuncContext, f *CloseOrderContext) {
}

func funcClosePosition(ctx wasmlib.ScFuncContext, f *ClosePositionContext) {
}

func funcCreateOrder(ctx wasmlib.ScFuncContext, f *CreateOrderContext) {
	// get params
	incoming := ctx.Incoming()
	amount := incoming.Balance(wasmlib.IOTA)
	amount = amount

	cropID := f.Params.CropID().Value()
	leverage := f.Params.Leverage().Value()
	leverage = leverage
	orderType := f.Params.Type().Value()

	// find matched order
	orders := f.State.Orders()
	ordersLen := orders.Length()
	for i := int32(0); i < ordersLen; i++ {
		order := orders.GetOrder(i).Value()
		if order.CropID != cropID || order.Type == orderType {
			continue
		}

		return
	}
}

func funcInit(ctx wasmlib.ScFuncContext, f *InitContext) {
	if f.Params.Owner().Exists() {
		f.State.Owner().SetValue(f.Params.Owner().Value())
		return
	}
	f.State.Owner().SetValue(ctx.ContractCreator())
}

func funcSetCrop(ctx wasmlib.ScFuncContext, f *SetCropContext) {
	// get params
	cropName := f.Params.Name().Value()
	cropCountry := f.Params.Country().Value()
	cropYield := f.Params.Yield().Value()

	crops := f.State.Crops()
	cropsLen := crops.Length()
	for i := int32(0); i < cropsLen; i++ {
		crop := crops.GetCrop(i).Value()
		if crop.Name != cropName || crop.Country != cropCountry {
			continue
		}

		crop.Yield = cropYield
		crop.UpdatedAt = time.Now().Unix()
		crops.GetCrop(i).SetValue(crop)

		// TOTO liquidated positions

		return
	}

	// if not found create a new crop
	newCrop := &Crop{
		Id:        fmt.Sprintf("%s_%s", cropName, cropCountry),
		Name:      cropName,
		Country:   cropCountry,
		Yield:     cropYield,
		UpdatedAt: time.Now().Unix(),
	}
	crops.GetCrop(cropsLen).SetValue(newCrop)
}

func funcSetOwner(ctx wasmlib.ScFuncContext, f *SetOwnerContext) {
	f.State.Owner().SetValue(f.Params.Owner().Value())
}

func viewGetOwner(ctx wasmlib.ScViewContext, f *GetOwnerContext) {
	f.Results.Owner().SetValue(f.State.Owner().Value())
}

func viewGetCrops(ctx wasmlib.ScViewContext, f *GetCropsContext) {
	resultCrops := f.Results.Crops()
	stateCrops := f.State.Crops()
	stateCropsLen := stateCrops.Length()
	for i := int32(0); i < stateCropsLen; i++ {
		resultCrops.GetCrop(i).SetValue(stateCrops.GetCrop(i).Value())
	}
}

func viewGetMyPositions(ctx wasmlib.ScViewContext, f *GetMyPositionsContext) {
}

func viewGetOrders(ctx wasmlib.ScViewContext, f *GetOrdersContext) {
}

func viewGetCrop(ctx wasmlib.ScViewContext, f *GetCropContext) {
	stateCrops := f.State.Crops()
	stateCrop := stateCrops.GetCrop(0).Value()
	f.Results.Crop().SetValue(stateCrop)
	// cropID := f.Params.CropID().Value()
	// stateCrops := f.State.Crops()
	// stateCropsLen := stateCrops.Length()
	// for i := int32(0); i < stateCropsLen; i++ {
	// 	stateCrop := stateCrops.GetCrop(i).Value()
	// 	if stateCrop.Id == cropID {
	// 		f.Results.Crop().SetValue(stateCrop)
	// 		return
	// 	}
	// }
}
