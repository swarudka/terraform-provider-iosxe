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
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/CiscoDevNet/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type AAAAccounting struct {
	Device                 types.String           `tfsdk:"device"`
	Id                     types.String           `tfsdk:"id"`
	DeleteMode             types.String           `tfsdk:"delete_mode"`
	UpdateNewinfoPeriodic  types.Int64            `tfsdk:"update_newinfo_periodic"`
	IdentityStartStopGroup types.String           `tfsdk:"identity_start_stop_group"`
	Exec                   []AAAAccountingExec    `tfsdk:"exec"`
	Network                []AAAAccountingNetwork `tfsdk:"network"`
	SystemGuaranteeFirst   types.Bool             `tfsdk:"system_guarantee_first"`
}

type AAAAccountingData struct {
	Device                 types.String           `tfsdk:"device"`
	Id                     types.String           `tfsdk:"id"`
	UpdateNewinfoPeriodic  types.Int64            `tfsdk:"update_newinfo_periodic"`
	IdentityStartStopGroup types.String           `tfsdk:"identity_start_stop_group"`
	Exec                   []AAAAccountingExec    `tfsdk:"exec"`
	Network                []AAAAccountingNetwork `tfsdk:"network"`
	SystemGuaranteeFirst   types.Bool             `tfsdk:"system_guarantee_first"`
}
type AAAAccountingExec struct {
	Name   types.String `tfsdk:"name"`
	Group1 types.String `tfsdk:"group1"`
}
type AAAAccountingNetwork struct {
	Id     types.String `tfsdk:"id"`
	Group1 types.String `tfsdk:"group1"`
	Group2 types.String `tfsdk:"group2"`
}

func (data AAAAccounting) getPath() string {
	return "Cisco-IOS-XE-native:native/aaa/Cisco-IOS-XE-aaa:accounting"
}

func (data AAAAccountingData) getPath() string {
	return "Cisco-IOS-XE-native:native/aaa/Cisco-IOS-XE-aaa:accounting"
}

// if last path element has a key -> remove it
func (data AAAAccounting) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data AAAAccounting) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.UpdateNewinfoPeriodic.IsNull() && !data.UpdateNewinfoPeriodic.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"update.newinfo.periodic", strconv.FormatInt(data.UpdateNewinfoPeriodic.ValueInt64(), 10))
	}
	if !data.IdentityStartStopGroup.IsNull() && !data.IdentityStartStopGroup.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"identity.default.start-stop.group", data.IdentityStartStopGroup.ValueString())
	}
	if !data.SystemGuaranteeFirst.IsNull() && !data.SystemGuaranteeFirst.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"system.guarantee-first", data.SystemGuaranteeFirst.ValueBool())
	}
	if len(data.Exec) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"exec", []interface{}{})
		for index, item := range data.Exec {
			if !item.Name.IsNull() && !item.Name.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"exec"+"."+strconv.Itoa(index)+"."+"name", item.Name.ValueString())
			}
			if !item.Group1.IsNull() && !item.Group1.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"exec"+"."+strconv.Itoa(index)+"."+"start-stop.group1.group", item.Group1.ValueString())
			}
		}
	}
	if len(data.Network) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"network", []interface{}{})
		for index, item := range data.Network {
			if !item.Id.IsNull() && !item.Id.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"network"+"."+strconv.Itoa(index)+"."+"id", item.Id.ValueString())
			}
			if !item.Group1.IsNull() && !item.Group1.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"network"+"."+strconv.Itoa(index)+"."+"start-stop.group-config.group1.group", item.Group1.ValueString())
			}
			if !item.Group2.IsNull() && !item.Group2.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"network"+"."+strconv.Itoa(index)+"."+"start-stop.group-config.group2.group", item.Group2.ValueString())
			}
		}
	}
	return body
}

