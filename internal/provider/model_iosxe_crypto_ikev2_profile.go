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
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/CiscoDevNet/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type CryptoIKEv2Profile struct {
	Device                           types.String                                         `tfsdk:"device"`
	Id                               types.String                                         `tfsdk:"id"`
	DeleteMode                       types.String                                         `tfsdk:"delete_mode"`
	Name                             types.String                                         `tfsdk:"name"`
	Description                      types.String                                         `tfsdk:"description"`
	AuthenticationRemotePreShare     types.Bool                                           `tfsdk:"authentication_remote_pre_share"`
	AuthenticationLocalPreShare      types.Bool                                           `tfsdk:"authentication_local_pre_share"`
	IdentityLocalAddress             types.String                                         `tfsdk:"identity_local_address"`
	IdentityLocalKeyId               types.String                                         `tfsdk:"identity_local_key_id"`
	MatchInboundOnly                 types.Bool                                           `tfsdk:"match_inbound_only"`
	MatchAddressLocalIp              types.String                                         `tfsdk:"match_address_local_ip"`
	MatchFvrf                        types.String                                         `tfsdk:"match_fvrf"`
	MatchFvrfAny                     types.Bool                                           `tfsdk:"match_fvrf_any"`
	MatchIdentityRemoteIpv4Addresses []CryptoIKEv2ProfileMatchIdentityRemoteIpv4Addresses `tfsdk:"match_identity_remote_ipv4_addresses"`
	MatchIdentityRemoteIpv6Prefixes  types.List                                           `tfsdk:"match_identity_remote_ipv6_prefixes"`
	MatchIdentityRemoteKeys          types.List                                           `tfsdk:"match_identity_remote_keys"`
	KeyringLocal                     types.String                                         `tfsdk:"keyring_local"`
	DpdInterval                      types.Int64                                          `tfsdk:"dpd_interval"`
	DpdRetry                         types.Int64                                          `tfsdk:"dpd_retry"`
	DpdQuery                         types.String                                         `tfsdk:"dpd_query"`
	ConfigExchangeRequest            types.Bool                                           `tfsdk:"config_exchange_request"`
}

type CryptoIKEv2ProfileData struct {
	Device                           types.String                                         `tfsdk:"device"`
	Id                               types.String                                         `tfsdk:"id"`
	Name                             types.String                                         `tfsdk:"name"`
	Description                      types.String                                         `tfsdk:"description"`
	AuthenticationRemotePreShare     types.Bool                                           `tfsdk:"authentication_remote_pre_share"`
	AuthenticationLocalPreShare      types.Bool                                           `tfsdk:"authentication_local_pre_share"`
	IdentityLocalAddress             types.String                                         `tfsdk:"identity_local_address"`
	IdentityLocalKeyId               types.String                                         `tfsdk:"identity_local_key_id"`
	MatchInboundOnly                 types.Bool                                           `tfsdk:"match_inbound_only"`
	MatchAddressLocalIp              types.String                                         `tfsdk:"match_address_local_ip"`
	MatchFvrf                        types.String                                         `tfsdk:"match_fvrf"`
	MatchFvrfAny                     types.Bool                                           `tfsdk:"match_fvrf_any"`
	MatchIdentityRemoteIpv4Addresses []CryptoIKEv2ProfileMatchIdentityRemoteIpv4Addresses `tfsdk:"match_identity_remote_ipv4_addresses"`
	MatchIdentityRemoteIpv6Prefixes  types.List                                           `tfsdk:"match_identity_remote_ipv6_prefixes"`
	MatchIdentityRemoteKeys          types.List                                           `tfsdk:"match_identity_remote_keys"`
	KeyringLocal                     types.String                                         `tfsdk:"keyring_local"`
	DpdInterval                      types.Int64                                          `tfsdk:"dpd_interval"`
	DpdRetry                         types.Int64                                          `tfsdk:"dpd_retry"`
	DpdQuery                         types.String                                         `tfsdk:"dpd_query"`
	ConfigExchangeRequest            types.Bool                                           `tfsdk:"config_exchange_request"`
}
type CryptoIKEv2ProfileMatchIdentityRemoteIpv4Addresses struct {
	Address types.String `tfsdk:"address"`
	Mask    types.String `tfsdk:"mask"`
}

