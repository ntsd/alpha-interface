// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package alphainterface2

import "github.com/iotaledger/wasp/packages/vm/wasmlib/go/wasmlib"

const (
	ORDER_TYPE_SELL = "sell"
	ORDER_TYPE_BUY  = "buy"

	ORDER_STATUS_OPENING  = "opening"
	ORDER_STATUS_MATCHED  = "matched"
	ORDER_STATUS_CANCELED = "canceled"
)

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func funcCloseOrder(ctx wasmlib.ScFuncContext, f *CloseOrderContext) {
	// get states
	orders := f.State.Orders()

	// get params
	orderId := f.Params.OrderIdx().Value()
	order := orders.GetOrder(orderId).Value()
	defer orders.GetOrder(orderId).SetValue(order)

	order.Status = ORDER_STATUS_CANCELED
	if order.Type == ORDER_TYPE_BUY {
		// tranfer iota back to order owner
		ctx.TransferToAddress(order.Owner.Address(), wasmlib.NewScTransferIotas(order.Iota))
		order.Iota = 0
	}
}

func findMatchedOrderAndAdjust(ctx wasmlib.ScFuncContext, f *CreateOrderContext, newOrder *Order, newOrderWallet *Wallet) {
	// get states
	orders := f.State.Orders()
	ordersLen := orders.Length()
	wallets := f.State.Wallets()

	// find match orders
	for i := int32(0); i < ordersLen; i++ {
		adjOrder := orders.GetOrder(i).Value()
		// check match sell with long opening order in same crop
		if adjOrder.Status != ORDER_STATUS_OPENING ||
			adjOrder.CropIdx != newOrder.CropIdx ||
			adjOrder.Type == newOrder.Type {
			continue
		}

		// validate offer price this is not solid because the adjust order doesn't sort by price
		if (newOrder.Type == ORDER_TYPE_BUY && newOrder.Price < adjOrder.Price) ||
			(newOrder.Type == ORDER_TYPE_SELL && newOrder.Price > adjOrder.Price) {
			continue
		}
		matchedPrice := adjOrder.Price
		matchedAmount := min(adjOrder.Amount, newOrder.Amount)

		newOrder.Amount -= matchedPrice
		if newOrder.Type == ORDER_TYPE_BUY {
			newOrderWallet.Amount += matchedAmount
			newOrder.Iota -= matchedAmount * matchedPrice
		} else {
			newOrderWallet.Amount -= matchedAmount
			returnIota := matchedAmount * matchedPrice
			ctx.TransferToAddress(newOrder.Owner.Address(), wasmlib.NewScTransferIotas(returnIota))
		}
		if newOrder.Amount <= 0 {
			newOrder.Status = ORDER_STATUS_MATCHED
			ctx.TransferToAddress(newOrder.Owner.Address(), wasmlib.NewScTransferIotas(newOrder.Iota))
			return
		}

		adjWallet := wallets.GetWallet(adjOrder.WalletIdx).Value()
		adjOrder.Amount -= matchedAmount
		if adjOrder.Type == ORDER_TYPE_BUY {
			adjWallet.Amount += matchedAmount
			adjOrder.Iota -= matchedAmount * matchedPrice
		} else {
			adjWallet.Amount -= matchedAmount
			returnIota := matchedAmount * matchedPrice
			ctx.TransferToAddress(adjOrder.Owner.Address(), wasmlib.NewScTransferIotas(returnIota))
		}
		if adjOrder.Amount <= 0 {
			adjOrder.Status = ORDER_STATUS_MATCHED
			ctx.TransferToAddress(adjOrder.Owner.Address(), wasmlib.NewScTransferIotas(adjOrder.Iota))
		}
		wallets.GetWallet(adjOrder.WalletIdx).SetValue(adjWallet)
		orders.GetOrder(adjOrder.Idx).SetValue(adjOrder)
	}
}

