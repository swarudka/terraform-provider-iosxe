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

type ClassMap struct {
	Device                                    types.String                             `tfsdk:"device"`
	Id                                        types.String                             `tfsdk:"id"`
	Name                                      types.String                             `tfsdk:"name"`
	Type                                      types.String                             `tfsdk:"type"`
	Subscriber                                types.Bool                               `tfsdk:"subscriber"`
	Prematch                                  types.String                             `tfsdk:"prematch"`
	MatchAuthorizationStatusAuthorized        types.Bool                               `tfsdk:"match_authorization_status_authorized"`
	MatchResultTypeAaaTimeout                 types.Bool                               `tfsdk:"match_result_type_aaa_timeout"`
	MatchAuthorizationStatusUnauthorized      types.Bool                               `tfsdk:"match_authorization_status_unauthorized"`
	MatchActivatedServiceTemplates            []ClassMapMatchActivatedServiceTemplates `tfsdk:"match_activated_service_templates"`
	MatchAuthorizingMethodPriorityGreaterThan types.List                               `tfsdk:"match_authorizing_method_priority_greater_than"`
	MatchMethodDot1x                          types.Bool                               `tfsdk:"match_method_dot1x"`
	MatchResultTypeMethodDot1xAuthoritative   types.Bool                               `tfsdk:"match_result_type_method_dot1x_authoritative"`
	MatchResultTypeMethodDot1xAgentNotFound   types.Bool                               `tfsdk:"match_result_type_method_dot1x_agent_not_found"`
	MatchResultTypeMethodDot1xMethodTimeout   types.Bool                               `tfsdk:"match_result_type_method_dot1x_method_timeout"`
	MatchMethodMab                            types.Bool                               `tfsdk:"match_method_mab"`
	MatchResultTypeMethodMabAuthoritative     types.Bool                               `tfsdk:"match_result_type_method_mab_authoritative"`
	Description                               types.String                             `tfsdk:"description"`
}

type ClassMapData struct {
	Device                                    types.String                             `tfsdk:"device"`
	Id                                        types.String                             `tfsdk:"id"`
	Name                                      types.String                             `tfsdk:"name"`
	Type                                      types.String                             `tfsdk:"type"`
	Subscriber                                types.Bool                               `tfsdk:"subscriber"`
	Prematch                                  types.String                             `tfsdk:"prematch"`
	MatchAuthorizationStatusAuthorized        types.Bool                               `tfsdk:"match_authorization_status_authorized"`
	MatchResultTypeAaaTimeout                 types.Bool                               `tfsdk:"match_result_type_aaa_timeout"`
	MatchAuthorizationStatusUnauthorized      types.Bool                               `tfsdk:"match_authorization_status_unauthorized"`
	MatchActivatedServiceTemplates            []ClassMapMatchActivatedServiceTemplates `tfsdk:"match_activated_service_templates"`
	MatchAuthorizingMethodPriorityGreaterThan types.List                               `tfsdk:"match_authorizing_method_priority_greater_than"`
	MatchMethodDot1x                          types.Bool                               `tfsdk:"match_method_dot1x"`
	MatchResultTypeMethodDot1xAuthoritative   types.Bool                               `tfsdk:"match_result_type_method_dot1x_authoritative"`
	MatchResultTypeMethodDot1xAgentNotFound   types.Bool                               `tfsdk:"match_result_type_method_dot1x_agent_not_found"`
	MatchResultTypeMethodDot1xMethodTimeout   types.Bool                               `tfsdk:"match_result_type_method_dot1x_method_timeout"`
	MatchMethodMab                            types.Bool                               `tfsdk:"match_method_mab"`
	MatchResultTypeMethodMabAuthoritative     types.Bool                               `tfsdk:"match_result_type_method_mab_authoritative"`
	Description                               types.String                             `tfsdk:"description"`
}
type ClassMapMatchActivatedServiceTemplates struct {
	ServiceName types.String `tfsdk:"service_name"`
}

func (data ClassMap) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/policy/Cisco-IOS-XE-policy:class-map=%v", url.QueryEscape(fmt.Sprintf("%v", data.Name.ValueString())))
}

