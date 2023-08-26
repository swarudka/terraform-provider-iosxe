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
	_ datasource.DataSource              = &InterfaceLoopbackDataSource{}
	_ datasource.DataSourceWithConfigure = &InterfaceLoopbackDataSource{}
)

func NewInterfaceLoopbackDataSource() datasource.DataSource {
	return &InterfaceLoopbackDataSource{}
}

type InterfaceLoopbackDataSource struct {
	clients map[string]*restconf.Client
}

func (d *InterfaceLoopbackDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_loopback"
}

func (d *InterfaceLoopbackDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Interface Loopback configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the retrieved object.",
				Computed:            true,
			},
			"name": schema.Int64Attribute{
				MarkdownDescription: "",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Interface specific description",
				Computed:            true,
			},
			"shutdown": schema.BoolAttribute{
				MarkdownDescription: "Shutdown the selected interface",
				Computed:            true,
			},
			"ip_proxy_arp": schema.BoolAttribute{
				MarkdownDescription: "Enable proxy ARP",
				Computed:            true,
			},
			"ip_redirects": schema.BoolAttribute{
				MarkdownDescription: "Enable sending ICMP Redirect messages",
				Computed:            true,
			},
			"unreachables": schema.BoolAttribute{
				MarkdownDescription: "Enable sending ICMP Unreachable messages",
				Computed:            true,
			},
			"vrf_forwarding": schema.StringAttribute{
				MarkdownDescription: "Configure forwarding table",
				Computed:            true,
			},
			"ipv4_address": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"ipv4_address_mask": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"ip_access_group_in": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"ip_access_group_in_enable": schema.BoolAttribute{
				MarkdownDescription: "inbound packets",
				Computed:            true,
			},
			"ip_access_group_out": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"ip_access_group_out_enable": schema.BoolAttribute{
				MarkdownDescription: "outbound packets",
				Computed:            true,
			},
			"arp_timeout": schema.Int64Attribute{
				MarkdownDescription: "Set ARP cache timeout",
				Computed:            true,
			},
			"ip_arp_inspection_trust": schema.BoolAttribute{
				MarkdownDescription: "Configure Trust state",
				Computed:            true,
			},
			"ip_arp_inspection_limit_rate": schema.Int64Attribute{
				MarkdownDescription: "Rate Limit",
				Computed:            true,
			},
			"snooping_trust": schema.BoolAttribute{
				MarkdownDescription: "DHCP Snooping trust config",
				Computed:            true,
			},
		},
	}
}

func (d *InterfaceLoopbackDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.clients = req.ProviderData.(map[string]*restconf.Client)
}

func (d *InterfaceLoopbackDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config InterfaceLoopbackData

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
		config = InterfaceLoopbackData{Device: config.Device}
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