func (data CryptoIKEv2Profile) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/crypto/Cisco-IOS-XE-crypto:ikev2/profile=%v", url.QueryEscape(fmt.Sprintf("%v", data.Name.ValueString())))
}

func (data CryptoIKEv2ProfileData) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/crypto/Cisco-IOS-XE-crypto:ikev2/profile=%v", url.QueryEscape(fmt.Sprintf("%v", data.Name.ValueString())))
}

// if last path element has a key -> remove it
func (data CryptoIKEv2Profile) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data CryptoIKEv2Profile) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"name", data.Name.ValueString())
	}
	if !data.Description.IsNull() && !data.Description.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"description", data.Description.ValueString())
	}
	if !data.AuthenticationRemotePreShare.IsNull() && !data.AuthenticationRemotePreShare.IsUnknown() {
		if data.AuthenticationRemotePreShare.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"authentication.remote.pre-share", map[string]string{})
		}
	}
	if !data.AuthenticationLocalPreShare.IsNull() && !data.AuthenticationLocalPreShare.IsUnknown() {
		if data.AuthenticationLocalPreShare.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"authentication.local.pre-share", map[string]string{})
		}
	}
	if !data.IdentityLocalAddress.IsNull() && !data.IdentityLocalAddress.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"identity.local.address", data.IdentityLocalAddress.ValueString())
	}
	if !data.IdentityLocalKeyId.IsNull() && !data.IdentityLocalKeyId.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"identity.local.key-id", data.IdentityLocalKeyId.ValueString())
	}
	if !data.MatchInboundOnly.IsNull() && !data.MatchInboundOnly.IsUnknown() {
		if data.MatchInboundOnly.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"match.inbound-only", map[string]string{})
		}
	}
	if !data.MatchAddressLocalIp.IsNull() && !data.MatchAddressLocalIp.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"match.address.local.ip", data.MatchAddressLocalIp.ValueString())
	}
	if !data.MatchFvrf.IsNull() && !data.MatchFvrf.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"match.fvrf.name", data.MatchFvrf.ValueString())
	}
	if !data.MatchFvrfAny.IsNull() && !data.MatchFvrfAny.IsUnknown() {
		if data.MatchFvrfAny.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"match.fvrf.any", map[string]string{})
		}
	}
	if !data.MatchIdentityRemoteIpv6Prefixes.IsNull() && !data.MatchIdentityRemoteIpv6Prefixes.IsUnknown() {
		var values []string
		data.MatchIdentityRemoteIpv6Prefixes.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"match.identity.remote.address.ipv6-prefix", values)
	}
	if !data.MatchIdentityRemoteKeys.IsNull() && !data.MatchIdentityRemoteKeys.IsUnknown() {
		var values []string
		data.MatchIdentityRemoteKeys.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"match.identity.remote.key-ids", values)
	}
	if !data.KeyringLocal.IsNull() && !data.KeyringLocal.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"keyring.local.name", data.KeyringLocal.ValueString())
	}
	if !data.DpdInterval.IsNull() && !data.DpdInterval.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"dpd.interval", strconv.FormatInt(data.DpdInterval.ValueInt64(), 10))
	}
	if !data.DpdRetry.IsNull() && !data.DpdRetry.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"dpd.retry", strconv.FormatInt(data.DpdRetry.ValueInt64(), 10))
	}
	if !data.DpdQuery.IsNull() && !data.DpdQuery.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"dpd.query", data.DpdQuery.ValueString())
	}
	if !data.ConfigExchangeRequest.IsNull() && !data.ConfigExchangeRequest.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"config-exchange.request-1", data.ConfigExchangeRequest.ValueBool())
	}
	if len(data.MatchIdentityRemoteIpv4Addresses) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"match.identity.remote.address.ipv4", []interface{}{})
		for index, item := range data.MatchIdentityRemoteIpv4Addresses {
			if !item.Address.IsNull() && !item.Address.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"match.identity.remote.address.ipv4"+"."+strconv.Itoa(index)+"."+"ipv4-address", item.Address.ValueString())
			}
			if !item.Mask.IsNull() && !item.Mask.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"match.identity.remote.address.ipv4"+"."+strconv.Itoa(index)+"."+"ipv4-mask", item.Mask.ValueString())
			}
		}
	}
	return body
}

