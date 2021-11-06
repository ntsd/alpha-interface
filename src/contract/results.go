// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package alphainterfacecontract

import "github.com/iotaledger/wasp/packages/vm/wasmlib"

type ImmutableGetCropResults struct {
	id int32
}

func (s ImmutableGetCropResults) Crop() ImmutableCrop {
	return ImmutableCrop{objID: s.id, keyID: idxMap[IdxResultCrop]}
}

type MutableGetCropResults struct {
	id int32
}

func (s MutableGetCropResults) Crop() MutableCrop {
	return MutableCrop{objID: s.id, keyID: idxMap[IdxResultCrop]}
}

type ArrayOfImmutableCrop struct {
	objID int32
}

func (a ArrayOfImmutableCrop) Length() int32 {
	return wasmlib.GetLength(a.objID)
}

func (a ArrayOfImmutableCrop) GetCrop(index int32) ImmutableCrop {
	return ImmutableCrop{objID: a.objID, keyID: wasmlib.Key32(index)}
}

type ImmutableGetCropsResults struct {
	id int32
}

func (s ImmutableGetCropsResults) Crops() ArrayOfImmutableCrop {
	arrID := wasmlib.GetObjectID(s.id, idxMap[IdxResultCrops], wasmlib.TYPE_ARRAY|wasmlib.TYPE_BYTES)
	return ArrayOfImmutableCrop{objID: arrID}
}

type ArrayOfMutableCrop struct {
	objID int32
}

func (a ArrayOfMutableCrop) Clear() {
	wasmlib.Clear(a.objID)
}

func (a ArrayOfMutableCrop) Length() int32 {
	return wasmlib.GetLength(a.objID)
}

func (a ArrayOfMutableCrop) GetCrop(index int32) MutableCrop {
	return MutableCrop{objID: a.objID, keyID: wasmlib.Key32(index)}
}

type MutableGetCropsResults struct {
	id int32
}

func (s MutableGetCropsResults) Crops() ArrayOfMutableCrop {
	arrID := wasmlib.GetObjectID(s.id, idxMap[IdxResultCrops], wasmlib.TYPE_ARRAY|wasmlib.TYPE_BYTES)
	return ArrayOfMutableCrop{objID: arrID}
}

type ArrayOfImmutablePosition struct {
	objID int32
}

func (a ArrayOfImmutablePosition) Length() int32 {
	return wasmlib.GetLength(a.objID)
}

func (a ArrayOfImmutablePosition) GetPosition(index int32) ImmutablePosition {
	return ImmutablePosition{objID: a.objID, keyID: wasmlib.Key32(index)}
}

type ImmutableGetMyPositionsResults struct {
	id int32
}

func (s ImmutableGetMyPositionsResults) Positions() ArrayOfImmutablePosition {
	arrID := wasmlib.GetObjectID(s.id, idxMap[IdxResultPositions], wasmlib.TYPE_ARRAY|wasmlib.TYPE_BYTES)
	return ArrayOfImmutablePosition{objID: arrID}
}

type ArrayOfMutablePosition struct {
	objID int32
}

func (a ArrayOfMutablePosition) Clear() {
	wasmlib.Clear(a.objID)
}

func (a ArrayOfMutablePosition) Length() int32 {
	return wasmlib.GetLength(a.objID)
}

func (a ArrayOfMutablePosition) GetPosition(index int32) MutablePosition {
	return MutablePosition{objID: a.objID, keyID: wasmlib.Key32(index)}
}

type MutableGetMyPositionsResults struct {
	id int32
}

func (s MutableGetMyPositionsResults) Positions() ArrayOfMutablePosition {
	arrID := wasmlib.GetObjectID(s.id, idxMap[IdxResultPositions], wasmlib.TYPE_ARRAY|wasmlib.TYPE_BYTES)
	return ArrayOfMutablePosition{objID: arrID}
}

type ArrayOfImmutableOrder struct {
	objID int32
}

func (a ArrayOfImmutableOrder) Length() int32 {
	return wasmlib.GetLength(a.objID)
}

func (a ArrayOfImmutableOrder) GetOrder(index int32) ImmutableOrder {
	return ImmutableOrder{objID: a.objID, keyID: wasmlib.Key32(index)}
}

type ImmutableGetOrdersResults struct {
	id int32
}

func (s ImmutableGetOrdersResults) Orders() ArrayOfImmutableOrder {
	arrID := wasmlib.GetObjectID(s.id, idxMap[IdxResultOrders], wasmlib.TYPE_ARRAY|wasmlib.TYPE_BYTES)
	return ArrayOfImmutableOrder{objID: arrID}
}

type ArrayOfMutableOrder struct {
	objID int32
}

func (a ArrayOfMutableOrder) Clear() {
	wasmlib.Clear(a.objID)
}

func (a ArrayOfMutableOrder) Length() int32 {
	return wasmlib.GetLength(a.objID)
}

func (a ArrayOfMutableOrder) GetOrder(index int32) MutableOrder {
	return MutableOrder{objID: a.objID, keyID: wasmlib.Key32(index)}
}

type MutableGetOrdersResults struct {
	id int32
}

func (s MutableGetOrdersResults) Orders() ArrayOfMutableOrder {
	arrID := wasmlib.GetObjectID(s.id, idxMap[IdxResultOrders], wasmlib.TYPE_ARRAY|wasmlib.TYPE_BYTES)
	return ArrayOfMutableOrder{objID: arrID}
}

type ImmutableGetOwnerResults struct {
	id int32
}

func (s ImmutableGetOwnerResults) Owner() wasmlib.ScImmutableAgentID {
	return wasmlib.NewScImmutableAgentID(s.id, idxMap[IdxResultOwner])
}

type MutableGetOwnerResults struct {
	id int32
}

func (s MutableGetOwnerResults) Owner() wasmlib.ScMutableAgentID {
	return wasmlib.NewScMutableAgentID(s.id, idxMap[IdxResultOwner])
}
