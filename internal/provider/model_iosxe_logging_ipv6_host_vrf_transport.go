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

type LoggingIPv6HostVRFTransport struct {
	Device            types.String                                   `tfsdk:"device"`
	Id                types.String                                   `tfsdk:"id"`
	DeleteMode        types.String                                   `tfsdk:"delete_mode"`
	Ipv6Host          types.String                                   `tfsdk:"ipv6_host"`
	Vrf               types.String                                   `tfsdk:"vrf"`
	TransportUdpPorts []LoggingIPv6HostVRFTransportTransportUdpPorts `tfsdk:"transport_udp_ports"`
	TransportTcpPorts []LoggingIPv6HostVRFTransportTransportTcpPorts `tfsdk:"transport_tcp_ports"`
	TransportTlsPorts []LoggingIPv6HostVRFTransportTransportTlsPorts `tfsdk:"transport_tls_ports"`
}

type LoggingIPv6HostVRFTransportData struct {
	Device            types.String                                   `tfsdk:"device"`
	Id                types.String                                   `tfsdk:"id"`
	Ipv6Host          types.String                                   `tfsdk:"ipv6_host"`
	Vrf               types.String                                   `tfsdk:"vrf"`
	TransportUdpPorts []LoggingIPv6HostVRFTransportTransportUdpPorts `tfsdk:"transport_udp_ports"`
	TransportTcpPorts []LoggingIPv6HostVRFTransportTransportTcpPorts `tfsdk:"transport_tcp_ports"`
	TransportTlsPorts []LoggingIPv6HostVRFTransportTransportTlsPorts `tfsdk:"transport_tls_ports"`
}
type LoggingIPv6HostVRFTransportTransportUdpPorts struct {
	PortNumber types.Int64 `tfsdk:"port_number"`
}
type LoggingIPv6HostVRFTransportTransportTcpPorts struct {
	PortNumber types.Int64 `tfsdk:"port_number"`
}
type LoggingIPv6HostVRFTransportTransportTlsPorts struct {
	PortNumber types.Int64  `tfsdk:"port_number"`
	Profile    types.String `tfsdk:"profile"`
}

func (data LoggingIPv6HostVRFTransport) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/logging/host/ipv6/ipv6-host-vrf-transport-list=%s,%s", url.QueryEscape(fmt.Sprintf("%v", data.Ipv6Host.ValueString())), url.QueryEscape(fmt.Sprintf("%v", data.Vrf.ValueString())))
}

func (data LoggingIPv6HostVRFTransportData) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/logging/host/ipv6/ipv6-host-vrf-transport-list=%s,%s", url.QueryEscape(fmt.Sprintf("%v", data.Ipv6Host.ValueString())), url.QueryEscape(fmt.Sprintf("%v", data.Vrf.ValueString())))
}

// if last path element has a key -> remove it
func (data LoggingIPv6HostVRFTransport) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data LoggingIPv6HostVRFTransport) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.Ipv6Host.IsNull() && !data.Ipv6Host.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ipv6-host", data.Ipv6Host.ValueString())
	}
	if !data.Vrf.IsNull() && !data.Vrf.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vrf", data.Vrf.ValueString())
	}
	if len(data.TransportUdpPorts) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"transport.udp.port-config", []interface{}{})
		for index, item := range data.TransportUdpPorts {
			if !item.PortNumber.IsNull() && !item.PortNumber.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"transport.udp.port-config"+"."+strconv.Itoa(index)+"."+"port-number", strconv.FormatInt(item.PortNumber.ValueInt64(), 10))
			}
		}
	}
	if len(data.TransportTcpPorts) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"transport.tcp.port-config", []interface{}{})
		for index, item := range data.TransportTcpPorts {
			if !item.PortNumber.IsNull() && !item.PortNumber.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"transport.tcp.port-config"+"."+strconv.Itoa(index)+"."+"port-number", strconv.FormatInt(item.PortNumber.ValueInt64(), 10))
			}
		}
	}
	if len(data.TransportTlsPorts) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"transport.tls.port", []interface{}{})
		for index, item := range data.TransportTlsPorts {
			if !item.PortNumber.IsNull() && !item.PortNumber.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"transport.tls.port"+"."+strconv.Itoa(index)+"."+"port-number", strconv.FormatInt(item.PortNumber.ValueInt64(), 10))
			}
			if !item.Profile.IsNull() && !item.Profile.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"transport.tls.port"+"."+strconv.Itoa(index)+"."+"profile", item.Profile.ValueString())
			}
		}
	}
	return body
}

