// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package alphainterfacecontract

import "github.com/iotaledger/wasp/packages/vm/wasmlib"

func funcInit(ctx wasmlib.ScFuncContext, f *InitContext) {
	if f.Params.Owner().Exists() {
		f.State.Owner().SetValue(f.Params.Owner().Value())
		return
	}
	f.State.Owner().SetValue(ctx.ContractCreator())
}

func funcSetOwner(ctx wasmlib.ScFuncContext, f *SetOwnerContext) {
	f.State.Owner().SetValue(f.Params.Owner().Value())
}

func viewGetOwner(ctx wasmlib.ScViewContext, f *GetOwnerContext) {
	f.Results.Owner().SetValue(f.State.Owner().Value())
}

func funcTransfer(ctx wasmlib.ScFuncContext, f *TransferContext) {
	incoming := ctx.Incoming()
	amount := incoming.Balance(wasmlib.IOTA)

	receiver := f.Params.Receiver().Value()

	tx := &Transaction{
		Sender:   ctx.Caller(),
		Amount:   amount,
		Receiver: receiver,
	}

	// Create transaction
	txs := f.State.Transactions()
	txsLen := txs.Length()
	txs.GetTransaction(txsLen).SetValue(tx)

	// Tranfer iotas
	transfers := wasmlib.NewScTransferIotas(amount)
	ctx.TransferToAddress(receiver.Address(), transfers)
}

func viewViewTransactions(ctx wasmlib.ScViewContext, f *ViewTransactionsContext) {
}
