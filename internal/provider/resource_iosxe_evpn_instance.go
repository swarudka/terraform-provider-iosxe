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

	"github.com/CiscoDevNet/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-restconf"
)

func NewEVPNInstanceResource() resource.Resource {
	return &EVPNInstanceResource{}
}

type EVPNInstanceResource struct {
	clients map[string]*restconf.Client
}

func (r *EVPNInstanceResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_evpn_instance"
}

func (r *EVPNInstanceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This resource can manage the EVPN Instance configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the object.",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"evpn_instance_num": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("evpn instance number").AddIntegerRangeDescription(1, 65535).String,
				Required:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
				},
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
			},
			"vlan_based_replication_type_ingress": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Ingress replication").String,
				Optional:            true,
			},
			"vlan_based_replication_type_static": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Static replication").String,
				Optional:            true,
			},
			"vlan_based_replication_type_p2mp": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("p2mp replication").String,
				Optional:            true,
			},
			"vlan_based_replication_type_mp2mp": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("mp2mp replication").String,
				Optional:            true,
			},
			"vlan_based_encapsulation": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Data encapsulation method").AddStringEnumDescription("mpls", "vxlan").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("mpls", "vxlan"),
				},
			},
			"vlan_based_auto_route_target": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Automatically set a route-target").String,
				Optional:            true,
			},
			"vlan_based_rd": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ASN:nn or IP-address:nn").String,
				Optional:            true,
			},
			"vlan_based_route_target": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ASN:nn or IP-address:nn").String,
				Optional:            true,
			},
			"vlan_based_route_target_both": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ASN:nn or IP-address:nn").String,
				Optional:            true,
			},
			"vlan_based_route_target_import": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ASN:nn or IP-address:nn (DEPRECATED, use rt-value-entry)").String,
				Optional:            true,
			},
			"vlan_based_route_target_export": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ASN:nn or IP-address:nn (DEPRECATED, use rt-value-entry)").String,
				Optional:            true,
			},
			"vlan_based_ip_local_learning_disable": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Disable IP local learning from dataplane").String,
				Optional:            true,
			},
			"vlan_based_ip_local_learning_enable": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable IP local learning from dataplane").String,
				Optional:            true,
			},
			"vlan_based_default_gateway_advertise": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Advertise Default Gateway MAC/IP routes").AddStringEnumDescription("disable", "enable").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("disable", "enable"),
				},
			},
			"vlan_based_re_originate_route_type5": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Re-originate route-type 5").String,
				Optional:            true,
			},
		},
	}
}

func (r *EVPNInstanceResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.clients = req.ProviderData.(map[string]*restconf.Client)
}

