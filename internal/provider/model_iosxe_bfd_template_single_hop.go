// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"regexp"
	"strconv"

	"github.com/CiscoDevNet/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type BFDTemplateSingleHop struct {
	Device                                types.String `tfsdk:"device"`
	Id                                    types.String `tfsdk:"id"`
	Name                                  types.String `tfsdk:"name"`
	AuthenticationMd5Keychain             types.String `tfsdk:"authentication_md5_keychain"`
	AuthenticationMeticulousMd5Keychain   types.String `tfsdk:"authentication_meticulous_md5_keychain"`
	AuthenticationMeticulousSha1Keychain  types.String `tfsdk:"authentication_meticulous_sha_1_keychain"`
	AuthenticationSha1Keychain            types.String `tfsdk:"authentication_sha_1_keychain"`
	IntervalSinglehopV2MillUnitMinTx      types.Int64  `tfsdk:"interval_singlehop_v2_mill_unit_min_tx"`
	IntervalSinglehopV2MillUnitMinRx      types.Int64  `tfsdk:"interval_singlehop_v2_mill_unit_min_rx"`
	IntervalSinglehopV2MillUnitBoth       types.Int64  `tfsdk:"interval_singlehop_v2_mill_unit_both"`
	IntervalSinglehopV2MillUnitMultiplier types.Int64  `tfsdk:"interval_singlehop_v2_mill_unit_multiplier"`
	IntervalSinglehopV2MsUnitMinRx        types.Int64  `tfsdk:"interval_singlehop_v2_ms_unit_min_rx"`
	IntervalSinglehopV2MsUnitMinTx        types.Int64  `tfsdk:"interval_singlehop_v2_ms_unit_min_tx"`
	Echo                                  types.Bool   `tfsdk:"echo"`
	DampeningHalfTime                     types.Int64  `tfsdk:"dampening_half_time"`
	DampeningUnsuppressTime               types.Int64  `tfsdk:"dampening_unsuppress_time"`
	DampeningSuppressTime                 types.Int64  `tfsdk:"dampening_suppress_time"`
	DampeningMaxSuppressingTime           types.Int64  `tfsdk:"dampening_max_suppressing_time"`
}

type BFDTemplateSingleHopData struct {
	Device                                types.String `tfsdk:"device"`
	Id                                    types.String `tfsdk:"id"`
	Name                                  types.String `tfsdk:"name"`
	AuthenticationMd5Keychain             types.String `tfsdk:"authentication_md5_keychain"`
	AuthenticationMeticulousMd5Keychain   types.String `tfsdk:"authentication_meticulous_md5_keychain"`
	AuthenticationMeticulousSha1Keychain  types.String `tfsdk:"authentication_meticulous_sha_1_keychain"`
	AuthenticationSha1Keychain            types.String `tfsdk:"authentication_sha_1_keychain"`
	IntervalSinglehopV2MillUnitMinTx      types.Int64  `tfsdk:"interval_singlehop_v2_mill_unit_min_tx"`
	IntervalSinglehopV2MillUnitMinRx      types.Int64  `tfsdk:"interval_singlehop_v2_mill_unit_min_rx"`
	IntervalSinglehopV2MillUnitBoth       types.Int64  `tfsdk:"interval_singlehop_v2_mill_unit_both"`
	IntervalSinglehopV2MillUnitMultiplier types.Int64  `tfsdk:"interval_singlehop_v2_mill_unit_multiplier"`
	IntervalSinglehopV2MsUnitMinRx        types.Int64  `tfsdk:"interval_singlehop_v2_ms_unit_min_rx"`
	IntervalSinglehopV2MsUnitMinTx        types.Int64  `tfsdk:"interval_singlehop_v2_ms_unit_min_tx"`
	Echo                                  types.Bool   `tfsdk:"echo"`
	DampeningHalfTime                     types.Int64  `tfsdk:"dampening_half_time"`
	DampeningUnsuppressTime               types.Int64  `tfsdk:"dampening_unsuppress_time"`
	DampeningSuppressTime                 types.Int64  `tfsdk:"dampening_suppress_time"`
	DampeningMaxSuppressingTime           types.Int64  `tfsdk:"dampening_max_suppressing_time"`
}

func (data BFDTemplateSingleHop) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/bfd-template/Cisco-IOS-XE-bfd:single-hop=%v", url.QueryEscape(fmt.Sprintf("%v", data.Name.ValueString())))
}

