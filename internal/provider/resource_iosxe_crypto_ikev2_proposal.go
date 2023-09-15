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
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-restconf"
)

func NewCryptoIKEv2ProposalResource() resource.Resource {
	return &CryptoIKEv2ProposalResource{}
}

type CryptoIKEv2ProposalResource struct {
	clients map[string]*restconf.Client
}

func (r *CryptoIKEv2ProposalResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_crypto_ikev2_proposal"
}

func (r *CryptoIKEv2ProposalResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This resource can manage the Crypto IKEv2 Proposal configuration.",

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
			"name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"encryption_en_3des": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("3DES").String,
				Optional:            true,
			},
			"encryption_aes_cbc_128": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("AES-CBC-128").String,
				Optional:            true,
			},
			"encryption_aes_cbc_192": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("AES-CBC-192").String,
				Optional:            true,
			},
			"encryption_aes_cbc_256": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("AES-CBC-256").String,
				Optional:            true,
			},
			"encryption_aes_gcm_128": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Combined-mode,128 bit key,16 byte ICV(Authentication Tag)").String,
				Optional:            true,
			},
			"encryption_aes_gcm_256": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Combined-mode,256 bit key,16 byte ICV(Authentication Tag)").String,
				Optional:            true,
			},
			"group_one": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("DH 768 MODP").String,
				Optional:            true,
			},
			"group_two": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("DH 1024 MODP").String,
				Optional:            true,
			},
			"group_fourteen": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("DH 2048 MODP").String,
				Optional:            true,
			},
			"group_fifteen": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("DH 3072 MODP").String,
				Optional:            true,
			},
			"group_sixteen": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("DH 4096 MODP").String,
				Optional:            true,
			},
			"group_nineteen": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("DH 256 ECP").String,
				Optional:            true,
			},
			"group_twenty": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("DH 384 ECP").String,
				Optional:            true,
			},
			"group_twenty_one": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("DH 521 ECP").String,
				Optional:            true,
			},
			"group_twenty_four": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("DH 2048 (256 subgroup) MODP").String,
				Optional:            true,
			},
			"integrity_md5": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Message Digest 5").String,
				Optional:            true,
			},
			"integrity_sha1": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Secure Hash Standard").String,
				Optional:            true,
			},
			"integrity_sha256": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Secure Hash Standard 2 (256 bit)").String,
				Optional:            true,
			},
			"integrity_sha384": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Secure Hash Standard 2 (384 bit)").String,
				Optional:            true,
			},
			"integrity_sha512": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Secure Hash Standard 2 (512 bit)").String,
				Optional:            true,
			},
			"prf_md5": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Message Digest 5").String,
				Optional:            true,
			},
			"prf_sha1": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Secure Hash Standard").String,
				Optional:            true,
			},
			"prf_sha256": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Secure Hash Standard 2 (256 bit)").String,
				Optional:            true,
			},
			"prf_sha384": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Secure Hash Standard 2 (384 bit)").String,
				Optional:            true,
			},
			"prf_sha512": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Secure Hash Standard 2 (512 bit)").String,
				Optional:            true,
			},
		},
	}
}

func (r *CryptoIKEv2ProposalResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.clients = req.ProviderData.(map[string]*restconf.Client)
}

func (r *CryptoIKEv2ProposalResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan CryptoIKEv2Proposal

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

func (r *CryptoIKEv2ProposalResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state CryptoIKEv2Proposal

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
		state = CryptoIKEv2Proposal{Device: state.Device, Id: state.Id}
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

func (r *CryptoIKEv2ProposalResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state CryptoIKEv2Proposal

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

	deletedItems := plan.getDeletedItems(ctx, state)
	tflog.Debug(ctx, fmt.Sprintf("Removed items to delete: %+v", deletedItems))

	emptyLeafsDelete := plan.getEmptyLeafsDelete(ctx)
	tflog.Debug(ctx, fmt.Sprintf("List of empty leafs to delete: %+v", emptyLeafsDelete))

	if YangPatch {
		edits := []restconf.YangPatchEdit{restconf.NewYangPatchEdit("merge", plan.getPath(), restconf.Body{Str: body})}
		for _, i := range deletedItems {
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
		for _, i := range deletedItems {
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

func (r *CryptoIKEv2ProposalResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state CryptoIKEv2Proposal

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

func (r *CryptoIKEv2ProposalResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
