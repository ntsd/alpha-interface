// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package alphainterfacecontract

import (
	"strconv"

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

func viewGetPositions(ctx wasmlib.ScViewContext, f *GetPositionsContext) {
	resultPositions := f.Results.Positions()
	statePositions := f.State.Positions()
	statePositionsLen := statePositions.Length()
	for i := int32(0); i < statePositionsLen; i++ {
		position := statePositions.GetPosition(i).Value()
		resultPositions.GetPosition(i).SetValue(position)
	}
}

func viewGetCropsString(ctx wasmlib.ScViewContext, f *GetCropsStringContext) {
	stateCrops := f.State.Crops()
	stateCropsLen := stateCrops.Length()
	cropsString := "["
	for i := int32(0); i < stateCropsLen; i++ {
		stateCrop := stateCrops.GetCrop(i).Value()
		cropsString += "{Idx:" + strconv.FormatInt(int64(stateCrop.Idx), 10) + "," +
			"Name:" + stateCrop.Name + "," +
			"Country:" + stateCrop.Country + "," +
			"Yield:" + strconv.FormatInt(stateCrop.Yield, 10) + "," +
			"UpdatedAt:" + string(stateCrop.UpdatedAt) + "},"
	}
	cropsString += "]"
	f.Results.CropsString().SetValue(cropsString)
}

func viewGetOrdersString(ctx wasmlib.ScViewContext, f *GetOrdersStringContext) {
	stateOrders := f.State.Orders()
	stateOrdersLen := stateOrders.Length()
	ordersString := "["
	for i := int32(0); i < stateOrdersLen; i++ {
		stateOrder := stateOrders.GetOrder(i).Value()
		ordersString += "{Status:" + stateOrder.Status + "," +
			"Idx:" + strconv.FormatInt(int64(stateOrder.Idx), 10) + "," +
			"PositionIdx:" + strconv.FormatInt(int64(stateOrder.PositionIdx), 10) + "," +
			"CropIdx:" + strconv.FormatInt(int64(stateOrder.CropIdx), 10) + "," +
			"Type:" + stateOrder.Type + "," +
			"CurAmount:" + strconv.FormatInt(stateOrder.CurAmount, 10) + "," +
			"FullAmount:" + strconv.FormatInt(stateOrder.FullAmount, 10) + "," +
			"Owner.Address:" + stateOrder.Owner.Address().String() + "},"
	}
	ordersString += "]"
	f.Results.OrdersString().SetValue(ordersString)
}

func viewGetPositionsString(ctx wasmlib.ScViewContext, f *GetPositionsStringContext) {
	statePositions := f.State.Positions()
	statePositionsLen := statePositions.Length()
	positionsString := "["
	for i := int32(0); i < statePositionsLen; i++ {
		statePosition := statePositions.GetPosition(i).Value()
		positionsString += "{Status:" + statePosition.Status + "," +
			"Idx:" + strconv.FormatInt(int64(statePosition.Idx), 10) + "," +
			"CropIdx:" + strconv.FormatInt(int64(statePosition.CropIdx), 10) + "," +
			"Type:" + statePosition.Type + "," +
			"Amount:" + strconv.FormatInt(statePosition.Amount, 10) + "," +
			"AveragePrice:" + strconv.FormatInt(statePosition.AveragePrice, 10) + "," +
			"Owner.Address:" + statePosition.Owner.Address().String() + "},"
	}
	positionsString += "]"
	f.Results.PositionsString().SetValue(positionsString)
}
