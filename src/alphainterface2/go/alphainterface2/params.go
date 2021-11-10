// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package alphainterface2
import "github.com/iotaledger/wasp/packages/vm/wasmlib/go/wasmlib"

type ImmutableCloseOrderParams struct {
	id int32
}

func (s ImmutableCloseOrderParams) OrderIdx() wasmlib.ScImmutableInt32 {
	return wasmlib.NewScImmutableInt32(s.id, idxMap[IdxParamOrderIdx])
}

type MutableCloseOrderParams struct {
	id int32
}

func (s MutableCloseOrderParams) OrderIdx() wasmlib.ScMutableInt32 {
	return wasmlib.NewScMutableInt32(s.id, idxMap[IdxParamOrderIdx])
}

type ImmutableCreateOrderParams struct {
	id int32
}

func (s ImmutableCreateOrderParams) Amount() wasmlib.ScImmutableInt64 {
	return wasmlib.NewScImmutableInt64(s.id, idxMap[IdxParamAmount])
}

func (s ImmutableCreateOrderParams) CropIdx() wasmlib.ScImmutableInt32 {
	return wasmlib.NewScImmutableInt32(s.id, idxMap[IdxParamCropIdx])
}

func (s ImmutableCreateOrderParams) Price() wasmlib.ScImmutableInt64 {
	return wasmlib.NewScImmutableInt64(s.id, idxMap[IdxParamPrice])
}

func (s ImmutableCreateOrderParams) Type() wasmlib.ScImmutableString {
	return wasmlib.NewScImmutableString(s.id, idxMap[IdxParamType])
}

type MutableCreateOrderParams struct {
	id int32
}

func (s MutableCreateOrderParams) Amount() wasmlib.ScMutableInt64 {
	return wasmlib.NewScMutableInt64(s.id, idxMap[IdxParamAmount])
}

func (s MutableCreateOrderParams) CropIdx() wasmlib.ScMutableInt32 {
	return wasmlib.NewScMutableInt32(s.id, idxMap[IdxParamCropIdx])
}

func (s MutableCreateOrderParams) Price() wasmlib.ScMutableInt64 {
	return wasmlib.NewScMutableInt64(s.id, idxMap[IdxParamPrice])
}

func (s MutableCreateOrderParams) Type() wasmlib.ScMutableString {
	return wasmlib.NewScMutableString(s.id, idxMap[IdxParamType])
}

type ImmutableInitParams struct {
	id int32
}

func (s ImmutableInitParams) Owner() wasmlib.ScImmutableAgentID {
	return wasmlib.NewScImmutableAgentID(s.id, idxMap[IdxParamOwner])
}

type MutableInitParams struct {
	id int32
}

func (s MutableInitParams) Owner() wasmlib.ScMutableAgentID {
	return wasmlib.NewScMutableAgentID(s.id, idxMap[IdxParamOwner])
}

type ImmutableSetCropParams struct {
	id int32
}

func (s ImmutableSetCropParams) Country() wasmlib.ScImmutableString {
	return wasmlib.NewScImmutableString(s.id, idxMap[IdxParamCountry])
}

func (s ImmutableSetCropParams) Name() wasmlib.ScImmutableString {
	return wasmlib.NewScImmutableString(s.id, idxMap[IdxParamName])
}

type MutableSetCropParams struct {
	id int32
}

func (s MutableSetCropParams) Country() wasmlib.ScMutableString {
	return wasmlib.NewScMutableString(s.id, idxMap[IdxParamCountry])
}

func (s MutableSetCropParams) Name() wasmlib.ScMutableString {
	return wasmlib.NewScMutableString(s.id, idxMap[IdxParamName])
}

type ImmutableSetOwnerParams struct {
	id int32
}

func (s ImmutableSetOwnerParams) Owner() wasmlib.ScImmutableAgentID {
	return wasmlib.NewScImmutableAgentID(s.id, idxMap[IdxParamOwner])
}

type MutableSetOwnerParams struct {
	id int32
}

func (s MutableSetOwnerParams) Owner() wasmlib.ScMutableAgentID {
	return wasmlib.NewScMutableAgentID(s.id, idxMap[IdxParamOwner])
}

type ImmutableGetCropParams struct {
	id int32
}

func (s ImmutableGetCropParams) CropIdx() wasmlib.ScImmutableInt32 {
	return wasmlib.NewScImmutableInt32(s.id, idxMap[IdxParamCropIdx])
}

type MutableGetCropParams struct {
	id int32
}

func (s MutableGetCropParams) CropIdx() wasmlib.ScMutableInt32 {
	return wasmlib.NewScMutableInt32(s.id, idxMap[IdxParamCropIdx])
}
