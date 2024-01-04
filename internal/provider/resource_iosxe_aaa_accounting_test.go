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

func TestAccIosxeAAAAccounting(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_aaa_accounting.test", "update_newinfo_periodic", "2880"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_aaa_accounting.test", "identities.0.name", "test"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_aaa_accounting.test", "identities.0.start_stop_broadcast", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_aaa_accounting.test", "identities.0.start_stop_group_broadcast", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_aaa_accounting.test", "identities.0.start_stop_group_logger", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_aaa_accounting.test", "identities.0.start_stop_group1", "GROUP1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_aaa_accounting.test", "identities.0.start_stop_group2", "GROUP2"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_aaa_accounting.test", "identities.0.start_stop_group3", "GROUP3"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_aaa_accounting.test", "identities.0.start_stop_group4", "GROUP4"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_aaa_accounting.test", "identity_default_start_stop_group1", "RADIUS-GROUP"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_aaa_accounting.test", "identity_default_start_stop_group2", "RADIUS-GROUP2"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_aaa_accounting.test", "identity_default_start_stop_group3", "RADIUS-GROUP3"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_aaa_accounting.test", "identity_default_start_stop_group4", "RADIUS-GROUP4"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_aaa_accounting.test", "execs.0.name", "default"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_aaa_accounting.test", "execs.0.start_stop_group1", "T-Group"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_aaa_accounting.test", "networks.0.id", "network1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_aaa_accounting.test", "networks.0.start_stop_group1", "radius"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_aaa_accounting.test", "networks.0.start_stop_group2", "tacacs+"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_aaa_accounting.test", "system_guarantee_first", "false"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeAAAAccountingConfig_minimum(),
			},
			{
				Config: testAccIosxeAAAAccountingConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				ResourceName:  "iosxe_aaa_accounting.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XE-native:native/aaa/Cisco-IOS-XE-aaa:accounting",
			},
		},
	})
}

func testAccIosxeAAAAccountingConfig_minimum() string {
	config := `resource "iosxe_aaa_accounting" "test" {` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxeAAAAccountingConfig_all() string {
	config := `resource "iosxe_aaa_accounting" "test" {` + "\n"
	config += `	update_newinfo_periodic = 2880` + "\n"
	config += `	identities = [{` + "\n"
	config += `		name = "test"` + "\n"
	config += `		start_stop_broadcast = false` + "\n"
	config += `		start_stop_group_broadcast = false` + "\n"
	config += `		start_stop_group_logger = false` + "\n"
	config += `		start_stop_group1 = "GROUP1"` + "\n"
	config += `		start_stop_group2 = "GROUP2"` + "\n"
	config += `		start_stop_group3 = "GROUP3"` + "\n"
	config += `		start_stop_group4 = "GROUP4"` + "\n"
	config += `	}]` + "\n"
	config += `	identity_default_start_stop_group1 = "RADIUS-GROUP"` + "\n"
	config += `	identity_default_start_stop_group2 = "RADIUS-GROUP2"` + "\n"
	config += `	identity_default_start_stop_group3 = "RADIUS-GROUP3"` + "\n"
	config += `	identity_default_start_stop_group4 = "RADIUS-GROUP4"` + "\n"
	config += `	execs = [{` + "\n"
	config += `		name = "default"` + "\n"
	config += `		start_stop_group1 = "T-Group"` + "\n"
	config += `	}]` + "\n"
	config += `	networks = [{` + "\n"
	config += `		id = "network1"` + "\n"
	config += `		start_stop_group1 = "radius"` + "\n"
	config += `		start_stop_group2 = "tacacs+"` + "\n"
	config += `	}]` + "\n"
	config += `	system_guarantee_first = false` + "\n"
	config += `}` + "\n"
	return config
}