func (data *AAAAccounting) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "update.newinfo.periodic"); value.Exists() && !data.UpdateNewinfoPeriodic.IsNull() {
		data.UpdateNewinfoPeriodic = types.Int64Value(value.Int())
	} else {
		data.UpdateNewinfoPeriodic = types.Int64Null()
	}
	if value := res.Get(prefix + "identity.default.start-stop.group"); value.Exists() && !data.IdentityStartStopGroup.IsNull() {
		data.IdentityStartStopGroup = types.StringValue(value.String())
	} else {
		data.IdentityStartStopGroup = types.StringNull()
	}
	for i := range data.Exec {
		keys := [...]string{"name"}
		keyValues := [...]string{data.Exec[i].Name.ValueString()}

		var r gjson.Result
		res.Get(prefix + "exec").ForEach(
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
		if value := r.Get("name"); value.Exists() && !data.Exec[i].Name.IsNull() {
			data.Exec[i].Name = types.StringValue(value.String())
		} else {
			data.Exec[i].Name = types.StringNull()
		}
		if value := r.Get("start-stop.group1.group"); value.Exists() && !data.Exec[i].Group1.IsNull() {
			data.Exec[i].Group1 = types.StringValue(value.String())
		} else {
			data.Exec[i].Group1 = types.StringNull()
		}
	}
	for i := range data.Network {
		keys := [...]string{"id"}
		keyValues := [...]string{data.Network[i].Id.ValueString()}

		var r gjson.Result
		res.Get(prefix + "network").ForEach(
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
		if value := r.Get("id"); value.Exists() && !data.Network[i].Id.IsNull() {
			data.Network[i].Id = types.StringValue(value.String())
		} else {
			data.Network[i].Id = types.StringNull()
		}
		if value := r.Get("start-stop.group-config.group1.group"); value.Exists() && !data.Network[i].Group1.IsNull() {
			data.Network[i].Group1 = types.StringValue(value.String())
		} else {
			data.Network[i].Group1 = types.StringNull()
		}
		if value := r.Get("start-stop.group-config.group2.group"); value.Exists() && !data.Network[i].Group2.IsNull() {
			data.Network[i].Group2 = types.StringValue(value.String())
		} else {
			data.Network[i].Group2 = types.StringNull()
		}
	}
	if value := res.Get(prefix + "system.guarantee-first"); !data.SystemGuaranteeFirst.IsNull() {
		if value.Exists() {
			data.SystemGuaranteeFirst = types.BoolValue(value.Bool())
		}
	} else {
		data.SystemGuaranteeFirst = types.BoolNull()
	}
}

func (data *AAAAccountingData) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "update.newinfo.periodic"); value.Exists() {
		data.UpdateNewinfoPeriodic = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "identity.default.start-stop.group"); value.Exists() {
		data.IdentityStartStopGroup = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "exec"); value.Exists() {
		data.Exec = make([]AAAAccountingExec, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := AAAAccountingExec{}
			if cValue := v.Get("name"); cValue.Exists() {
				item.Name = types.StringValue(cValue.String())
			}
			if cValue := v.Get("start-stop.group1.group"); cValue.Exists() {
				item.Group1 = types.StringValue(cValue.String())
			}
			data.Exec = append(data.Exec, item)
			return true
		})
	}
	if value := res.Get(prefix + "network"); value.Exists() {
		data.Network = make([]AAAAccountingNetwork, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := AAAAccountingNetwork{}
			if cValue := v.Get("id"); cValue.Exists() {
				item.Id = types.StringValue(cValue.String())
			}
			if cValue := v.Get("start-stop.group-config.group1.group"); cValue.Exists() {
				item.Group1 = types.StringValue(cValue.String())
			}
			if cValue := v.Get("start-stop.group-config.group2.group"); cValue.Exists() {
				item.Group2 = types.StringValue(cValue.String())
			}
			data.Network = append(data.Network, item)
			return true
		})
	}
	if value := res.Get(prefix + "system.guarantee-first"); value.Exists() {
		data.SystemGuaranteeFirst = types.BoolValue(value.Bool())
	} else {
		data.SystemGuaranteeFirst = types.BoolValue(false)
	}
}

func (data *AAAAccounting) getDeletedListItems(ctx context.Context, state AAAAccounting) []string {
	deletedListItems := make([]string, 0)
	for i := range state.Exec {
		stateKeyValues := [...]string{state.Exec[i].Name.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.Exec[i].Name.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.Exec {
			found = true
			if state.Exec[i].Name.ValueString() != data.Exec[j].Name.ValueString() {
				found = false
			}
			if found {
				break
			}
		}
		if !found {
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/exec=%v", state.getPath(), strings.Join(stateKeyValues[:], ",")))
		}
	}
	for i := range state.Network {
		stateKeyValues := [...]string{state.Network[i].Id.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.Network[i].Id.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.Network {
			found = true
			if state.Network[i].Id.ValueString() != data.Network[j].Id.ValueString() {
				found = false
			}
			if found {
				break
			}
		}
		if !found {
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/network=%v", state.getPath(), strings.Join(stateKeyValues[:], ",")))
		}
	}
	return deletedListItems
}

func (data *AAAAccounting) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)

	return emptyLeafsDelete
}

func (data *AAAAccounting) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.UpdateNewinfoPeriodic.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/update/newinfo/periodic", data.getPath()))
	}
	if !data.IdentityStartStopGroup.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/identity/default/start-stop/group", data.getPath()))
	}
	for i := range data.Exec {
		keyValues := [...]string{data.Exec[i].Name.ValueString()}

		deletePaths = append(deletePaths, fmt.Sprintf("%v/exec=%v", data.getPath(), strings.Join(keyValues[:], ",")))
	}
	for i := range data.Network {
		keyValues := [...]string{data.Network[i].Id.ValueString()}

		deletePaths = append(deletePaths, fmt.Sprintf("%v/network=%v", data.getPath(), strings.Join(keyValues[:], ",")))
	}
	if !data.SystemGuaranteeFirst.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/system/guarantee-first", data.getPath()))
	}
	return deletePaths
}