func (data BFDTemplateSingleHopData) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/bfd-template/Cisco-IOS-XE-bfd:single-hop=%v", url.QueryEscape(fmt.Sprintf("%v", data.Name.ValueString())))
}

// if last path element has a key -> remove it
func (data BFDTemplateSingleHop) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data BFDTemplateSingleHop) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"name", data.Name.ValueString())
	}
	if !data.AuthenticationMd5Keychain.IsNull() && !data.AuthenticationMd5Keychain.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"authentication.md5.keychain", data.AuthenticationMd5Keychain.ValueString())
	}
	if !data.AuthenticationMeticulousMd5Keychain.IsNull() && !data.AuthenticationMeticulousMd5Keychain.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"authentication.meticulous-md5.keychain", data.AuthenticationMeticulousMd5Keychain.ValueString())
	}
	if !data.AuthenticationMeticulousSha1Keychain.IsNull() && !data.AuthenticationMeticulousSha1Keychain.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"authentication.meticulous-sha-1.keychain", data.AuthenticationMeticulousSha1Keychain.ValueString())
	}
	if !data.AuthenticationSha1Keychain.IsNull() && !data.AuthenticationSha1Keychain.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"authentication.sha-1.keychain", data.AuthenticationSha1Keychain.ValueString())
	}
	if !data.IntervalSinglehopV2MillUnitMinTx.IsNull() && !data.IntervalSinglehopV2MillUnitMinTx.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"interval-singlehop-v2.mill-unit.min-tx", strconv.FormatInt(data.IntervalSinglehopV2MillUnitMinTx.ValueInt64(), 10))
	}
	if !data.IntervalSinglehopV2MillUnitMinRx.IsNull() && !data.IntervalSinglehopV2MillUnitMinRx.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"interval-singlehop-v2.mill-unit.min-rx", strconv.FormatInt(data.IntervalSinglehopV2MillUnitMinRx.ValueInt64(), 10))
	}
	if !data.IntervalSinglehopV2MillUnitBoth.IsNull() && !data.IntervalSinglehopV2MillUnitBoth.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"interval-singlehop-v2.mill-unit.both", strconv.FormatInt(data.IntervalSinglehopV2MillUnitBoth.ValueInt64(), 10))
	}
	if !data.IntervalSinglehopV2MillUnitMultiplier.IsNull() && !data.IntervalSinglehopV2MillUnitMultiplier.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"interval-singlehop-v2.mill-unit.multiplier", strconv.FormatInt(data.IntervalSinglehopV2MillUnitMultiplier.ValueInt64(), 10))
	}
	if !data.IntervalSinglehopV2MsUnitMinRx.IsNull() && !data.IntervalSinglehopV2MsUnitMinRx.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"interval-singlehop-v2.ms-unit.min-rx", strconv.FormatInt(data.IntervalSinglehopV2MsUnitMinRx.ValueInt64(), 10))
	}
	if !data.IntervalSinglehopV2MsUnitMinTx.IsNull() && !data.IntervalSinglehopV2MsUnitMinTx.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"interval-singlehop-v2.ms-unit.min-tx", strconv.FormatInt(data.IntervalSinglehopV2MsUnitMinTx.ValueInt64(), 10))
	}
	if !data.Echo.IsNull() && !data.Echo.IsUnknown() {
		if data.Echo.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"echo", map[string]string{})
		}
	}
	if !data.DampeningHalfTime.IsNull() && !data.DampeningHalfTime.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"dampening.half-time", strconv.FormatInt(data.DampeningHalfTime.ValueInt64(), 10))
	}
	if !data.DampeningUnsuppressTime.IsNull() && !data.DampeningUnsuppressTime.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"dampening.unsuppress-time", strconv.FormatInt(data.DampeningUnsuppressTime.ValueInt64(), 10))
	}
	if !data.DampeningSuppressTime.IsNull() && !data.DampeningSuppressTime.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"dampening.suppress-time", strconv.FormatInt(data.DampeningSuppressTime.ValueInt64(), 10))
	}
	if !data.DampeningMaxSuppressingTime.IsNull() && !data.DampeningMaxSuppressingTime.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"dampening.max-suppressing-time", strconv.FormatInt(data.DampeningMaxSuppressingTime.ValueInt64(), 10))
	}
	return body
}

