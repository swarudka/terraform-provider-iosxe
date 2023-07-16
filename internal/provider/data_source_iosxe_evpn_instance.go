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
	_ datasource.DataSource              = &EVPNInstanceDataSource{}
	_ datasource.DataSourceWithConfigure = &EVPNInstanceDataSource{}
)

func NewEVPNInstanceDataSource() datasource.DataSource {
	return &EVPNInstanceDataSource{}
}

type EVPNInstanceDataSource struct {
	clients map[string]*restconf.Client
}

func (d *EVPNInstanceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_evpn_instance"
}

func (d *EVPNInstanceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the EVPN Instance configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the retrieved object.",
				Computed:            true,
			},
			"evpn_instance_num": schema.Int64Attribute{
				MarkdownDescription: "evpn instance number",
				Required:            true,
			},
			"vlan_based_replication_type_ingress": schema.BoolAttribute{
				MarkdownDescription: "Ingress replication",
				Computed:            true,
			},
			"vlan_based_replication_type_static": schema.BoolAttribute{
				MarkdownDescription: "Static replication",
				Computed:            true,
			},
			"vlan_based_replication_type_p2mp": schema.BoolAttribute{
				MarkdownDescription: "p2mp replication",
				Computed:            true,
			},
			"vlan_based_replication_type_mp2mp": schema.BoolAttribute{
				MarkdownDescription: "mp2mp replication",
				Computed:            true,
			},
			"vlan_based_encapsulation": schema.StringAttribute{
				MarkdownDescription: "Data encapsulation method",
				Computed:            true,
			},
			"vlan_based_auto_route_target": schema.BoolAttribute{
				MarkdownDescription: "Automatically set a route-target",
				Computed:            true,
			},
			"vlan_based_rd": schema.StringAttribute{
				MarkdownDescription: "ASN:nn or IP-address:nn",
				Computed:            true,
			},
			"vlan_based_route_target": schema.StringAttribute{
				MarkdownDescription: "ASN:nn or IP-address:nn",
				Computed:            true,
			},
			"vlan_based_route_target_both": schema.StringAttribute{
				MarkdownDescription: "ASN:nn or IP-address:nn",
				Computed:            true,
			},
			"vlan_based_route_target_import": schema.StringAttribute{
				MarkdownDescription: "ASN:nn or IP-address:nn (DEPRECATED, use rt-value-entry)",
				Computed:            true,
			},
			"vlan_based_route_target_export": schema.StringAttribute{
				MarkdownDescription: "ASN:nn or IP-address:nn (DEPRECATED, use rt-value-entry)",
				Computed:            true,
			},
			"vlan_based_ip_local_learning_disable": schema.BoolAttribute{
				MarkdownDescription: "Disable IP local learning from dataplane",
				Computed:            true,
			},
			"vlan_based_ip_local_learning_enable": schema.BoolAttribute{
				MarkdownDescription: "Enable IP local learning from dataplane",
				Computed:            true,
			},
			"vlan_based_default_gateway_advertise": schema.StringAttribute{
				MarkdownDescription: "Advertise Default Gateway MAC/IP routes",
				Computed:            true,
			},
			"vlan_based_re_originate_route_type5": schema.BoolAttribute{
				MarkdownDescription: "Re-originate route-type 5",
				Computed:            true,
			},
		},
	}
}

func (d *EVPNInstanceDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.clients = req.ProviderData.(map[string]*restconf.Client)
}

func (d *EVPNInstanceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config EVPNInstanceData

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
		config = EVPNInstanceData{Device: config.Device}
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