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

type Radius struct {
	Device                       types.String `tfsdk:"device"`
	Id                           types.String `tfsdk:"id"`
	Name                         types.String `tfsdk:"name"`
	Ipv4Address                  types.String `tfsdk:"ipv4_address"`
	AuthenticationPort           types.Int64  `tfsdk:"authentication_port"`
	AccountingPort               types.Int64  `tfsdk:"accounting_port"`
	Timeout                      types.Int64  `tfsdk:"timeout"`
	Retransmit                   types.Int64  `tfsdk:"retransmit"`
	Key                          types.String `tfsdk:"key"`
	AutomateTesterUsername       types.String `tfsdk:"automate_tester_username"`
	AutomateTesterIgnoreAcctPort types.Bool   `tfsdk:"automate_tester_ignore_acct_port"`
	AutomateTesterProbeOnConfig  types.Bool   `tfsdk:"automate_tester_probe_on_config"`
	PacKey                       types.String `tfsdk:"pac_key"`
	PacKeyEncryption             types.String `tfsdk:"pac_key_encryption"`
}

type RadiusData struct {
	Device                       types.String `tfsdk:"device"`
	Id                           types.String `tfsdk:"id"`
	Name                         types.String `tfsdk:"name"`
	Ipv4Address                  types.String `tfsdk:"ipv4_address"`
	AuthenticationPort           types.Int64  `tfsdk:"authentication_port"`
	AccountingPort               types.Int64  `tfsdk:"accounting_port"`
	Timeout                      types.Int64  `tfsdk:"timeout"`
	Retransmit                   types.Int64  `tfsdk:"retransmit"`
	Key                          types.String `tfsdk:"key"`
	AutomateTesterUsername       types.String `tfsdk:"automate_tester_username"`
	AutomateTesterIgnoreAcctPort types.Bool   `tfsdk:"automate_tester_ignore_acct_port"`
	AutomateTesterProbeOnConfig  types.Bool   `tfsdk:"automate_tester_probe_on_config"`
	PacKey                       types.String `tfsdk:"pac_key"`
	PacKeyEncryption             types.String `tfsdk:"pac_key_encryption"`
}

func (data Radius) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/radius/Cisco-IOS-XE-aaa:server=%v", url.QueryEscape(fmt.Sprintf("%v", data.Name.ValueString())))
}

func (data RadiusData) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/radius/Cisco-IOS-XE-aaa:server=%v", url.QueryEscape(fmt.Sprintf("%v", data.Name.ValueString())))
}

// if last path element has a key -> remove it
func (data Radius) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data Radius) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"id", data.Name.ValueString())
	}
	if !data.Ipv4Address.IsNull() && !data.Ipv4Address.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"address.ipv4", data.Ipv4Address.ValueString())
	}
	if !data.AuthenticationPort.IsNull() && !data.AuthenticationPort.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"address.auth-port", strconv.FormatInt(data.AuthenticationPort.ValueInt64(), 10))
	}
	if !data.AccountingPort.IsNull() && !data.AccountingPort.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"address.acct-port", strconv.FormatInt(data.AccountingPort.ValueInt64(), 10))
	}
	if !data.Timeout.IsNull() && !data.Timeout.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"timeout", strconv.FormatInt(data.Timeout.ValueInt64(), 10))
	}
	if !data.Retransmit.IsNull() && !data.Retransmit.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"retransmit", strconv.FormatInt(data.Retransmit.ValueInt64(), 10))
	}
	if !data.Key.IsNull() && !data.Key.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"key.key", data.Key.ValueString())
	}
	if !data.AutomateTesterUsername.IsNull() && !data.AutomateTesterUsername.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"automate-tester.username", data.AutomateTesterUsername.ValueString())
	}
	if !data.AutomateTesterIgnoreAcctPort.IsNull() && !data.AutomateTesterIgnoreAcctPort.IsUnknown() {
		if data.AutomateTesterIgnoreAcctPort.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"automate-tester.ignore-acct-port", map[string]string{})
		}
	}
	if !data.AutomateTesterProbeOnConfig.IsNull() && !data.AutomateTesterProbeOnConfig.IsUnknown() {
		if data.AutomateTesterProbeOnConfig.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"automate-tester.probe-on-config", map[string]string{})
		}
	}
	if !data.PacKey.IsNull() && !data.PacKey.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"pac.key.key", data.PacKey.ValueString())
	}
	if !data.PacKeyEncryption.IsNull() && !data.PacKeyEncryption.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"pac.key.encryption", data.PacKeyEncryption.ValueString())
	}
	return body
}

