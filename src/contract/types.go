// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package alphainterfacecontract

import "github.com/iotaledger/wasp/packages/vm/wasmlib"

type Crop struct {
	Country string // the country name
	Name    string // name of the crop
	Year    int64  // year of the report
	Yield   int64  // yield value of the crop per hg/ha
}

func NewCropFromBytes(bytes []byte) *Crop {
	decode := wasmlib.NewBytesDecoder(bytes)
	data := &Crop{}
	data.Country = decode.String()
	data.Name = decode.String()
	data.Year = decode.Int64()
	data.Yield = decode.Int64()
	decode.Close()
	return data
}

func (o *Crop) Bytes() []byte {
	return wasmlib.NewBytesEncoder().
		String(o.Country).
		String(o.Name).
		Int64(o.Year).
		Int64(o.Yield).
		Data()
}

type ImmutableCrop struct {
	objID int32
	keyID wasmlib.Key32
}

func (o ImmutableCrop) Exists() bool {
	return wasmlib.Exists(o.objID, o.keyID, wasmlib.TYPE_BYTES)
}

func (o ImmutableCrop) Value() *Crop {
	return NewCropFromBytes(wasmlib.GetBytes(o.objID, o.keyID, wasmlib.TYPE_BYTES))
}

type MutableCrop struct {
	objID int32
	keyID wasmlib.Key32
}

func (o MutableCrop) Exists() bool {
	return wasmlib.Exists(o.objID, o.keyID, wasmlib.TYPE_BYTES)
}

func (o MutableCrop) SetValue(value *Crop) {
	wasmlib.SetBytes(o.objID, o.keyID, wasmlib.TYPE_BYTES, value.Bytes())
}

func (o MutableCrop) Value() *Crop {
	return NewCropFromBytes(wasmlib.GetBytes(o.objID, o.keyID, wasmlib.TYPE_BYTES))
}

type Transaction struct {
	Amount   int64             // amount to transfer
	Receiver wasmlib.ScAgentID // agent id of receiver
	Sender   wasmlib.ScAgentID // agent id of sender
}

func NewTransactionFromBytes(bytes []byte) *Transaction {
	decode := wasmlib.NewBytesDecoder(bytes)
	data := &Transaction{}
	data.Amount = decode.Int64()
	data.Receiver = decode.AgentID()
	data.Sender = decode.AgentID()
	decode.Close()
	return data
}

func (o *Transaction) Bytes() []byte {
	return wasmlib.NewBytesEncoder().
		Int64(o.Amount).
		AgentID(o.Receiver).
		AgentID(o.Sender).
		Data()
}

type ImmutableTransaction struct {
	objID int32
	keyID wasmlib.Key32
}

func (o ImmutableTransaction) Exists() bool {
	return wasmlib.Exists(o.objID, o.keyID, wasmlib.TYPE_BYTES)
}

func (o ImmutableTransaction) Value() *Transaction {
	return NewTransactionFromBytes(wasmlib.GetBytes(o.objID, o.keyID, wasmlib.TYPE_BYTES))
}

type MutableTransaction struct {
	objID int32
	keyID wasmlib.Key32
}

func (o MutableTransaction) Exists() bool {
	return wasmlib.Exists(o.objID, o.keyID, wasmlib.TYPE_BYTES)
}

func (o MutableTransaction) SetValue(value *Transaction) {
	wasmlib.SetBytes(o.objID, o.keyID, wasmlib.TYPE_BYTES, value.Bytes())
}

func (o MutableTransaction) Value() *Transaction {
	return NewTransactionFromBytes(wasmlib.GetBytes(o.objID, o.keyID, wasmlib.TYPE_BYTES))
}
