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
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxeFlowRecord(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_flow_record.test", "description", "My flow record"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_flow_record.test", "match_ipv4_source_address", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_flow_record.test", "match_ipv4_destination_address", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_flow_record.test", "match_ipv4_protocol", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_flow_record.test", "match_ipv4_tos", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_flow_record.test", "match_transport_source_port", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_flow_record.test", "match_transport_destination_port", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_flow_record.test", "match_interface_input", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_flow_record.test", "match_flow_direction", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_flow_record.test", "collect_interface_output", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_flow_record.test", "collect_counter_bytes_long", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_flow_record.test", "collect_counter_packets_long", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_flow_record.test", "collect_transport_tcp_flags", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_flow_record.test", "collect_timestamp_absolute_first", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_flow_record.test", "collect_timestamp_absolute_last", "true"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxeFlowRecordConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

func testAccDataSourceIosxeFlowRecordConfig() string {
	config := `resource "iosxe_flow_record" "test" {` + "\n"
	config += `	delete_mode = "attributes"` + "\n"
	config += `	name = "FNF1"` + "\n"
	config += `	description = "My flow record"` + "\n"
	config += `	match_ipv4_source_address = true` + "\n"
	config += `	match_ipv4_destination_address = true` + "\n"
	config += `	match_ipv4_protocol = true` + "\n"
	config += `	match_ipv4_tos = true` + "\n"
	config += `	match_transport_source_port = true` + "\n"
	config += `	match_transport_destination_port = true` + "\n"
	config += `	match_interface_input = true` + "\n"
	config += `	match_flow_direction = true` + "\n"
	config += `	collect_interface_output = true` + "\n"
	config += `	collect_counter_bytes_long = true` + "\n"
	config += `	collect_counter_packets_long = true` + "\n"
	config += `	collect_transport_tcp_flags = true` + "\n"
	config += `	collect_timestamp_absolute_first = true` + "\n"
	config += `	collect_timestamp_absolute_last = true` + "\n"
	config += `}` + "\n"

	config += `
		data "iosxe_flow_record" "test" {
			name = "FNF1"
			depends_on = [iosxe_flow_record.test]
		}
	`
	return config
}