func (data ClassMapData) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/policy/Cisco-IOS-XE-policy:class-map=%v", url.QueryEscape(fmt.Sprintf("%v", data.Name.ValueString())))
}

// if last path element has a key -> remove it
func (data ClassMap) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data ClassMap) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"name", data.Name.ValueString())
	}
	if !data.Type.IsNull() && !data.Type.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"type", data.Type.ValueString())
	}
	if !data.Subscriber.IsNull() && !data.Subscriber.IsUnknown() {
		if data.Subscriber.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"subscriber", map[string]string{})
		}
	}
	if !data.Prematch.IsNull() && !data.Prematch.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"prematch", data.Prematch.ValueString())
	}
	if !data.MatchAuthorizationStatusAuthorized.IsNull() && !data.MatchAuthorizationStatusAuthorized.IsUnknown() {
		if data.MatchAuthorizationStatusAuthorized.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"match.authorization-status.authorized", map[string]string{})
		}
	}
	if !data.MatchResultTypeAaaTimeout.IsNull() && !data.MatchResultTypeAaaTimeout.IsUnknown() {
		if data.MatchResultTypeAaaTimeout.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"match.result-type.aaa-timeout", map[string]string{})
		}
	}
	if !data.MatchAuthorizationStatusUnauthorized.IsNull() && !data.MatchAuthorizationStatusUnauthorized.IsUnknown() {
		if data.MatchAuthorizationStatusUnauthorized.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"match.authorization-status.unauthorized", map[string]string{})
		}
	}
	if !data.MatchAuthorizingMethodPriorityGreaterThan.IsNull() && !data.MatchAuthorizingMethodPriorityGreaterThan.IsUnknown() {
		var values []int
		data.MatchAuthorizingMethodPriorityGreaterThan.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"match.authorizing-method-priority.greater-than", values)
	}
	if !data.MatchMethodDot1x.IsNull() && !data.MatchMethodDot1x.IsUnknown() {
		if data.MatchMethodDot1x.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"match.method.dot1x", map[string]string{})
		}
	}
	if !data.MatchResultTypeMethodDot1xAuthoritative.IsNull() && !data.MatchResultTypeMethodDot1xAuthoritative.IsUnknown() {
		if data.MatchResultTypeMethodDot1xAuthoritative.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"match.result-type.method.dot1x.authoritative", map[string]string{})
		}
	}
	if !data.MatchResultTypeMethodDot1xAgentNotFound.IsNull() && !data.MatchResultTypeMethodDot1xAgentNotFound.IsUnknown() {
		if data.MatchResultTypeMethodDot1xAgentNotFound.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"match.result-type.method.dot1x.agent-not-found", map[string]string{})
		}
	}
	if !data.MatchResultTypeMethodDot1xMethodTimeout.IsNull() && !data.MatchResultTypeMethodDot1xMethodTimeout.IsUnknown() {
		if data.MatchResultTypeMethodDot1xMethodTimeout.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"match.result-type.method.dot1x.method-timeout", map[string]string{})
		}
	}
	if !data.MatchMethodMab.IsNull() && !data.MatchMethodMab.IsUnknown() {
		if data.MatchMethodMab.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"match.method.mab", map[string]string{})
		}
	}
	if !data.MatchResultTypeMethodMabAuthoritative.IsNull() && !data.MatchResultTypeMethodMabAuthoritative.IsUnknown() {
		if data.MatchResultTypeMethodMabAuthoritative.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"match.result-type.method.mab.authoritative", map[string]string{})
		}
	}
	if !data.Description.IsNull() && !data.Description.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"description", data.Description.ValueString())
	}
	if len(data.MatchActivatedServiceTemplates) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"match.activated-service-template", []interface{}{})
		for index, item := range data.MatchActivatedServiceTemplates {
			if !item.ServiceName.IsNull() && !item.ServiceName.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"match.activated-service-template"+"."+strconv.Itoa(index)+"."+"service-name", item.ServiceName.ValueString())
			}
		}
	}
	return body
}

