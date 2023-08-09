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

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-restconf"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &CryptoIKEv2ProfileDataSource{}
	_ datasource.DataSourceWithConfigure = &CryptoIKEv2ProfileDataSource{}
)

func NewCryptoIKEv2ProfileDataSource() datasource.DataSource {
	return &CryptoIKEv2ProfileDataSource{}
}

type CryptoIKEv2ProfileDataSource struct {
	clients map[string]*restconf.Client
}

func (d *CryptoIKEv2ProfileDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_crypto_ikev2_profile"
}

func (d *CryptoIKEv2ProfileDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Crypto IKEv2 Profile configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the retrieved object.",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Specify a description of this profile",
				Computed:            true,
			},
			"authentication_remote_pre_share": schema.BoolAttribute{
				MarkdownDescription: "Pre-Shared Key",
				Computed:            true,
			},
			"authentication_local_pre_share": schema.BoolAttribute{
				MarkdownDescription: "Pre-Shared Key",
				Computed:            true,
			},
			"identity_local_address": schema.StringAttribute{
				MarkdownDescription: "address",
				Computed:            true,
			},
			"identity_local_key_id": schema.StringAttribute{
				MarkdownDescription: "key-id opaque string - proprietary types of identification key-id string",
				Computed:            true,
			},
			"match_inbound_only": schema.BoolAttribute{
				MarkdownDescription: "Match the profile for incoming connections only",
				Computed:            true,
			},
			"match_address_local_ip": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"match_fvrf": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"match_fvrf_any": schema.BoolAttribute{
				MarkdownDescription: "Any fvrf",
				Computed:            true,
			},
			"match_identity_remote_ipv4_addresses": schema.ListNestedAttribute{
				MarkdownDescription: "",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"mask": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
					},
				},
			},
			"match_identity_remote_ipv6_prefixes": schema.ListAttribute{
				MarkdownDescription: "",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"match_identity_remote_keys": schema.ListAttribute{
				MarkdownDescription: "key-id opaque string",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"keyring_local": schema.StringAttribute{
				MarkdownDescription: "Keyring name",
				Computed:            true,
			},
			"dpd_interval": schema.Int64Attribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"dpd_retry": schema.Int64Attribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"dpd_query": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"config_exchange_request": schema.BoolAttribute{
				MarkdownDescription: "enable config-exchange request",
				Computed:            true,
			},
		},
	}
}

func (d *CryptoIKEv2ProfileDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.clients = req.ProviderData.(map[string]*restconf.Client)
}

func (d *CryptoIKEv2ProfileDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config CryptoIKEv2ProfileData

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if _, ok := d.clients[config.Device.ValueString()]; !ok {
		resp.Diagnostics.AddAttributeError(path.Root("device"), "Invalid device", fmt.Sprintf("Device '%s' does not exist in provider configuration.", config.Device.ValueString()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.getPath()))

	res, err := d.clients[config.Device.ValueString()].GetData(config.getPath())
	if res.StatusCode == 404 {
		config = CryptoIKEv2ProfileData{Device: config.Device}
	} else {
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		}

		config.fromBody(ctx, res.Res)
	}

	config.Id = types.StringValue(config.getPath())

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.getPath()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}