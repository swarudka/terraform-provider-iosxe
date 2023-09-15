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

	"github.com/CiscoDevNet/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type CryptoIPSecTransformSet struct {
	Device     types.String `tfsdk:"device"`
	Id         types.String `tfsdk:"id"`
	Name       types.String `tfsdk:"name"`
	Esp        types.String `tfsdk:"esp"`
	EspHmac    types.String `tfsdk:"esp_hmac"`
	ModeTunnel types.Bool   `tfsdk:"mode_tunnel"`
}

type CryptoIPSecTransformSetData struct {
	Device     types.String `tfsdk:"device"`
	Id         types.String `tfsdk:"id"`
	Name       types.String `tfsdk:"name"`
	Esp        types.String `tfsdk:"esp"`
	EspHmac    types.String `tfsdk:"esp_hmac"`
	ModeTunnel types.Bool   `tfsdk:"mode_tunnel"`
}

func (data CryptoIPSecTransformSet) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/crypto/Cisco-IOS-XE-crypto:ipsec/transform-set=%s", url.QueryEscape(fmt.Sprintf("%v", data.Name.ValueString())))
}

func (data CryptoIPSecTransformSetData) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/crypto/Cisco-IOS-XE-crypto:ipsec/transform-set=%s", url.QueryEscape(fmt.Sprintf("%v", data.Name.ValueString())))
}

// if last path element has a key -> remove it
func (data CryptoIPSecTransformSet) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data CryptoIPSecTransformSet) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"tag", data.Name.ValueString())
	}
	if !data.Esp.IsNull() && !data.Esp.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"esp", data.Esp.ValueString())
	}
	if !data.EspHmac.IsNull() && !data.EspHmac.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"esp-hmac", data.EspHmac.ValueString())
	}
	if !data.ModeTunnel.IsNull() && !data.ModeTunnel.IsUnknown() {
		if data.ModeTunnel.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"mode.tunnel-choice", map[string]string{})
		}
	}
	return body
}

func (data *CryptoIPSecTransformSet) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "tag"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get(prefix + "esp"); value.Exists() && !data.Esp.IsNull() {
		data.Esp = types.StringValue(value.String())
	} else {
		data.Esp = types.StringNull()
	}
	if value := res.Get(prefix + "esp-hmac"); value.Exists() && !data.EspHmac.IsNull() {
		data.EspHmac = types.StringValue(value.String())
	} else {
		data.EspHmac = types.StringNull()
	}
	if value := res.Get(prefix + "mode.tunnel-choice"); !data.ModeTunnel.IsNull() {
		if value.Exists() {
			data.ModeTunnel = types.BoolValue(true)
		} else {
			data.ModeTunnel = types.BoolValue(false)
		}
	} else {
		data.ModeTunnel = types.BoolNull()
	}
}

func (data *CryptoIPSecTransformSetData) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "esp"); value.Exists() {
		data.Esp = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "esp-hmac"); value.Exists() {
		data.EspHmac = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "mode.tunnel-choice"); value.Exists() {
		data.ModeTunnel = types.BoolValue(true)
	} else {
		data.ModeTunnel = types.BoolValue(false)
	}
}

func (data *CryptoIPSecTransformSet) getDeletedItems(ctx context.Context, state CryptoIPSecTransformSet) []string {
	deletedItems := make([]string, 0)
	if !state.Esp.IsNull() && data.Esp.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/esp", state.getPath()))
	}
	if !state.EspHmac.IsNull() && data.EspHmac.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/esp-hmac", state.getPath()))
	}
	if !state.ModeTunnel.IsNull() && data.ModeTunnel.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/mode/tunnel-choice", state.getPath()))
	}
	return deletedItems
}

func (data *CryptoIPSecTransformSet) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	if !data.ModeTunnel.IsNull() && !data.ModeTunnel.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/mode/tunnel-choice", data.getPath()))
	}
	return emptyLeafsDelete
}

func (data *CryptoIPSecTransformSet) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.Esp.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/esp", data.getPath()))
	}
	if !data.EspHmac.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/esp-hmac", data.getPath()))
	}
	if !data.ModeTunnel.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/mode/tunnel-choice", data.getPath()))
	}
	return deletePaths
}