func (data *ClassMap) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get(prefix + "type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get(prefix + "subscriber"); !data.Subscriber.IsNull() {
		if value.Exists() {
			data.Subscriber = types.BoolValue(true)
		} else {
			data.Subscriber = types.BoolValue(false)
		}
	} else {
		data.Subscriber = types.BoolNull()
	}
	if value := res.Get(prefix + "prematch"); value.Exists() && !data.Prematch.IsNull() {
		data.Prematch = types.StringValue(value.String())
	} else {
		data.Prematch = types.StringNull()
	}
	if value := res.Get(prefix + "match.authorization-status.authorized"); !data.MatchAuthorizationStatusAuthorized.IsNull() {
		if value.Exists() {
			data.MatchAuthorizationStatusAuthorized = types.BoolValue(true)
		} else {
			data.MatchAuthorizationStatusAuthorized = types.BoolValue(false)
		}
	} else {
		data.MatchAuthorizationStatusAuthorized = types.BoolNull()
	}
	if value := res.Get(prefix + "match.result-type.aaa-timeout"); !data.MatchResultTypeAaaTimeout.IsNull() {
		if value.Exists() {
			data.MatchResultTypeAaaTimeout = types.BoolValue(true)
		} else {
			data.MatchResultTypeAaaTimeout = types.BoolValue(false)
		}
	} else {
		data.MatchResultTypeAaaTimeout = types.BoolNull()
	}
	if value := res.Get(prefix + "match.authorization-status.unauthorized"); !data.MatchAuthorizationStatusUnauthorized.IsNull() {
		if value.Exists() {
			data.MatchAuthorizationStatusUnauthorized = types.BoolValue(true)
		} else {
			data.MatchAuthorizationStatusUnauthorized = types.BoolValue(false)
		}
	} else {
		data.MatchAuthorizationStatusUnauthorized = types.BoolNull()
	}
	for i := range data.MatchActivatedServiceTemplates {
		keys := [...]string{"service-name"}
		keyValues := [...]string{data.MatchActivatedServiceTemplates[i].ServiceName.ValueString()}

		var r gjson.Result
		res.Get(prefix + "match.activated-service-template").ForEach(
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
		if value := r.Get("service-name"); value.Exists() && !data.MatchActivatedServiceTemplates[i].ServiceName.IsNull() {
			data.MatchActivatedServiceTemplates[i].ServiceName = types.StringValue(value.String())
		} else {
			data.MatchActivatedServiceTemplates[i].ServiceName = types.StringNull()
		}
	}
	if value := res.Get(prefix + "match.authorizing-method-priority.greater-than"); value.Exists() && !data.MatchAuthorizingMethodPriorityGreaterThan.IsNull() {
		data.MatchAuthorizingMethodPriorityGreaterThan = helpers.GetInt64List(value.Array())
	} else {
		data.MatchAuthorizingMethodPriorityGreaterThan = types.ListNull(types.Int64Type)
	}
	if value := res.Get(prefix + "match.method.dot1x"); !data.MatchMethodDot1x.IsNull() {
		if value.Exists() {
			data.MatchMethodDot1x = types.BoolValue(true)
		} else {
			data.MatchMethodDot1x = types.BoolValue(false)
		}
	} else {
		data.MatchMethodDot1x = types.BoolNull()
	}
	if value := res.Get(prefix + "match.result-type.method.dot1x.authoritative"); !data.MatchResultTypeMethodDot1xAuthoritative.IsNull() {
		if value.Exists() {
			data.MatchResultTypeMethodDot1xAuthoritative = types.BoolValue(true)
		} else {
			data.MatchResultTypeMethodDot1xAuthoritative = types.BoolValue(false)
		}
	} else {
		data.MatchResultTypeMethodDot1xAuthoritative = types.BoolNull()
	}
	if value := res.Get(prefix + "match.result-type.method.dot1x.agent-not-found"); !data.MatchResultTypeMethodDot1xAgentNotFound.IsNull() {
		if value.Exists() {
			data.MatchResultTypeMethodDot1xAgentNotFound = types.BoolValue(true)
		} else {
			data.MatchResultTypeMethodDot1xAgentNotFound = types.BoolValue(false)
		}
	} else {
		data.MatchResultTypeMethodDot1xAgentNotFound = types.BoolNull()
	}
	if value := res.Get(prefix + "match.result-type.method.dot1x.method-timeout"); !data.MatchResultTypeMethodDot1xMethodTimeout.IsNull() {
		if value.Exists() {
			data.MatchResultTypeMethodDot1xMethodTimeout = types.BoolValue(true)
		} else {
			data.MatchResultTypeMethodDot1xMethodTimeout = types.BoolValue(false)
		}
	} else {
		data.MatchResultTypeMethodDot1xMethodTimeout = types.BoolNull()
	}
	if value := res.Get(prefix + "match.method.mab"); !data.MatchMethodMab.IsNull() {
		if value.Exists() {
			data.MatchMethodMab = types.BoolValue(true)
		} else {
			data.MatchMethodMab = types.BoolValue(false)
		}
	} else {
		data.MatchMethodMab = types.BoolNull()
	}
	if value := res.Get(prefix + "match.result-type.method.mab.authoritative"); !data.MatchResultTypeMethodMabAuthoritative.IsNull() {
		if value.Exists() {
			data.MatchResultTypeMethodMabAuthoritative = types.BoolValue(true)
		} else {
			data.MatchResultTypeMethodMabAuthoritative = types.BoolValue(false)
		}
	} else {
		data.MatchResultTypeMethodMabAuthoritative = types.BoolNull()
	}
	if value := res.Get(prefix + "description"); value.Exists() && !data.Description.IsNull() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
}

