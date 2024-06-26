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

func TestAccDataSourceIosxeCryptoIKEv2Proposal(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_crypto_ikev2_proposal.test", "encryption_aes_gcm_256", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_crypto_ikev2_proposal.test", "group_one", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_crypto_ikev2_proposal.test", "group_two", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_crypto_ikev2_proposal.test", "group_fourteen", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_crypto_ikev2_proposal.test", "group_fifteen", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_crypto_ikev2_proposal.test", "group_sixteen", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_crypto_ikev2_proposal.test", "group_nineteen", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_crypto_ikev2_proposal.test", "group_twenty", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_crypto_ikev2_proposal.test", "group_twenty_one", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_crypto_ikev2_proposal.test", "group_twenty_four", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_crypto_ikev2_proposal.test", "integrity_sha1", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_crypto_ikev2_proposal.test", "prf_sha1", "true"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxeCryptoIKEv2ProposalConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

func testAccDataSourceIosxeCryptoIKEv2ProposalConfig() string {
	config := `resource "iosxe_crypto_ikev2_proposal" "test" {` + "\n"
	config += `	name = "PROPOSAL1"` + "\n"
	config += `	encryption_aes_gcm_256 = true` + "\n"
	config += `	group_one = true` + "\n"
	config += `	group_two = true` + "\n"
	config += `	group_fourteen = true` + "\n"
	config += `	group_fifteen = true` + "\n"
	config += `	group_sixteen = true` + "\n"
	config += `	group_nineteen = true` + "\n"
	config += `	group_twenty = true` + "\n"
	config += `	group_twenty_one = true` + "\n"
	config += `	group_twenty_four = true` + "\n"
	config += `	integrity_sha1 = true` + "\n"
	config += `	prf_sha1 = true` + "\n"
	config += `}` + "\n"

	config += `
		data "iosxe_crypto_ikev2_proposal" "test" {
			name = "PROPOSAL1"
			depends_on = [iosxe_crypto_ikev2_proposal.test]
		}
	`
	return config
}
