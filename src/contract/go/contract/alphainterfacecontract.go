// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package alphainterfacecontract

import (
	"github.com/iotaledger/wasp/packages/vm/wasmlib/go/wasmlib"
)

const (
	AMOUNT_SCALE = 1000000 // amount = unit * AMOUNT_SCALE

	ORDER_TYPE_SHORT = "short"
	ORDER_TYPE_LONG  = "long"

	POSITION_STATUS_OPENING   = "opening"
	POSITION_STATUS_CLOSING   = "closing"
	POSITION_STATUS_CLOSED    = "closed"
	POSITION_STATUS_LIQUIDATE = "liquidate"

	ORDER_STATUS_OPENING  = "opening"
	ORDER_STATUS_MATCHED  = "matched"
	ORDER_STATUS_CANCELED = "canceled"
)

func funcCloseOrder(ctx wasmlib.ScFuncContext, f *CloseOrderContext) {
	// TODO
}

func funcClosePosition(ctx wasmlib.ScFuncContext, f *ClosePositionContext) {
	caller := ctx.Caller()

	// get params
	positionIdx := f.Params.PositionIdx().Value()

	// get states
	orders := f.State.Orders()
	ordersLen := orders.Length()
	positions := f.State.Positions()
	position := positions.GetPosition(positionIdx).Value()
	crop := f.State.Crops().GetCrop(position.CropIdx).Value()

	defer positions.GetPosition(positionIdx).SetValue(position)

	// check owner position
	if position.Owner.Address() != caller.Address() {
		ctx.Log("you're not the owner of this position")
		return
	}

	// check status is opening
	if position.Status != POSITION_STATUS_OPENING {
		ctx.Log("the status is not opening")
		return
	}

	position.Status = POSITION_STATUS_CLOSING

	oppositeType := ORDER_TYPE_SHORT
	if position.Type == ORDER_TYPE_SHORT {
		oppositeType = ORDER_TYPE_LONG
	}

	newOrder := &Order{
		Idx:         ordersLen,
		PositionIdx: positionIdx,
		CropIdx:     position.CropIdx,
		CurAmount:   position.Amount,
		FullAmount:  position.Amount,
		Owner:       caller,
		Status:      ORDER_STATUS_OPENING,
		Type:        oppositeType,
	}
	orders.GetOrder(ordersLen).SetValue(newOrder)

	// find match orders
	for i := int32(0); i < ordersLen; i++ {
		adjOrder := orders.GetOrder(i).Value()
		// check match sell with long opening order in same crop
		if adjOrder.Status != ORDER_STATUS_OPENING ||
			adjOrder.CropIdx != newOrder.CropIdx ||
			adjOrder.Type == newOrder.Type {
			continue
		}

		adjPosition := positions.GetPosition(adjOrder.PositionIdx).Value()

		// adjust order amount more than new order
		if adjOrder.CurAmount > newOrder.CurAmount {
			// update adjust position
			updatePositionAmount(ctx, adjPosition, newOrder.CurAmount, crop.Yield)
			positions.GetPosition(adjOrder.PositionIdx).SetValue(adjPosition)

			// update adjust order
			adjOrder.CurAmount -= newOrder.CurAmount
			orders.GetOrder(i).SetValue(adjOrder)

			// update new position
			updatePositionAmount(ctx, position, newOrder.CurAmount, crop.Yield)

			// update new order
			newOrder.CurAmount = 0
			newOrder.Status = ORDER_STATUS_MATCHED

			return
		}
		// adjust order less or equal than new order

		// update new position
		updatePositionAmount(ctx, position, adjOrder.CurAmount, crop.Yield)

		// update new order
		newOrder.CurAmount -= adjOrder.CurAmount

		// update adjust position
		updatePositionAmount(ctx, adjPosition, adjOrder.CurAmount, crop.Yield)
		positions.GetPosition(adjOrder.PositionIdx).SetValue(adjPosition)

		// update adjust order
		adjOrder.CurAmount = 0
		adjOrder.Status = ORDER_STATUS_MATCHED
		orders.GetOrder(i).SetValue(adjOrder)
	}
}

func updatePositionAmount(ctx wasmlib.ScFuncContext, position *Position, amount int64, price int64) {
	if position.Status == POSITION_STATUS_CLOSING {
		iotaAmount := amount * price / AMOUNT_SCALE
		transfers := wasmlib.NewScTransferIotas(iotaAmount)
		ctx.TransferToAddress(position.Owner.Address(), transfers)

		position.Amount -= amount
		if position.Amount <= 0 {
			position.Status = POSITION_STATUS_CLOSED
		}
	}
	if position.Status == POSITION_STATUS_OPENING {
		position.Amount += amount
		position.AveragePrice = (position.AveragePrice + price) / 2
		if position.Amount >= 0 {
			position.Status = POSITION_STATUS_OPENING
		}
	}
}

