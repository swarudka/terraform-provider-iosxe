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
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxeInterfacePortChannel(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_port_channel.test", "description", "My Interface Description"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_port_channel.test", "shutdown", "false"))
	if os.Getenv("C8000V") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_port_channel.test", "ip_proxy_arp", "false"))
	}
	if os.Getenv("C8000V") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_port_channel.test", "ip_redirects", "false"))
	}
	if os.Getenv("C8000V") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_port_channel.test", "ip_unreachables", "false"))
	}
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_port_channel.test", "vrf_forwarding", "VRF1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_port_channel.test", "ipv4_address", "192.0.2.1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_port_channel.test", "ipv4_address_mask", "255.255.255.0"))
	if os.Getenv("C9000V") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_port_channel.test", "switchport", "false"))
	}
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_port_channel.test", "ip_access_group_in", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_port_channel.test", "ip_access_group_in_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_port_channel.test", "ip_access_group_out", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_port_channel.test", "ip_access_group_out_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_port_channel.test", "ip_dhcp_relay_source_interface", "Loopback100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_port_channel.test", "helper_addresses.0.address", "10.10.10.10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_port_channel.test", "helper_addresses.0.global", "false"))
	if os.Getenv("IOSXE179") != "" || os.Getenv("IOSXE1710") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_port_channel.test", "bfd_template", "bfd_template1"))
	}
	if os.Getenv("IOSXE179") != "" || os.Getenv("IOSXE1710") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_port_channel.test", "bfd_enable", "true"))
	}
	if os.Getenv("IOSXE179") != "" || os.Getenv("IOSXE1710") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_port_channel.test", "bfd_local_address", "1.2.3.4"))
	}
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_port_channel.test", "ipv6_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_port_channel.test", "ipv6_mtu", "1300"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_port_channel.test", "ipv6_nd_ra_suppress_all", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_port_channel.test", "ipv6_address_dhcp", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_port_channel.test", "ipv6_link_local_addresses.0.address", "fe80::64"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_port_channel.test", "ipv6_link_local_addresses.0.link_local", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_port_channel.test", "ipv6_addresses.0.prefix", "2224:DB8::/32"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_port_channel.test", "ipv6_addresses.0.eui_64", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_port_channel.test", "arp_timeout", "2147"))
	if os.Getenv("C9000V") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_port_channel.test", "ip_arp_inspection_trust", "true"))
	}
	if os.Getenv("C9000V") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_port_channel.test", "ip_arp_inspection_limit_rate", "1000"))
	}
	if os.Getenv("C9000V") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_port_channel.test", "spanning_tree_link_type", "point-to-point"))
	}
	if os.Getenv("C9000V") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_port_channel.test", "ip_dhcp_snooping_trust", "true"))
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxeInterfacePortChannelPrerequisitesConfig + testAccDataSourceIosxeInterfacePortChannelConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

const testAccDataSourceIosxeInterfacePortChannelPrerequisitesConfig = `
resource "iosxe_restconf" "PreReq0" {
	path = "Cisco-IOS-XE-native:native/vrf/definition=VRF1"
	delete = false
	attributes = {
		"name" = "VRF1"
		"address-family/ipv4" = ""
	}
}

`

func testAccDataSourceIosxeInterfacePortChannelConfig() string {
	config := `resource "iosxe_interface_port_channel" "test" {` + "\n"
	config += `	delete_mode = "attributes"` + "\n"
	config += `	name = 10` + "\n"
	config += `	description = "My Interface Description"` + "\n"
	config += `	shutdown = false` + "\n"
	if os.Getenv("C8000V") != "" {
		config += `	ip_proxy_arp = false` + "\n"
	}
	if os.Getenv("C8000V") != "" {
		config += `	ip_redirects = false` + "\n"
	}
	if os.Getenv("C8000V") != "" {
		config += `	ip_unreachables = false` + "\n"
	}
	config += `	vrf_forwarding = "VRF1"` + "\n"
	config += `	ipv4_address = "192.0.2.1"` + "\n"
	config += `	ipv4_address_mask = "255.255.255.0"` + "\n"
	if os.Getenv("C9000V") != "" {
		config += `	switchport = false` + "\n"
	}
	config += `	ip_access_group_in = "1"` + "\n"
	config += `	ip_access_group_in_enable = true` + "\n"
	config += `	ip_access_group_out = "1"` + "\n"
	config += `	ip_access_group_out_enable = true` + "\n"
	config += `	ip_dhcp_relay_source_interface = "Loopback100"` + "\n"
	config += `	helper_addresses = [{` + "\n"
	config += `		address = "10.10.10.10"` + "\n"
	config += `		global = false` + "\n"
	config += `	}]` + "\n"
	if os.Getenv("IOSXE179") != "" || os.Getenv("IOSXE1710") != "" {
		config += `	bfd_template = "bfd_template1"` + "\n"
	}
	if os.Getenv("IOSXE179") != "" || os.Getenv("IOSXE1710") != "" {
		config += `	bfd_enable = true` + "\n"
	}
	if os.Getenv("IOSXE179") != "" || os.Getenv("IOSXE1710") != "" {
		config += `	bfd_local_address = "1.2.3.4"` + "\n"
	}
	config += `	ipv6_enable = true` + "\n"
	config += `	ipv6_mtu = 1300` + "\n"
	config += `	ipv6_nd_ra_suppress_all = true` + "\n"
	config += `	ipv6_address_dhcp = true` + "\n"
	config += `	ipv6_link_local_addresses = [{` + "\n"
	config += `		address = "fe80::64"` + "\n"
	config += `		link_local = true` + "\n"
	config += `	}]` + "\n"
	config += `	ipv6_addresses = [{` + "\n"
	config += `		prefix = "2224:DB8::/32"` + "\n"
	config += `		eui_64 = true` + "\n"
	config += `	}]` + "\n"
	config += `	arp_timeout = 2147` + "\n"
	if os.Getenv("C9000V") != "" {
		config += `	ip_arp_inspection_trust = true` + "\n"
	}
	if os.Getenv("C9000V") != "" {
		config += `	ip_arp_inspection_limit_rate = 1000` + "\n"
	}
	if os.Getenv("C9000V") != "" {
		config += `	spanning_tree_link_type = "point-to-point"` + "\n"
	}
	if os.Getenv("C9000V") != "" {
		config += `	ip_dhcp_snooping_trust = true` + "\n"
	}
	config += `	depends_on = [iosxe_restconf.PreReq0, ]` + "\n"
	config += `}` + "\n"

	config += `
		data "iosxe_interface_port_channel" "test" {
			name = 10
			depends_on = [iosxe_interface_port_channel.test]
		}
	`
	return config
}