func (data *BFDTemplateSingleHop) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get(prefix + "authentication.md5.keychain"); value.Exists() && !data.AuthenticationMd5Keychain.IsNull() {
		data.AuthenticationMd5Keychain = types.StringValue(value.String())
	} else {
		data.AuthenticationMd5Keychain = types.StringNull()
	}
	if value := res.Get(prefix + "authentication.meticulous-md5.keychain"); value.Exists() && !data.AuthenticationMeticulousMd5Keychain.IsNull() {
		data.AuthenticationMeticulousMd5Keychain = types.StringValue(value.String())
	} else {
		data.AuthenticationMeticulousMd5Keychain = types.StringNull()
	}
	if value := res.Get(prefix + "authentication.meticulous-sha-1.keychain"); value.Exists() && !data.AuthenticationMeticulousSha1Keychain.IsNull() {
		data.AuthenticationMeticulousSha1Keychain = types.StringValue(value.String())
	} else {
		data.AuthenticationMeticulousSha1Keychain = types.StringNull()
	}
	if value := res.Get(prefix + "authentication.sha-1.keychain"); value.Exists() && !data.AuthenticationSha1Keychain.IsNull() {
		data.AuthenticationSha1Keychain = types.StringValue(value.String())
	} else {
		data.AuthenticationSha1Keychain = types.StringNull()
	}
	if value := res.Get(prefix + "interval-singlehop-v2.mill-unit.min-tx"); value.Exists() && !data.IntervalSinglehopV2MillUnitMinTx.IsNull() {
		data.IntervalSinglehopV2MillUnitMinTx = types.Int64Value(value.Int())
	} else {
		data.IntervalSinglehopV2MillUnitMinTx = types.Int64Null()
	}
	if value := res.Get(prefix + "interval-singlehop-v2.mill-unit.min-rx"); value.Exists() && !data.IntervalSinglehopV2MillUnitMinRx.IsNull() {
		data.IntervalSinglehopV2MillUnitMinRx = types.Int64Value(value.Int())
	} else {
		data.IntervalSinglehopV2MillUnitMinRx = types.Int64Null()
	}
	if value := res.Get(prefix + "interval-singlehop-v2.mill-unit.both"); value.Exists() && !data.IntervalSinglehopV2MillUnitBoth.IsNull() {
		data.IntervalSinglehopV2MillUnitBoth = types.Int64Value(value.Int())
	} else {
		data.IntervalSinglehopV2MillUnitBoth = types.Int64Null()
	}
	if value := res.Get(prefix + "interval-singlehop-v2.mill-unit.multiplier"); value.Exists() && !data.IntervalSinglehopV2MillUnitMultiplier.IsNull() {
		data.IntervalSinglehopV2MillUnitMultiplier = types.Int64Value(value.Int())
	} else {
		data.IntervalSinglehopV2MillUnitMultiplier = types.Int64Null()
	}
	if value := res.Get(prefix + "interval-singlehop-v2.ms-unit.min-rx"); value.Exists() && !data.IntervalSinglehopV2MsUnitMinRx.IsNull() {
		data.IntervalSinglehopV2MsUnitMinRx = types.Int64Value(value.Int())
	} else {
		data.IntervalSinglehopV2MsUnitMinRx = types.Int64Null()
	}
	if value := res.Get(prefix + "interval-singlehop-v2.ms-unit.min-tx"); value.Exists() && !data.IntervalSinglehopV2MsUnitMinTx.IsNull() {
		data.IntervalSinglehopV2MsUnitMinTx = types.Int64Value(value.Int())
	} else {
		data.IntervalSinglehopV2MsUnitMinTx = types.Int64Null()
	}
	if value := res.Get(prefix + "echo"); !data.Echo.IsNull() {
		if value.Exists() {
			data.Echo = types.BoolValue(true)
		} else {
			data.Echo = types.BoolValue(false)
		}
	} else {
		data.Echo = types.BoolNull()
	}
	if value := res.Get(prefix + "dampening.half-time"); value.Exists() && !data.DampeningHalfTime.IsNull() {
		data.DampeningHalfTime = types.Int64Value(value.Int())
	} else {
		data.DampeningHalfTime = types.Int64Null()
	}
	if value := res.Get(prefix + "dampening.unsuppress-time"); value.Exists() && !data.DampeningUnsuppressTime.IsNull() {
		data.DampeningUnsuppressTime = types.Int64Value(value.Int())
	} else {
		data.DampeningUnsuppressTime = types.Int64Null()
	}
	if value := res.Get(prefix + "dampening.suppress-time"); value.Exists() && !data.DampeningSuppressTime.IsNull() {
		data.DampeningSuppressTime = types.Int64Value(value.Int())
	} else {
		data.DampeningSuppressTime = types.Int64Null()
	}
	if value := res.Get(prefix + "dampening.max-suppressing-time"); value.Exists() && !data.DampeningMaxSuppressingTime.IsNull() {
		data.DampeningMaxSuppressingTime = types.Int64Value(value.Int())
	} else {
		data.DampeningMaxSuppressingTime = types.Int64Null()
	}
}