func funcCreateOrder(ctx wasmlib.ScFuncContext, f *CreateOrderContext) {
	caller := ctx.Caller()

	incoming := ctx.Incoming()
	incomingAmount := incoming.Balance(wasmlib.IOTA)
	incomingAmountScaled := incomingAmount * AMOUNT_SCALE

	// get params
	cropIdx := f.Params.CropIdx().Value()
	orderType := f.Params.Type().Value()

	// get states
	orders := f.State.Orders()
	ordersLen := orders.Length()
	positions := f.State.Positions()
	positionsLen := positions.Length()
	crops := f.State.Crops()

	crop := crops.GetCrop(cropIdx).Value()

	amount := incomingAmountScaled / crop.Yield

	newPosition := &Position{
		CropIdx:      cropIdx,
		Amount:       0,
		AveragePrice: crop.Yield,
		Owner:        caller,
		Status:       POSITION_STATUS_OPENING,
		Type:         orderType,
	}
	defer positions.GetPosition(positionsLen).SetValue(newPosition)

	newOrder := &Order{
		CropIdx:     cropIdx,
		PositionIdx: positionsLen,
		CurAmount:   amount,
		FullAmount:  amount,
		Owner:       caller,
		Type:        orderType,
		Status:      ORDER_STATUS_OPENING,
	}
	defer orders.GetOrder(ordersLen).SetValue(newOrder)

	// find match orders
	for i := int32(0); i < ordersLen; i++ {
		adjOrder := orders.GetOrder(i).Value()
		// check match sell with long opening order in same crop
		if adjOrder.Status != ORDER_STATUS_OPENING ||
			adjOrder.CropIdx != newOrder.CropIdx ||
			adjOrder.Type == newOrder.Type {
			continue
		}

		adjPosition := positions.GetPosition(adjOrder.PositionIdx).Value()

		// adjust order amount more than new order
		if adjOrder.CurAmount > newOrder.CurAmount {
			// update adjust position
			updatePositionAmount(ctx, adjPosition, newOrder.CurAmount, crop.Yield)
			positions.GetPosition(adjOrder.PositionIdx).SetValue(adjPosition)

			// update adjust order
			adjOrder.CurAmount -= newOrder.CurAmount
			orders.GetOrder(i).SetValue(adjOrder)

			// update new position
			updatePositionAmount(ctx, newPosition, newOrder.CurAmount, crop.Yield)

			// update new order
			newOrder.CurAmount = 0
			newOrder.Status = ORDER_STATUS_MATCHED

			return
		}
		// adjust order less or equal than new order

		// update new position
		updatePositionAmount(ctx, newPosition, adjOrder.CurAmount, crop.Yield)

		// update new order
		newOrder.CurAmount -= adjOrder.CurAmount

		// update adjust position
		updatePositionAmount(ctx, adjPosition, adjOrder.CurAmount, crop.Yield)
		positions.GetPosition(adjOrder.PositionIdx).SetValue(adjPosition)

		// update adjust order
		adjOrder.CurAmount = 0
		adjOrder.Status = ORDER_STATUS_MATCHED
		orders.GetOrder(i).SetValue(adjOrder)
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
		crop.UpdatedAt = ctx.Timestamp()
		crops.GetCrop(i).SetValue(crop)

		// TOTO liquidate positions

		return
	}

	// if not found create a new crop
	newCrop := &Crop{
		Idx:       cropsLen,
		Name:      cropName,
		Country:   cropCountry,
		Yield:     cropYield,
		UpdatedAt: ctx.Timestamp(),
	}
	crops.GetCrop(cropsLen).SetValue(newCrop)
}

func funcSetOwner(ctx wasmlib.ScFuncContext, f *SetOwnerContext) {
	f.State.Owner().SetValue(f.Params.Owner().Value())
}

func viewGetCrop(ctx wasmlib.ScViewContext, f *GetCropContext) {
	cropIdx := f.Params.CropIdx().Value()
	f.Results.Crop().SetValue(f.State.Crops().GetCrop(cropIdx).Value())
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
	resultPositions := f.Results.Positions()
	statePositions := f.State.Positions()
	statePositionsLen := statePositions.Length()
	var n int32
	for i := int32(0); i < statePositionsLen; i++ {
		position := statePositions.GetPosition(i).Value()
		if ctx.Caller().Address() == position.Owner.Address() {
			resultPositions.GetPosition(n).SetValue(position)
			n++
		}
	}
}

func viewGetOrders(ctx wasmlib.ScViewContext, f *GetOrdersContext) {
	resultOrders := f.Results.Orders()
	stateOrders := f.State.Orders()
	stateOrdersLen := stateOrders.Length()
	for i := int32(0); i < stateOrdersLen; i++ {
		resultOrders.GetOrder(i).SetValue(stateOrders.GetOrder(i).Value())
	}
}

func viewGetOwner(ctx wasmlib.ScViewContext, f *GetOwnerContext) {
	f.Results.Owner().SetValue(f.State.Owner().Value())
}