func (r *EVPNInstanceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan EVPNInstance

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if _, ok := r.clients[plan.Device.ValueString()]; !ok {
		resp.Diagnostics.AddAttributeError(path.Root("device"), "Invalid device", fmt.Sprintf("Device '%s' does not exist in provider configuration.", plan.Device.ValueString()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.getPath()))

	// Create object
	body := plan.toBody(ctx)

	emptyLeafsDelete := plan.getEmptyLeafsDelete(ctx)
	tflog.Debug(ctx, fmt.Sprintf("List of empty leafs to delete: %+v", emptyLeafsDelete))

	if YangPatch {
		edits := []restconf.YangPatchEdit{restconf.NewYangPatchEdit("merge", plan.getPath(), restconf.Body{Str: body})}
		for _, i := range emptyLeafsDelete {
			edits = append(edits, restconf.NewYangPatchEdit("remove", i, restconf.Body{}))
		}
		_, err := r.clients[plan.Device.ValueString()].YangPatchData("", "1", "", edits)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object, got error: %s", err))
			return
		}
	} else {
		res, err := r.clients[plan.Device.ValueString()].PatchData(plan.getPathShort(), body)
		if len(res.Errors.Error) > 0 && res.Errors.Error[0].ErrorMessage == "patch to a nonexistent resource" {
			_, err = r.clients[plan.Device.ValueString()].PutData(plan.getPath(), body)
		}
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PATCH), got error: %s", err))
			return
		}
		for _, i := range emptyLeafsDelete {
			res, err := r.clients[plan.Device.ValueString()].DeleteData(i)
			if err != nil && res.StatusCode != 404 {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object, got error: %s", err))
				return
			}
		}
	}

	plan.Id = types.StringValue(plan.getPath())

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.getPath()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *EVPNInstanceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state EVPNInstance

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if _, ok := r.clients[state.Device.ValueString()]; !ok {
		resp.Diagnostics.AddAttributeError(path.Root("device"), "Invalid device", fmt.Sprintf("Device '%s' does not exist in provider configuration.", state.Device.ValueString()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.ValueString()))

	res, err := r.clients[state.Device.ValueString()].GetData(state.Id.ValueString())
	if res.StatusCode == 404 {
		state = EVPNInstance{Device: state.Device, Id: state.Id}
	} else {
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		}

		state.updateFromBody(ctx, res.Res)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *EVPNInstanceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state EVPNInstance

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Read state
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if _, ok := r.clients[plan.Device.ValueString()]; !ok {
		resp.Diagnostics.AddAttributeError(path.Root("device"), "Invalid device", fmt.Sprintf("Device '%s' does not exist in provider configuration.", plan.Device.ValueString()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Id.ValueString()))

	body := plan.toBody(ctx)

	deletedListItems := plan.getDeletedListItems(ctx, state)
	tflog.Debug(ctx, fmt.Sprintf("List items to delete: %+v", deletedListItems))

	emptyLeafsDelete := plan.getEmptyLeafsDelete(ctx)
	tflog.Debug(ctx, fmt.Sprintf("List of empty leafs to delete: %+v", emptyLeafsDelete))

	if YangPatch {
		edits := []restconf.YangPatchEdit{restconf.NewYangPatchEdit("merge", plan.getPath(), restconf.Body{Str: body})}
		for _, i := range deletedListItems {
			edits = append(edits, restconf.NewYangPatchEdit("remove", i, restconf.Body{}))
		}
		for _, i := range emptyLeafsDelete {
			edits = append(edits, restconf.NewYangPatchEdit("remove", i, restconf.Body{}))
		}
		_, err := r.clients[plan.Device.ValueString()].YangPatchData("", "1", "", edits)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object, got error: %s", err))
			return
		}
	} else {
		res, err := r.clients[plan.Device.ValueString()].PatchData(plan.getPathShort(), body)
		if len(res.Errors.Error) > 0 && res.Errors.Error[0].ErrorMessage == "patch to a nonexistent resource" {
			_, err = r.clients[plan.Device.ValueString()].PutData(plan.getPath(), body)
		}
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PATCH), got error: %s", err))
			return
		}
		for _, i := range deletedListItems {
			res, err := r.clients[state.Device.ValueString()].DeleteData(i)
			if err != nil && res.StatusCode != 404 {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object, got error: %s", err))
				return
			}
		}
		for _, i := range emptyLeafsDelete {
			res, err := r.clients[plan.Device.ValueString()].DeleteData(i)
			if err != nil && res.StatusCode != 404 {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object, got error: %s", err))
				return
			}
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *EVPNInstanceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state EVPNInstance

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if _, ok := r.clients[state.Device.ValueString()]; !ok {
		resp.Diagnostics.AddAttributeError(path.Root("device"), "Invalid device", fmt.Sprintf("Device '%s' does not exist in provider configuration.", state.Device.ValueString()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	deleteMode := "all"

	if deleteMode == "all" {
		res, err := r.clients[state.Device.ValueString()].DeleteData(state.Id.ValueString())
		if err != nil && res.StatusCode != 404 && res.StatusCode != 400 {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object, got error: %s", err))
			return
		}
	} else {
		deletePaths := state.getDeletePaths(ctx)
		tflog.Debug(ctx, fmt.Sprintf("Paths to delete: %+v", deletePaths))

		if YangPatch {
			edits := []restconf.YangPatchEdit{}
			for _, i := range deletePaths {
				edits = append(edits, restconf.NewYangPatchEdit("remove", i, restconf.Body{}))
			}
			_, err := r.clients[state.Device.ValueString()].YangPatchData("", "1", "", edits)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object, got error: %s", err))
				return
			}
		} else {
			for _, i := range deletePaths {
				res, err := r.clients[state.Device.ValueString()].DeleteData(i)
				if err != nil && res.StatusCode != 404 {
					resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object, got error: %s", err))
					return
				}
			}
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

func (r *EVPNInstanceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}