func (data *CryptoIKEv2Profile) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get(prefix + "description"); value.Exists() && !data.Description.IsNull() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get(prefix + "authentication.remote.pre-share"); !data.AuthenticationRemotePreShare.IsNull() {
		if value.Exists() {
			data.AuthenticationRemotePreShare = types.BoolValue(true)
		} else {
			data.AuthenticationRemotePreShare = types.BoolValue(false)
		}
	} else {
		data.AuthenticationRemotePreShare = types.BoolNull()
	}
	if value := res.Get(prefix + "authentication.local.pre-share"); !data.AuthenticationLocalPreShare.IsNull() {
		if value.Exists() {
			data.AuthenticationLocalPreShare = types.BoolValue(true)
		} else {
			data.AuthenticationLocalPreShare = types.BoolValue(false)
		}
	} else {
		data.AuthenticationLocalPreShare = types.BoolNull()
	}
	if value := res.Get(prefix + "identity.local.address"); value.Exists() && !data.IdentityLocalAddress.IsNull() {
		data.IdentityLocalAddress = types.StringValue(value.String())
	} else {
		data.IdentityLocalAddress = types.StringNull()
	}
	if value := res.Get(prefix + "identity.local.key-id"); value.Exists() && !data.IdentityLocalKeyId.IsNull() {
		data.IdentityLocalKeyId = types.StringValue(value.String())
	} else {
		data.IdentityLocalKeyId = types.StringNull()
	}
	if value := res.Get(prefix + "match.inbound-only"); !data.MatchInboundOnly.IsNull() {
		if value.Exists() {
			data.MatchInboundOnly = types.BoolValue(true)
		} else {
			data.MatchInboundOnly = types.BoolValue(false)
		}
	} else {
		data.MatchInboundOnly = types.BoolNull()
	}
	if value := res.Get(prefix + "match.address.local.ip"); value.Exists() && !data.MatchAddressLocalIp.IsNull() {
		data.MatchAddressLocalIp = types.StringValue(value.String())
	} else {
		data.MatchAddressLocalIp = types.StringNull()
	}
	if value := res.Get(prefix + "match.fvrf.name"); value.Exists() && !data.MatchFvrf.IsNull() {
		data.MatchFvrf = types.StringValue(value.String())
	} else {
		data.MatchFvrf = types.StringNull()
	}
	if value := res.Get(prefix + "match.fvrf.any"); !data.MatchFvrfAny.IsNull() {
		if value.Exists() {
			data.MatchFvrfAny = types.BoolValue(true)
		} else {
			data.MatchFvrfAny = types.BoolValue(false)
		}
	} else {
		data.MatchFvrfAny = types.BoolNull()
	}
	for i := range data.MatchIdentityRemoteIpv4Addresses {
		keys := [...]string{"ipv4-address"}
		keyValues := [...]string{data.MatchIdentityRemoteIpv4Addresses[i].Address.ValueString()}

		var r gjson.Result
		res.Get(prefix + "match.identity.remote.address.ipv4").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() == keyValues[ik] {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if value := r.Get("ipv4-address"); value.Exists() && !data.MatchIdentityRemoteIpv4Addresses[i].Address.IsNull() {
			data.MatchIdentityRemoteIpv4Addresses[i].Address = types.StringValue(value.String())
		} else {
			data.MatchIdentityRemoteIpv4Addresses[i].Address = types.StringNull()
		}
		if value := r.Get("ipv4-mask"); value.Exists() && !data.MatchIdentityRemoteIpv4Addresses[i].Mask.IsNull() {
			data.MatchIdentityRemoteIpv4Addresses[i].Mask = types.StringValue(value.String())
		} else {
			data.MatchIdentityRemoteIpv4Addresses[i].Mask = types.StringNull()
		}
	}
	if value := res.Get(prefix + "match.identity.remote.address.ipv6-prefix"); value.Exists() && !data.MatchIdentityRemoteIpv6Prefixes.IsNull() {
		data.MatchIdentityRemoteIpv6Prefixes = helpers.GetStringList(value.Array())
	} else {
		data.MatchIdentityRemoteIpv6Prefixes = types.ListNull(types.StringType)
	}
	if value := res.Get(prefix + "match.identity.remote.key-ids"); value.Exists() && !data.MatchIdentityRemoteKeys.IsNull() {
		data.MatchIdentityRemoteKeys = helpers.GetStringList(value.Array())
	} else {
		data.MatchIdentityRemoteKeys = types.ListNull(types.StringType)
	}
	if value := res.Get(prefix + "keyring.local.name"); value.Exists() && !data.KeyringLocal.IsNull() {
		data.KeyringLocal = types.StringValue(value.String())
	} else {
		data.KeyringLocal = types.StringNull()
	}
	if value := res.Get(prefix + "dpd.interval"); value.Exists() && !data.DpdInterval.IsNull() {
		data.DpdInterval = types.Int64Value(value.Int())
	} else {
		data.DpdInterval = types.Int64Null()
	}
	if value := res.Get(prefix + "dpd.retry"); value.Exists() && !data.DpdRetry.IsNull() {
		data.DpdRetry = types.Int64Value(value.Int())
	} else {
		data.DpdRetry = types.Int64Null()
	}
	if value := res.Get(prefix + "dpd.query"); value.Exists() && !data.DpdQuery.IsNull() {
		data.DpdQuery = types.StringValue(value.String())
	} else {
		data.DpdQuery = types.StringNull()
	}
	if value := res.Get(prefix + "config-exchange.request-1"); !data.ConfigExchangeRequest.IsNull() {
		if value.Exists() {
			data.ConfigExchangeRequest = types.BoolValue(value.Bool())
		}
	} else {
		data.ConfigExchangeRequest = types.BoolNull()
	}
}