func (data *LoggingIPv6HostVRFTransport) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "ipv6-host"); value.Exists() && !data.Ipv6Host.IsNull() {
		data.Ipv6Host = types.StringValue(value.String())
	} else {
		data.Ipv6Host = types.StringNull()
	}
	if value := res.Get(prefix + "vrf"); value.Exists() && !data.Vrf.IsNull() {
		data.Vrf = types.StringValue(value.String())
	} else {
		data.Vrf = types.StringNull()
	}
	for i := range data.TransportUdpPorts {
		keys := [...]string{"port-number"}
		keyValues := [...]string{strconv.FormatInt(data.TransportUdpPorts[i].PortNumber.ValueInt64(), 10)}

		var r gjson.Result
		res.Get(prefix + "transport.udp.port-config").ForEach(
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
		if value := r.Get("port-number"); value.Exists() && !data.TransportUdpPorts[i].PortNumber.IsNull() {
			data.TransportUdpPorts[i].PortNumber = types.Int64Value(value.Int())
		} else {
			data.TransportUdpPorts[i].PortNumber = types.Int64Null()
		}
	}
	for i := range data.TransportTcpPorts {
		keys := [...]string{"port-number"}
		keyValues := [...]string{strconv.FormatInt(data.TransportTcpPorts[i].PortNumber.ValueInt64(), 10)}

		var r gjson.Result
		res.Get(prefix + "transport.tcp.port-config").ForEach(
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
		if value := r.Get("port-number"); value.Exists() && !data.TransportTcpPorts[i].PortNumber.IsNull() {
			data.TransportTcpPorts[i].PortNumber = types.Int64Value(value.Int())
		} else {
			data.TransportTcpPorts[i].PortNumber = types.Int64Null()
		}
	}
	for i := range data.TransportTlsPorts {
		keys := [...]string{"port-number"}
		keyValues := [...]string{strconv.FormatInt(data.TransportTlsPorts[i].PortNumber.ValueInt64(), 10)}

		var r gjson.Result
		res.Get(prefix + "transport.tls.port").ForEach(
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
		if value := r.Get("port-number"); value.Exists() && !data.TransportTlsPorts[i].PortNumber.IsNull() {
			data.TransportTlsPorts[i].PortNumber = types.Int64Value(value.Int())
		} else {
			data.TransportTlsPorts[i].PortNumber = types.Int64Null()
		}
		if value := r.Get("profile"); value.Exists() && !data.TransportTlsPorts[i].Profile.IsNull() {
			data.TransportTlsPorts[i].Profile = types.StringValue(value.String())
		} else {
			data.TransportTlsPorts[i].Profile = types.StringNull()
		}
	}
}

func (data *LoggingIPv6HostVRFTransportData) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "transport.udp.port-config"); value.Exists() {
		data.TransportUdpPorts = make([]LoggingIPv6HostVRFTransportTransportUdpPorts, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := LoggingIPv6HostVRFTransportTransportUdpPorts{}
			if cValue := v.Get("port-number"); cValue.Exists() {
				item.PortNumber = types.Int64Value(cValue.Int())
			}
			data.TransportUdpPorts = append(data.TransportUdpPorts, item)
			return true
		})
	}
	if value := res.Get(prefix + "transport.tcp.port-config"); value.Exists() {
		data.TransportTcpPorts = make([]LoggingIPv6HostVRFTransportTransportTcpPorts, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := LoggingIPv6HostVRFTransportTransportTcpPorts{}
			if cValue := v.Get("port-number"); cValue.Exists() {
				item.PortNumber = types.Int64Value(cValue.Int())
			}
			data.TransportTcpPorts = append(data.TransportTcpPorts, item)
			return true
		})
	}
	if value := res.Get(prefix + "transport.tls.port"); value.Exists() {
		data.TransportTlsPorts = make([]LoggingIPv6HostVRFTransportTransportTlsPorts, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := LoggingIPv6HostVRFTransportTransportTlsPorts{}
			if cValue := v.Get("port-number"); cValue.Exists() {
				item.PortNumber = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("profile"); cValue.Exists() {
				item.Profile = types.StringValue(cValue.String())
			}
			data.TransportTlsPorts = append(data.TransportTlsPorts, item)
			return true
		})
	}
}