func (data *Radius) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "id"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get(prefix + "address.ipv4"); value.Exists() && !data.Ipv4Address.IsNull() {
		data.Ipv4Address = types.StringValue(value.String())
	} else {
		data.Ipv4Address = types.StringNull()
	}
	if value := res.Get(prefix + "address.auth-port"); value.Exists() && !data.AuthenticationPort.IsNull() {
		data.AuthenticationPort = types.Int64Value(value.Int())
	} else {
		data.AuthenticationPort = types.Int64Null()
	}
	if value := res.Get(prefix + "address.acct-port"); value.Exists() && !data.AccountingPort.IsNull() {
		data.AccountingPort = types.Int64Value(value.Int())
	} else {
		data.AccountingPort = types.Int64Null()
	}
	if value := res.Get(prefix + "timeout"); value.Exists() && !data.Timeout.IsNull() {
		data.Timeout = types.Int64Value(value.Int())
	} else {
		data.Timeout = types.Int64Null()
	}
	if value := res.Get(prefix + "retransmit"); value.Exists() && !data.Retransmit.IsNull() {
		data.Retransmit = types.Int64Value(value.Int())
	} else {
		data.Retransmit = types.Int64Null()
	}
	if value := res.Get(prefix + "key.key"); value.Exists() && !data.Key.IsNull() {
		data.Key = types.StringValue(value.String())
	} else {
		data.Key = types.StringNull()
	}
	if value := res.Get(prefix + "automate-tester.username"); value.Exists() && !data.AutomateTesterUsername.IsNull() {
		data.AutomateTesterUsername = types.StringValue(value.String())
	} else {
		data.AutomateTesterUsername = types.StringNull()
	}
	if value := res.Get(prefix + "automate-tester.ignore-acct-port"); !data.AutomateTesterIgnoreAcctPort.IsNull() {
		if value.Exists() {
			data.AutomateTesterIgnoreAcctPort = types.BoolValue(true)
		} else {
			data.AutomateTesterIgnoreAcctPort = types.BoolValue(false)
		}
	} else {
		data.AutomateTesterIgnoreAcctPort = types.BoolNull()
	}
	if value := res.Get(prefix + "automate-tester.probe-on-config"); !data.AutomateTesterProbeOnConfig.IsNull() {
		if value.Exists() {
			data.AutomateTesterProbeOnConfig = types.BoolValue(true)
		} else {
			data.AutomateTesterProbeOnConfig = types.BoolValue(false)
		}
	} else {
		data.AutomateTesterProbeOnConfig = types.BoolNull()
	}
	if value := res.Get(prefix + "pac.key.key"); value.Exists() && !data.PacKey.IsNull() {
		data.PacKey = types.StringValue(value.String())
	} else {
		data.PacKey = types.StringNull()
	}
	if value := res.Get(prefix + "pac.key.encryption"); value.Exists() && !data.PacKeyEncryption.IsNull() {
		data.PacKeyEncryption = types.StringValue(value.String())
	} else {
		data.PacKeyEncryption = types.StringNull()
	}
}