func (data *ClassMapData) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "subscriber"); value.Exists() {
		data.Subscriber = types.BoolValue(true)
	} else {
		data.Subscriber = types.BoolValue(false)
	}
	if value := res.Get(prefix + "prematch"); value.Exists() {
		data.Prematch = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "match.authorization-status.authorized"); value.Exists() {
		data.MatchAuthorizationStatusAuthorized = types.BoolValue(true)
	} else {
		data.MatchAuthorizationStatusAuthorized = types.BoolValue(false)
	}
	if value := res.Get(prefix + "match.result-type.aaa-timeout"); value.Exists() {
		data.MatchResultTypeAaaTimeout = types.BoolValue(true)
	} else {
		data.MatchResultTypeAaaTimeout = types.BoolValue(false)
	}
	if value := res.Get(prefix + "match.authorization-status.unauthorized"); value.Exists() {
		data.MatchAuthorizationStatusUnauthorized = types.BoolValue(true)
	} else {
		data.MatchAuthorizationStatusUnauthorized = types.BoolValue(false)
	}
	if value := res.Get(prefix + "match.activated-service-template"); value.Exists() {
		data.MatchActivatedServiceTemplates = make([]ClassMapMatchActivatedServiceTemplates, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ClassMapMatchActivatedServiceTemplates{}
			if cValue := v.Get("service-name"); cValue.Exists() {
				item.ServiceName = types.StringValue(cValue.String())
			}
			data.MatchActivatedServiceTemplates = append(data.MatchActivatedServiceTemplates, item)
			return true
		})
	}
	if value := res.Get(prefix + "match.authorizing-method-priority.greater-than"); value.Exists() {
		data.MatchAuthorizingMethodPriorityGreaterThan = helpers.GetInt64List(value.Array())
	} else {
		data.MatchAuthorizingMethodPriorityGreaterThan = types.ListNull(types.Int64Type)
	}
	if value := res.Get(prefix + "match.method.dot1x"); value.Exists() {
		data.MatchMethodDot1x = types.BoolValue(true)
	} else {
		data.MatchMethodDot1x = types.BoolValue(false)
	}
	if value := res.Get(prefix + "match.result-type.method.dot1x.authoritative"); value.Exists() {
		data.MatchResultTypeMethodDot1xAuthoritative = types.BoolValue(true)
	} else {
		data.MatchResultTypeMethodDot1xAuthoritative = types.BoolValue(false)
	}
	if value := res.Get(prefix + "match.result-type.method.dot1x.agent-not-found"); value.Exists() {
		data.MatchResultTypeMethodDot1xAgentNotFound = types.BoolValue(true)
	} else {
		data.MatchResultTypeMethodDot1xAgentNotFound = types.BoolValue(false)
	}
	if value := res.Get(prefix + "match.result-type.method.dot1x.method-timeout"); value.Exists() {
		data.MatchResultTypeMethodDot1xMethodTimeout = types.BoolValue(true)
	} else {
		data.MatchResultTypeMethodDot1xMethodTimeout = types.BoolValue(false)
	}
	if value := res.Get(prefix + "match.method.mab"); value.Exists() {
		data.MatchMethodMab = types.BoolValue(true)
	} else {
		data.MatchMethodMab = types.BoolValue(false)
	}
	if value := res.Get(prefix + "match.result-type.method.mab.authoritative"); value.Exists() {
		data.MatchResultTypeMethodMabAuthoritative = types.BoolValue(true)
	} else {
		data.MatchResultTypeMethodMabAuthoritative = types.BoolValue(false)
	}
	if value := res.Get(prefix + "description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	}
}