func (data *CryptoIKEv2ProfileData) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "authentication.remote.pre-share"); value.Exists() {
		data.AuthenticationRemotePreShare = types.BoolValue(true)
	} else {
		data.AuthenticationRemotePreShare = types.BoolValue(false)
	}
	if value := res.Get(prefix + "authentication.local.pre-share"); value.Exists() {
		data.AuthenticationLocalPreShare = types.BoolValue(true)
	} else {
		data.AuthenticationLocalPreShare = types.BoolValue(false)
	}
	if value := res.Get(prefix + "identity.local.address"); value.Exists() {
		data.IdentityLocalAddress = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "identity.local.key-id"); value.Exists() {
		data.IdentityLocalKeyId = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "match.inbound-only"); value.Exists() {
		data.MatchInboundOnly = types.BoolValue(true)
	} else {
		data.MatchInboundOnly = types.BoolValue(false)
	}
	if value := res.Get(prefix + "match.address.local.ip"); value.Exists() {
		data.MatchAddressLocalIp = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "match.fvrf.name"); value.Exists() {
		data.MatchFvrf = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "match.fvrf.any"); value.Exists() {
		data.MatchFvrfAny = types.BoolValue(true)
	} else {
		data.MatchFvrfAny = types.BoolValue(false)
	}
	if value := res.Get(prefix + "match.identity.remote.address.ipv4"); value.Exists() {
		data.MatchIdentityRemoteIpv4Addresses = make([]CryptoIKEv2ProfileMatchIdentityRemoteIpv4Addresses, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CryptoIKEv2ProfileMatchIdentityRemoteIpv4Addresses{}
			if cValue := v.Get("ipv4-address"); cValue.Exists() {
				item.Address = types.StringValue(cValue.String())
			}
			if cValue := v.Get("ipv4-mask"); cValue.Exists() {
				item.Mask = types.StringValue(cValue.String())
			}
			data.MatchIdentityRemoteIpv4Addresses = append(data.MatchIdentityRemoteIpv4Addresses, item)
			return true
		})
	}
	if value := res.Get(prefix + "match.identity.remote.address.ipv6-prefix"); value.Exists() {
		data.MatchIdentityRemoteIpv6Prefixes = helpers.GetStringList(value.Array())
	} else {
		data.MatchIdentityRemoteIpv6Prefixes = types.ListNull(types.StringType)
	}
	if value := res.Get(prefix + "match.identity.remote.key-ids"); value.Exists() {
		data.MatchIdentityRemoteKeys = helpers.GetStringList(value.Array())
	} else {
		data.MatchIdentityRemoteKeys = types.ListNull(types.StringType)
	}
	if value := res.Get(prefix + "keyring.local.name"); value.Exists() {
		data.KeyringLocal = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "dpd.interval"); value.Exists() {
		data.DpdInterval = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "dpd.retry"); value.Exists() {
		data.DpdRetry = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "dpd.query"); value.Exists() {
		data.DpdQuery = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "config-exchange.request-1"); value.Exists() {
		data.ConfigExchangeRequest = types.BoolValue(value.Bool())
	} else {
		data.ConfigExchangeRequest = types.BoolValue(false)
	}
}