func (data *RadiusData) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "address.ipv4"); value.Exists() {
		data.Ipv4Address = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "address.auth-port"); value.Exists() {
		data.AuthenticationPort = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "address.acct-port"); value.Exists() {
		data.AccountingPort = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "timeout"); value.Exists() {
		data.Timeout = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "retransmit"); value.Exists() {
		data.Retransmit = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "key.key"); value.Exists() {
		data.Key = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "automate-tester.username"); value.Exists() {
		data.AutomateTesterUsername = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "automate-tester.ignore-acct-port"); value.Exists() {
		data.AutomateTesterIgnoreAcctPort = types.BoolValue(true)
	} else {
		data.AutomateTesterIgnoreAcctPort = types.BoolValue(false)
	}
	if value := res.Get(prefix + "automate-tester.probe-on-config"); value.Exists() {
		data.AutomateTesterProbeOnConfig = types.BoolValue(true)
	} else {
		data.AutomateTesterProbeOnConfig = types.BoolValue(false)
	}
	if value := res.Get(prefix + "pac.key.key"); value.Exists() {
		data.PacKey = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "pac.key.encryption"); value.Exists() {
		data.PacKeyEncryption = types.StringValue(value.String())
	}
}

func (data *Radius) getDeletedItems(ctx context.Context, state Radius) []string {
	deletedItems := make([]string, 0)
	if !state.Ipv4Address.IsNull() && data.Ipv4Address.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/address/ipv4", state.getPath()))
	}
	if !state.AuthenticationPort.IsNull() && data.AuthenticationPort.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/address/auth-port", state.getPath()))
	}
	if !state.AccountingPort.IsNull() && data.AccountingPort.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/address/acct-port", state.getPath()))
	}
	if !state.Timeout.IsNull() && data.Timeout.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/timeout", state.getPath()))
	}
	if !state.Retransmit.IsNull() && data.Retransmit.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/retransmit", state.getPath()))
	}
	if !state.Key.IsNull() && data.Key.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/key/key", state.getPath()))
	}
	if !state.AutomateTesterUsername.IsNull() && data.AutomateTesterUsername.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/automate-tester/username", state.getPath()))
	}
	if !state.AutomateTesterIgnoreAcctPort.IsNull() && data.AutomateTesterIgnoreAcctPort.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/automate-tester/ignore-acct-port", state.getPath()))
	}
	if !state.AutomateTesterProbeOnConfig.IsNull() && data.AutomateTesterProbeOnConfig.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/automate-tester/probe-on-config", state.getPath()))
	}
	if !state.PacKey.IsNull() && data.PacKey.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/pac/key/key", state.getPath()))
	}
	if !state.PacKeyEncryption.IsNull() && data.PacKeyEncryption.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/pac/key/encryption", state.getPath()))
	}
	return deletedItems
}

func (data *Radius) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	if !data.AutomateTesterIgnoreAcctPort.IsNull() && !data.AutomateTesterIgnoreAcctPort.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/automate-tester/ignore-acct-port", data.getPath()))
	}
	if !data.AutomateTesterProbeOnConfig.IsNull() && !data.AutomateTesterProbeOnConfig.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/automate-tester/probe-on-config", data.getPath()))
	}
	return emptyLeafsDelete
}

func (data *Radius) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.Ipv4Address.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/address/ipv4", data.getPath()))
	}
	if !data.AuthenticationPort.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/address/auth-port", data.getPath()))
	}
	if !data.AccountingPort.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/address/acct-port", data.getPath()))
	}
	if !data.Timeout.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/timeout", data.getPath()))
	}
	if !data.Retransmit.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/retransmit", data.getPath()))
	}
	if !data.Key.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/key/key", data.getPath()))
	}
	if !data.AutomateTesterUsername.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/automate-tester/username", data.getPath()))
	}
	if !data.AutomateTesterIgnoreAcctPort.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/automate-tester/ignore-acct-port", data.getPath()))
	}
	if !data.AutomateTesterProbeOnConfig.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/automate-tester/probe-on-config", data.getPath()))
	}
	if !data.PacKey.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/pac/key/key", data.getPath()))
	}
	if !data.PacKeyEncryption.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/pac/key/encryption", data.getPath()))
	}
	return deletePaths
}