func (data *ClassMap) getDeletedItems(ctx context.Context, state ClassMap) []string {
	deletedItems := make([]string, 0)
	if !state.Type.IsNull() && data.Type.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/type", state.getPath()))
	}
	if !state.Subscriber.IsNull() && data.Subscriber.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/subscriber", state.getPath()))
	}
	if !state.Prematch.IsNull() && data.Prematch.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/prematch", state.getPath()))
	}
	if !state.MatchAuthorizationStatusAuthorized.IsNull() && data.MatchAuthorizationStatusAuthorized.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/match/authorization-status/authorized", state.getPath()))
	}
	if !state.MatchResultTypeAaaTimeout.IsNull() && data.MatchResultTypeAaaTimeout.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/match/result-type/aaa-timeout", state.getPath()))
	}
	if !state.MatchAuthorizationStatusUnauthorized.IsNull() && data.MatchAuthorizationStatusUnauthorized.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/match/authorization-status/unauthorized", state.getPath()))
	}
	for i := range state.MatchActivatedServiceTemplates {
		stateKeyValues := [...]string{state.MatchActivatedServiceTemplates[i].ServiceName.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.MatchActivatedServiceTemplates[i].ServiceName.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.MatchActivatedServiceTemplates {
			found = true
			if state.MatchActivatedServiceTemplates[i].ServiceName.ValueString() != data.MatchActivatedServiceTemplates[j].ServiceName.ValueString() {
				found = false
			}
			if found {
				break
			}
		}
		if !found {
			deletedItems = append(deletedItems, fmt.Sprintf("%v/match/activated-service-template=%v", state.getPath(), strings.Join(stateKeyValues[:], ",")))
		}
	}
	if !state.MatchAuthorizingMethodPriorityGreaterThan.IsNull() {
		if data.MatchAuthorizingMethodPriorityGreaterThan.IsNull() {
			deletedItems = append(deletedItems, fmt.Sprintf("%v/match/authorizing-method-priority/greater-than", state.getPath()))
		} else {
			var dataValues, stateValues []int
			data.MatchAuthorizingMethodPriorityGreaterThan.ElementsAs(ctx, &dataValues, false)
			state.MatchAuthorizingMethodPriorityGreaterThan.ElementsAs(ctx, &stateValues, false)
			for _, v := range stateValues {
				found := false
				for _, vv := range dataValues {
					if v == vv {
						found = true
						break
					}
				}
				if !found {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/match/authorizing-method-priority/greater-than=%v", state.getPath(), v))
				}
			}
		}
	}
	if !state.MatchMethodDot1x.IsNull() && data.MatchMethodDot1x.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/match/method/dot1x", state.getPath()))
	}
	if !state.MatchResultTypeMethodDot1xAuthoritative.IsNull() && data.MatchResultTypeMethodDot1xAuthoritative.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/match/result-type/method/dot1x/authoritative", state.getPath()))
	}
	if !state.MatchResultTypeMethodDot1xAgentNotFound.IsNull() && data.MatchResultTypeMethodDot1xAgentNotFound.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/match/result-type/method/dot1x/agent-not-found", state.getPath()))
	}
	if !state.MatchResultTypeMethodDot1xMethodTimeout.IsNull() && data.MatchResultTypeMethodDot1xMethodTimeout.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/match/result-type/method/dot1x/method-timeout", state.getPath()))
	}
	if !state.MatchMethodMab.IsNull() && data.MatchMethodMab.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/match/method/mab", state.getPath()))
	}
	if !state.MatchResultTypeMethodMabAuthoritative.IsNull() && data.MatchResultTypeMethodMabAuthoritative.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/match/result-type/method/mab/authoritative", state.getPath()))
	}
	if !state.Description.IsNull() && data.Description.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/description", state.getPath()))
	}
	return deletedItems
}