func (data *LoggingIPv6HostVRFTransport) getDeletedItems(ctx context.Context, state LoggingIPv6HostVRFTransport) []string {
	deletedItems := make([]string, 0)
	for i := range state.TransportUdpPorts {
		stateKeyValues := [...]string{strconv.FormatInt(state.TransportUdpPorts[i].PortNumber.ValueInt64(), 10)}

		emptyKeys := true
		if !reflect.ValueOf(state.TransportUdpPorts[i].PortNumber.ValueInt64()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.TransportUdpPorts {
			found = true
			if state.TransportUdpPorts[i].PortNumber.ValueInt64() != data.TransportUdpPorts[j].PortNumber.ValueInt64() {
				found = false
			}
			if found {
				break
			}
		}
		if !found {
			deletedItems = append(deletedItems, fmt.Sprintf("%v/transport/udp/port-config=%v", state.getPath(), strings.Join(stateKeyValues[:], ",")))
		}
	}
	for i := range state.TransportTcpPorts {
		stateKeyValues := [...]string{strconv.FormatInt(state.TransportTcpPorts[i].PortNumber.ValueInt64(), 10)}

		emptyKeys := true
		if !reflect.ValueOf(state.TransportTcpPorts[i].PortNumber.ValueInt64()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.TransportTcpPorts {
			found = true
			if state.TransportTcpPorts[i].PortNumber.ValueInt64() != data.TransportTcpPorts[j].PortNumber.ValueInt64() {
				found = false
			}
			if found {
				break
			}
		}
		if !found {
			deletedItems = append(deletedItems, fmt.Sprintf("%v/transport/tcp/port-config=%v", state.getPath(), strings.Join(stateKeyValues[:], ",")))
		}
	}
	for i := range state.TransportTlsPorts {
		stateKeyValues := [...]string{strconv.FormatInt(state.TransportTlsPorts[i].PortNumber.ValueInt64(), 10)}

		emptyKeys := true
		if !reflect.ValueOf(state.TransportTlsPorts[i].PortNumber.ValueInt64()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.TransportTlsPorts {
			found = true
			if state.TransportTlsPorts[i].PortNumber.ValueInt64() != data.TransportTlsPorts[j].PortNumber.ValueInt64() {
				found = false
			}
			if found {
				if !state.TransportTlsPorts[i].Profile.IsNull() && data.TransportTlsPorts[j].Profile.IsNull() {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/transport/tls/port=%v/profile", state.getPath(), strings.Join(stateKeyValues[:], ",")))
				}
				break
			}
		}
		if !found {
			deletedItems = append(deletedItems, fmt.Sprintf("%v/transport/tls/port=%v", state.getPath(), strings.Join(stateKeyValues[:], ",")))
		}
	}
	return deletedItems
}

func (data *LoggingIPv6HostVRFTransport) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)

	return emptyLeafsDelete
}

func (data *LoggingIPv6HostVRFTransport) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	for i := range data.TransportUdpPorts {
		keyValues := [...]string{strconv.FormatInt(data.TransportUdpPorts[i].PortNumber.ValueInt64(), 10)}

		deletePaths = append(deletePaths, fmt.Sprintf("%v/transport/udp/port-config=%v", data.getPath(), strings.Join(keyValues[:], ",")))
	}
	for i := range data.TransportTcpPorts {
		keyValues := [...]string{strconv.FormatInt(data.TransportTcpPorts[i].PortNumber.ValueInt64(), 10)}

		deletePaths = append(deletePaths, fmt.Sprintf("%v/transport/tcp/port-config=%v", data.getPath(), strings.Join(keyValues[:], ",")))
	}
	for i := range data.TransportTlsPorts {
		keyValues := [...]string{strconv.FormatInt(data.TransportTlsPorts[i].PortNumber.ValueInt64(), 10)}

		deletePaths = append(deletePaths, fmt.Sprintf("%v/transport/tls/port=%v", data.getPath(), strings.Join(keyValues[:], ",")))
	}
	return deletePaths
}