func (data *CryptoIKEv2Profile) getDeletedItems(ctx context.Context, state CryptoIKEv2Profile) []string {
	deletedItems := make([]string, 0)
	if !state.Description.IsNull() && data.Description.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/description", state.getPath()))
	}
	if !state.AuthenticationRemotePreShare.IsNull() && data.AuthenticationRemotePreShare.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/authentication/remote/pre-share", state.getPath()))
	}
	if !state.AuthenticationLocalPreShare.IsNull() && data.AuthenticationLocalPreShare.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/authentication/local/pre-share", state.getPath()))
	}
	if !state.IdentityLocalAddress.IsNull() && data.IdentityLocalAddress.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/identity/local/address", state.getPath()))
	}
	if !state.IdentityLocalKeyId.IsNull() && data.IdentityLocalKeyId.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/identity/local/key-id", state.getPath()))
	}
	if !state.MatchInboundOnly.IsNull() && data.MatchInboundOnly.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/match/inbound-only", state.getPath()))
	}
	if !state.MatchAddressLocalIp.IsNull() && data.MatchAddressLocalIp.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/match/address/local/ip", state.getPath()))
	}
	if !state.MatchFvrf.IsNull() && data.MatchFvrf.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/match/fvrf/name", state.getPath()))
	}
	if !state.MatchFvrfAny.IsNull() && data.MatchFvrfAny.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/match/fvrf/any", state.getPath()))
	}
	for i := range state.MatchIdentityRemoteIpv4Addresses {
		stateKeyValues := [...]string{state.MatchIdentityRemoteIpv4Addresses[i].Address.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.MatchIdentityRemoteIpv4Addresses[i].Address.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.MatchIdentityRemoteIpv4Addresses {
			found = true
			if state.MatchIdentityRemoteIpv4Addresses[i].Address.ValueString() != data.MatchIdentityRemoteIpv4Addresses[j].Address.ValueString() {
				found = false
			}
			if found {
				if !state.MatchIdentityRemoteIpv4Addresses[i].Mask.IsNull() && data.MatchIdentityRemoteIpv4Addresses[j].Mask.IsNull() {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/match/identity/remote/address/ipv4=%v/ipv4-mask", state.getPath(), strings.Join(stateKeyValues[:], ",")))
				}
				break
			}
		}
		if !found {
			deletedItems = append(deletedItems, fmt.Sprintf("%v/match/identity/remote/address/ipv4=%v", state.getPath(), strings.Join(stateKeyValues[:], ",")))
		}
	}
	if !state.MatchIdentityRemoteIpv6Prefixes.IsNull() && data.MatchIdentityRemoteIpv6Prefixes.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/match/identity/remote/address/ipv6-prefix", state.getPath()))
	}
	if !state.MatchIdentityRemoteKeys.IsNull() && data.MatchIdentityRemoteKeys.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/match/identity/remote/key-ids", state.getPath()))
	}
	if !state.KeyringLocal.IsNull() && data.KeyringLocal.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/keyring/local/name", state.getPath()))
	}
	if !state.DpdInterval.IsNull() && data.DpdInterval.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/dpd", state.getPath()))
	}
	if !state.DpdRetry.IsNull() && data.DpdRetry.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/dpd", state.getPath()))
	}
	if !state.DpdQuery.IsNull() && data.DpdQuery.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/dpd", state.getPath()))
	}
	if !state.ConfigExchangeRequest.IsNull() && data.ConfigExchangeRequest.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/config-exchange/request-1", state.getPath()))
	}
	return deletedItems
}