func (data *ClassMap) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	if !data.Subscriber.IsNull() && !data.Subscriber.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/subscriber", data.getPath()))
	}
	if !data.MatchAuthorizationStatusAuthorized.IsNull() && !data.MatchAuthorizationStatusAuthorized.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/match/authorization-status/authorized", data.getPath()))
	}
	if !data.MatchResultTypeAaaTimeout.IsNull() && !data.MatchResultTypeAaaTimeout.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/match/result-type/aaa-timeout", data.getPath()))
	}
	if !data.MatchAuthorizationStatusUnauthorized.IsNull() && !data.MatchAuthorizationStatusUnauthorized.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/match/authorization-status/unauthorized", data.getPath()))
	}

	if !data.MatchMethodDot1x.IsNull() && !data.MatchMethodDot1x.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/match/method/dot1x", data.getPath()))
	}
	if !data.MatchResultTypeMethodDot1xAuthoritative.IsNull() && !data.MatchResultTypeMethodDot1xAuthoritative.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/match/result-type/method/dot1x/authoritative", data.getPath()))
	}
	if !data.MatchResultTypeMethodDot1xAgentNotFound.IsNull() && !data.MatchResultTypeMethodDot1xAgentNotFound.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/match/result-type/method/dot1x/agent-not-found", data.getPath()))
	}
	if !data.MatchResultTypeMethodDot1xMethodTimeout.IsNull() && !data.MatchResultTypeMethodDot1xMethodTimeout.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/match/result-type/method/dot1x/method-timeout", data.getPath()))
	}
	if !data.MatchMethodMab.IsNull() && !data.MatchMethodMab.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/match/method/mab", data.getPath()))
	}
	if !data.MatchResultTypeMethodMabAuthoritative.IsNull() && !data.MatchResultTypeMethodMabAuthoritative.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/match/result-type/method/mab/authoritative", data.getPath()))
	}
	return emptyLeafsDelete
}

func (data *ClassMap) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.Type.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/type", data.getPath()))
	}
	if !data.Subscriber.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/subscriber", data.getPath()))
	}
	if !data.Prematch.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/prematch", data.getPath()))
	}
	if !data.MatchAuthorizationStatusAuthorized.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/match/authorization-status/authorized", data.getPath()))
	}
	if !data.MatchResultTypeAaaTimeout.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/match/result-type/aaa-timeout", data.getPath()))
	}
	if !data.MatchAuthorizationStatusUnauthorized.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/match/authorization-status/unauthorized", data.getPath()))
	}
	for i := range data.MatchActivatedServiceTemplates {
		keyValues := [...]string{data.MatchActivatedServiceTemplates[i].ServiceName.ValueString()}

		deletePaths = append(deletePaths, fmt.Sprintf("%v/match/activated-service-template=%v", data.getPath(), strings.Join(keyValues[:], ",")))
	}
	if !data.MatchAuthorizingMethodPriorityGreaterThan.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/match/authorizing-method-priority/greater-than", data.getPath()))
	}
	if !data.MatchMethodDot1x.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/match/method/dot1x", data.getPath()))
	}
	if !data.MatchResultTypeMethodDot1xAuthoritative.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/match/result-type/method/dot1x/authoritative", data.getPath()))
	}
	if !data.MatchResultTypeMethodDot1xAgentNotFound.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/match/result-type/method/dot1x/agent-not-found", data.getPath()))
	}
	if !data.MatchResultTypeMethodDot1xMethodTimeout.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/match/result-type/method/dot1x/method-timeout", data.getPath()))
	}
	if !data.MatchMethodMab.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/match/method/mab", data.getPath()))
	}
	if !data.MatchResultTypeMethodMabAuthoritative.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/match/result-type/method/mab/authoritative", data.getPath()))
	}
	if !data.Description.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/description", data.getPath()))
	}
	return deletePaths
}
