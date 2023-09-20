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

func TestAccIosxeAAA(t *testing.T) {
	if os.Getenv("AAA") == "" {
		t.Skip("skipping test, set environment variable AAA")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_aaa.test", "new_model", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_aaa.test", "server_radius_dynamic_author", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_aaa.test", "session_id", "common"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_aaa.test", "server_radius_dynamic_author_clients.0.ip", "11.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_aaa.test", "server_radius_dynamic_author_clients.0.server_key_type", "0"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_aaa.test", "server_radius_dynamic_author_clients.0.server_key", "abcd123"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_aaa.test", "group_server_radius.0.name", "T-Group"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_aaa.test", "group_server_radius.0.server_names.0.name", "TESTRADIUS"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_aaa.test", "group_server_radius.0.ip_radius_source_interface_loopback", "0"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_aaa.test", "group_tacacsplus.0.name", "tacacs-group"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_aaa.test", "group_tacacsplus.0.servers.0.name", "tacacs_10.10.15.12"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeAAAConfig_minimum(),
			},
			{
				Config: testAccIosxeAAAConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				ResourceName:  "iosxe_aaa.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XE-native:native/aaa",
			},
		},
	})
}

func testAccIosxeAAAConfig_minimum() string {
	config := `resource "iosxe_aaa" "test" {` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxeAAAConfig_all() string {
	config := `resource "iosxe_aaa" "test" {` + "\n"
	config += `	new_model = true` + "\n"
	config += `	server_radius_dynamic_author = true` + "\n"
	config += `	session_id = "common"` + "\n"
	config += `	server_radius_dynamic_author_clients = [{` + "\n"
	config += `		ip = "11.1.1.1"` + "\n"
	config += `		server_key_type = "0"` + "\n"
	config += `		server_key = "abcd123"` + "\n"
	config += `	}]` + "\n"
	config += `	group_server_radius = [{` + "\n"
	config += `		name = "T-Group"` + "\n"
	config += `		server_names = [{` + "\n"
	config += `			name = "TESTRADIUS"` + "\n"
	config += `		}]` + "\n"
	config += `		ip_radius_source_interface_loopback = 0` + "\n"
	config += `	}]` + "\n"
	config += `	group_tacacsplus = [{` + "\n"
	config += `		name = "tacacs-group"` + "\n"
	config += `		servers = [{` + "\n"
	config += `			name = "tacacs_10.10.15.12"` + "\n"
	config += `		}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}