func (data *BFDTemplateSingleHopData) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "authentication.md5.keychain"); value.Exists() {
		data.AuthenticationMd5Keychain = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "authentication.meticulous-md5.keychain"); value.Exists() {
		data.AuthenticationMeticulousMd5Keychain = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "authentication.meticulous-sha-1.keychain"); value.Exists() {
		data.AuthenticationMeticulousSha1Keychain = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "authentication.sha-1.keychain"); value.Exists() {
		data.AuthenticationSha1Keychain = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "interval-singlehop-v2.mill-unit.min-tx"); value.Exists() {
		data.IntervalSinglehopV2MillUnitMinTx = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "interval-singlehop-v2.mill-unit.min-rx"); value.Exists() {
		data.IntervalSinglehopV2MillUnitMinRx = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "interval-singlehop-v2.mill-unit.both"); value.Exists() {
		data.IntervalSinglehopV2MillUnitBoth = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "interval-singlehop-v2.mill-unit.multiplier"); value.Exists() {
		data.IntervalSinglehopV2MillUnitMultiplier = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "interval-singlehop-v2.ms-unit.min-rx"); value.Exists() {
		data.IntervalSinglehopV2MsUnitMinRx = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "interval-singlehop-v2.ms-unit.min-tx"); value.Exists() {
		data.IntervalSinglehopV2MsUnitMinTx = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "echo"); value.Exists() {
		data.Echo = types.BoolValue(true)
	} else {
		data.Echo = types.BoolValue(false)
	}
	if value := res.Get(prefix + "dampening.half-time"); value.Exists() {
		data.DampeningHalfTime = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "dampening.unsuppress-time"); value.Exists() {
		data.DampeningUnsuppressTime = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "dampening.suppress-time"); value.Exists() {
		data.DampeningSuppressTime = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "dampening.max-suppressing-time"); value.Exists() {
		data.DampeningMaxSuppressingTime = types.Int64Value(value.Int())
	}
}

func (data *BFDTemplateSingleHop) getDeletedListItems(ctx context.Context, state BFDTemplateSingleHop) []string {
	deletedListItems := make([]string, 0)
	return deletedListItems
}

func (data *BFDTemplateSingleHop) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	if !data.Echo.IsNull() && !data.Echo.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/echo", data.getPath()))
	}
	return emptyLeafsDelete
}

func (data *BFDTemplateSingleHop) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.AuthenticationMd5Keychain.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/authentication/md5/keychain", data.getPath()))
	}
	if !data.AuthenticationMeticulousMd5Keychain.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/authentication/meticulous-md5/keychain", data.getPath()))
	}
	if !data.AuthenticationMeticulousSha1Keychain.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/authentication/meticulous-sha-1/keychain", data.getPath()))
	}
	if !data.AuthenticationSha1Keychain.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/authentication/sha-1/keychain", data.getPath()))
	}
	if !data.IntervalSinglehopV2MillUnitMinTx.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/interval-singlehop-v2/mill-unit/min-tx", data.getPath()))
	}
	if !data.IntervalSinglehopV2MillUnitMinRx.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/interval-singlehop-v2/mill-unit/min-rx", data.getPath()))
	}
	if !data.IntervalSinglehopV2MillUnitBoth.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/interval-singlehop-v2/mill-unit/both", data.getPath()))
	}
	if !data.IntervalSinglehopV2MillUnitMultiplier.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/interval-singlehop-v2/mill-unit/multiplier", data.getPath()))
	}
	if !data.IntervalSinglehopV2MsUnitMinRx.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/interval-singlehop-v2/ms-unit/min-rx", data.getPath()))
	}
	if !data.IntervalSinglehopV2MsUnitMinTx.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/interval-singlehop-v2/ms-unit/min-tx", data.getPath()))
	}
	if !data.Echo.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/echo", data.getPath()))
	}
	if !data.DampeningHalfTime.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/dampening/half-time", data.getPath()))
	}
	if !data.DampeningUnsuppressTime.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/dampening/unsuppress-time", data.getPath()))
	}
	if !data.DampeningSuppressTime.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/dampening/suppress-time", data.getPath()))
	}
	if !data.DampeningMaxSuppressingTime.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/dampening/max-suppressing-time", data.getPath()))
	}
	return deletePaths
}