func (data *CryptoIKEv2Profile) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	if !data.AuthenticationRemotePreShare.IsNull() && !data.AuthenticationRemotePreShare.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/authentication/remote/pre-share", data.getPath()))
	}
	if !data.AuthenticationLocalPreShare.IsNull() && !data.AuthenticationLocalPreShare.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/authentication/local/pre-share", data.getPath()))
	}
	if !data.MatchInboundOnly.IsNull() && !data.MatchInboundOnly.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/match/inbound-only", data.getPath()))
	}
	if !data.MatchFvrfAny.IsNull() && !data.MatchFvrfAny.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/match/fvrf/any", data.getPath()))
	}

	return emptyLeafsDelete
}

func (data *CryptoIKEv2Profile) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.Description.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/description", data.getPath()))
	}
	if !data.AuthenticationRemotePreShare.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/authentication/remote/pre-share", data.getPath()))
	}
	if !data.AuthenticationLocalPreShare.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/authentication/local/pre-share", data.getPath()))
	}
	if !data.IdentityLocalAddress.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/identity/local/address", data.getPath()))
	}
	if !data.IdentityLocalKeyId.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/identity/local/key-id", data.getPath()))
	}
	if !data.MatchInboundOnly.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/match/inbound-only", data.getPath()))
	}
	if !data.MatchAddressLocalIp.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/match/address/local/ip", data.getPath()))
	}
	if !data.MatchFvrf.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/match/fvrf/name", data.getPath()))
	}
	if !data.MatchFvrfAny.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/match/fvrf/any", data.getPath()))
	}
	for i := range data.MatchIdentityRemoteIpv4Addresses {
		keyValues := [...]string{data.MatchIdentityRemoteIpv4Addresses[i].Address.ValueString()}

		deletePaths = append(deletePaths, fmt.Sprintf("%v/match/identity/remote/address/ipv4=%v", data.getPath(), strings.Join(keyValues[:], ",")))
	}
	if !data.MatchIdentityRemoteIpv6Prefixes.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/match/identity/remote/address/ipv6-prefix", data.getPath()))
	}
	if !data.MatchIdentityRemoteKeys.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/match/identity/remote/key-ids", data.getPath()))
	}
	if !data.KeyringLocal.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/keyring/local/name", data.getPath()))
	}
	if !data.DpdInterval.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/dpd", data.getPath()))
	}
	if !data.DpdRetry.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/dpd", data.getPath()))
	}
	if !data.DpdQuery.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/dpd", data.getPath()))
	}
	if !data.ConfigExchangeRequest.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/config-exchange/request-1", data.getPath()))
	}
	return deletePaths
}