func funcCreateOrder(ctx wasmlib.ScFuncContext, f *CreateOrderContext) {
	caller := ctx.Caller()

	incoming := ctx.Incoming()
	incomingAmount := incoming.Balance(wasmlib.IOTA)

	// get params
	cropIdx := f.Params.CropIdx().Value()
	price := f.Params.Price().Value()
	amount := f.Params.Amount().Value()
	orderType := f.Params.Type().Value()

	// validate
	if (orderType != ORDER_TYPE_SELL && orderType != ORDER_TYPE_BUY) ||
		price*amount > incomingAmount {
		ctx.Log("Validate failed")
		ctx.TransferToAddress(caller.Address(), wasmlib.NewScTransferIotas(incomingAmount))
	}

	// get states
	orders := f.State.Orders()
	ordersLen := orders.Length()
	wallets := f.State.Wallets()
	walletsLen := wallets.Length()
	crops := f.State.Crops()
	crop := crops.GetCrop(cropIdx).Value()

	if crop == nil {
		ctx.Log("Crop not found")
		ctx.TransferToAddress(caller.Address(), wasmlib.NewScTransferIotas(incomingAmount))
	}

	// get owner wallet by crop id
	var newOrderWallet *Wallet
	for i := int32(0); i < walletsLen; i++ {
		w := wallets.GetWallet(i).Value()
		if w != nil && w.Owner.Address() == caller.Address() && w.CropIdx == cropIdx {
			newOrderWallet = w
		}
	}
	if newOrderWallet == nil { // new wallet if not found
		newOrderWallet = &Wallet{
			Idx:     walletsLen,
			Owner:   caller,
			CropIdx: cropIdx,
			Amount:  0,
		}
	}
	wallets.GetWallet(newOrderWallet.Idx).SetValue(newOrderWallet)

	newOrder := &Order{
		Idx:       ordersLen,
		CropIdx:   cropIdx,
		WalletIdx: newOrderWallet.Idx,
		Amount:    amount,
		Price:     price,
		Owner:     caller,
		Type:      orderType,
		Status:    ORDER_STATUS_OPENING,
		Iota:      incomingAmount,
	}
	defer orders.GetOrder(ordersLen).SetValue(newOrder)

	findMatchedOrderAndAdjust(ctx, f, newOrder, newOrderWallet)
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

	crops := f.State.Crops()
	cropsLen := crops.Length()
	for i := int32(0); i < cropsLen; i++ {
		crop := crops.GetCrop(i).Value()
		if crop.Name != cropName || crop.Country != cropCountry {
			continue
		}

		crops.GetCrop(i).SetValue(crop)
		return
	}

	// if not found create a new crop
	newCrop := &Crop{
		Idx:     cropsLen,
		Name:    cropName,
		Country: cropCountry,
	}
	crops.GetCrop(cropsLen).SetValue(newCrop)
}

func funcSetOwner(ctx wasmlib.ScFuncContext, f *SetOwnerContext) {
	f.State.Owner().SetValue(f.Params.Owner().Value())
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

func viewGetWallets(ctx wasmlib.ScViewContext, f *GetWalletsContext) {
	resultWallets := f.Results.Wallets()
	stateWallets := f.State.Wallets()
	stateWalletsLen := stateWallets.Length()
	for i := int32(0); i < stateWalletsLen; i++ {
		resultWallets.GetWallet(i).SetValue(stateWallets.GetWallet(i).Value())
	}
}

func funcViewGetMyWallets(ctx wasmlib.ScFuncContext, f *ViewGetMyWalletsContext) {
	resultWallets := f.Results.Wallets()
	stateWallets := f.State.Wallets()
	stateWalletsLen := stateWallets.Length()
	var n int32
	for i := int32(0); i < stateWalletsLen; i++ {
		wallet := stateWallets.GetWallet(i).Value()
		if wallet.Owner.Address() == ctx.Caller().Address() {
			resultWallets.GetWallet(n).SetValue(wallet)
			n++
		}
	}
}

func funcViewGetOrders(ctx wasmlib.ScFuncContext, f *ViewGetOrdersContext) {
	resultOrders := f.Results.Orders()
	stateOrders := f.State.Orders()
	stateOrdersLen := stateOrders.Length()
	var n int32
	for i := int32(0); i < stateOrdersLen; i++ {
		wallet := stateOrders.GetOrder(i).Value()
		if wallet.Owner.Address() == ctx.Caller().Address() {
			resultOrders.GetOrder(n).SetValue(wallet)
			n++
		}
	}
}
