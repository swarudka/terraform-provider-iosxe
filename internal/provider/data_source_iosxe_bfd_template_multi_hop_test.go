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

func TestAccDataSourceIosxeBFDTemplateMultiHop(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_bfd_template_multi_hop.test", "echo", "true"))
	if os.Getenv("IOSXE179") != "" || os.Getenv("IOSXE1710") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_bfd_template_multi_hop.test", "interval_milliseconds_min_tx", "4500"))
	}
	if os.Getenv("IOSXE179") != "" || os.Getenv("IOSXE1710") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_bfd_template_multi_hop.test", "interval_milliseconds_min_rx", "5500"))
	}
	if os.Getenv("IOSXE179") != "" || os.Getenv("IOSXE1710") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_bfd_template_multi_hop.test", "interval_milliseconds_multiplier", "40"))
	}
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_bfd_template_multi_hop.test", "authentication_md5_keychain", "KEYNAME"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_bfd_template_multi_hop.test", "dampening_half_time", "21"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_bfd_template_multi_hop.test", "dampening_unsuppress_time", "1800"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_bfd_template_multi_hop.test", "dampening_suppress_time", "1900"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_bfd_template_multi_hop.test", "dampening_max_suppressing_time", "70"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxeBFDTemplateMultiHopConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

func testAccDataSourceIosxeBFDTemplateMultiHopConfig() string {
	config := `resource "iosxe_bfd_template_multi_hop" "test" {` + "\n"
	config += `	name = "T11"` + "\n"
	config += `	echo = true` + "\n"
	if os.Getenv("IOSXE179") != "" || os.Getenv("IOSXE1710") != "" {
		config += `	interval_milliseconds_min_tx = 4500` + "\n"
	}
	if os.Getenv("IOSXE179") != "" || os.Getenv("IOSXE1710") != "" {
		config += `	interval_milliseconds_min_rx = 5500` + "\n"
	}
	if os.Getenv("IOSXE179") != "" || os.Getenv("IOSXE1710") != "" {
		config += `	interval_milliseconds_multiplier = 40` + "\n"
	}
	config += `	authentication_md5_keychain = "KEYNAME"` + "\n"
	config += `	dampening_half_time = 21` + "\n"
	config += `	dampening_unsuppress_time = 1800` + "\n"
	config += `	dampening_suppress_time = 1900` + "\n"
	config += `	dampening_max_suppressing_time = 70` + "\n"
	config += `}` + "\n"

	config += `
		data "iosxe_bfd_template_multi_hop" "test" {
			name = "T11"
			depends_on = [iosxe_bfd_template_multi_hop.test]
		}
	`
	return config
}
