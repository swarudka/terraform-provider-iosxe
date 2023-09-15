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
	"net/url"
	"regexp"

	"github.com/CiscoDevNet/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type CryptoIKEv2Proposal struct {
	Device              types.String `tfsdk:"device"`
	Id                  types.String `tfsdk:"id"`
	Name                types.String `tfsdk:"name"`
	EncryptionEn3des    types.Bool   `tfsdk:"encryption_en_3des"`
	EncryptionAesCbc128 types.Bool   `tfsdk:"encryption_aes_cbc_128"`
	EncryptionAesCbc192 types.Bool   `tfsdk:"encryption_aes_cbc_192"`
	EncryptionAesCbc256 types.Bool   `tfsdk:"encryption_aes_cbc_256"`
	EncryptionAesGcm128 types.Bool   `tfsdk:"encryption_aes_gcm_128"`
	EncryptionAesGcm256 types.Bool   `tfsdk:"encryption_aes_gcm_256"`
	GroupOne            types.Bool   `tfsdk:"group_one"`
	GroupTwo            types.Bool   `tfsdk:"group_two"`
	GroupFourteen       types.Bool   `tfsdk:"group_fourteen"`
	GroupFifteen        types.Bool   `tfsdk:"group_fifteen"`
	GroupSixteen        types.Bool   `tfsdk:"group_sixteen"`
	GroupNineteen       types.Bool   `tfsdk:"group_nineteen"`
	GroupTwenty         types.Bool   `tfsdk:"group_twenty"`
	GroupTwentyOne      types.Bool   `tfsdk:"group_twenty_one"`
	GroupTwentyFour     types.Bool   `tfsdk:"group_twenty_four"`
	IntegrityMd5        types.Bool   `tfsdk:"integrity_md5"`
	IntegritySha1       types.Bool   `tfsdk:"integrity_sha1"`
	IntegritySha256     types.Bool   `tfsdk:"integrity_sha256"`
	IntegritySha384     types.Bool   `tfsdk:"integrity_sha384"`
	IntegritySha512     types.Bool   `tfsdk:"integrity_sha512"`
	PrfMd5              types.Bool   `tfsdk:"prf_md5"`
	PrfSha1             types.Bool   `tfsdk:"prf_sha1"`
	PrfSha256           types.Bool   `tfsdk:"prf_sha256"`
	PrfSha384           types.Bool   `tfsdk:"prf_sha384"`
	PrfSha512           types.Bool   `tfsdk:"prf_sha512"`
}

type CryptoIKEv2ProposalData struct {
	Device              types.String `tfsdk:"device"`
	Id                  types.String `tfsdk:"id"`
	Name                types.String `tfsdk:"name"`
	EncryptionEn3des    types.Bool   `tfsdk:"encryption_en_3des"`
	EncryptionAesCbc128 types.Bool   `tfsdk:"encryption_aes_cbc_128"`
	EncryptionAesCbc192 types.Bool   `tfsdk:"encryption_aes_cbc_192"`
	EncryptionAesCbc256 types.Bool   `tfsdk:"encryption_aes_cbc_256"`
	EncryptionAesGcm128 types.Bool   `tfsdk:"encryption_aes_gcm_128"`
	EncryptionAesGcm256 types.Bool   `tfsdk:"encryption_aes_gcm_256"`
	GroupOne            types.Bool   `tfsdk:"group_one"`
	GroupTwo            types.Bool   `tfsdk:"group_two"`
	GroupFourteen       types.Bool   `tfsdk:"group_fourteen"`
	GroupFifteen        types.Bool   `tfsdk:"group_fifteen"`
	GroupSixteen        types.Bool   `tfsdk:"group_sixteen"`
	GroupNineteen       types.Bool   `tfsdk:"group_nineteen"`
	GroupTwenty         types.Bool   `tfsdk:"group_twenty"`
	GroupTwentyOne      types.Bool   `tfsdk:"group_twenty_one"`
	GroupTwentyFour     types.Bool   `tfsdk:"group_twenty_four"`
	IntegrityMd5        types.Bool   `tfsdk:"integrity_md5"`
	IntegritySha1       types.Bool   `tfsdk:"integrity_sha1"`
	IntegritySha256     types.Bool   `tfsdk:"integrity_sha256"`
	IntegritySha384     types.Bool   `tfsdk:"integrity_sha384"`
	IntegritySha512     types.Bool   `tfsdk:"integrity_sha512"`
	PrfMd5              types.Bool   `tfsdk:"prf_md5"`
	PrfSha1             types.Bool   `tfsdk:"prf_sha1"`
	PrfSha256           types.Bool   `tfsdk:"prf_sha256"`
	PrfSha384           types.Bool   `tfsdk:"prf_sha384"`
	PrfSha512           types.Bool   `tfsdk:"prf_sha512"`
}

func (data CryptoIKEv2Proposal) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/crypto/Cisco-IOS-XE-crypto:ikev2/proposal=%v", url.QueryEscape(fmt.Sprintf("%v", data.Name.ValueString())))
}

func (data CryptoIKEv2ProposalData) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/crypto/Cisco-IOS-XE-crypto:ikev2/proposal=%v", url.QueryEscape(fmt.Sprintf("%v", data.Name.ValueString())))
}

// if last path element has a key -> remove it
func (data CryptoIKEv2Proposal) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data CryptoIKEv2Proposal) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"name", data.Name.ValueString())
	}
	if !data.EncryptionEn3des.IsNull() && !data.EncryptionEn3des.IsUnknown() {
		if data.EncryptionEn3des.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"encryption.en-3des", map[string]string{})
		}
	}
	if !data.EncryptionAesCbc128.IsNull() && !data.EncryptionAesCbc128.IsUnknown() {
		if data.EncryptionAesCbc128.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"encryption.aes-cbc-128", map[string]string{})
		}
	}
	if !data.EncryptionAesCbc192.IsNull() && !data.EncryptionAesCbc192.IsUnknown() {
		if data.EncryptionAesCbc192.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"encryption.aes-cbc-192", map[string]string{})
		}
	}
	if !data.EncryptionAesCbc256.IsNull() && !data.EncryptionAesCbc256.IsUnknown() {
		if data.EncryptionAesCbc256.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"encryption.aes-cbc-256", map[string]string{})
		}
	}
	if !data.EncryptionAesGcm128.IsNull() && !data.EncryptionAesGcm128.IsUnknown() {
		if data.EncryptionAesGcm128.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"encryption.aes-gcm-128", map[string]string{})
		}
	}
	if !data.EncryptionAesGcm256.IsNull() && !data.EncryptionAesGcm256.IsUnknown() {
		if data.EncryptionAesGcm256.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"encryption.aes-gcm-256", map[string]string{})
		}
	}
	if !data.GroupOne.IsNull() && !data.GroupOne.IsUnknown() {
		if data.GroupOne.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"group.one", map[string]string{})
		}
	}
	if !data.GroupTwo.IsNull() && !data.GroupTwo.IsUnknown() {
		if data.GroupTwo.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"group.two", map[string]string{})
		}
	}
	if !data.GroupFourteen.IsNull() && !data.GroupFourteen.IsUnknown() {
		if data.GroupFourteen.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"group.fourteen", map[string]string{})
		}
	}
	if !data.GroupFifteen.IsNull() && !data.GroupFifteen.IsUnknown() {
		if data.GroupFifteen.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"group.fifteen", map[string]string{})
		}
	}
	if !data.GroupSixteen.IsNull() && !data.GroupSixteen.IsUnknown() {
		if data.GroupSixteen.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"group.sixteen", map[string]string{})
		}
	}
	if !data.GroupNineteen.IsNull() && !data.GroupNineteen.IsUnknown() {
		if data.GroupNineteen.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"group.nineteen", map[string]string{})
		}
	}
	if !data.GroupTwenty.IsNull() && !data.GroupTwenty.IsUnknown() {
		if data.GroupTwenty.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"group.twenty", map[string]string{})
		}
	}
	if !data.GroupTwentyOne.IsNull() && !data.GroupTwentyOne.IsUnknown() {
		if data.GroupTwentyOne.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"group.twenty-one", map[string]string{})
		}
	}
	if !data.GroupTwentyFour.IsNull() && !data.GroupTwentyFour.IsUnknown() {
		if data.GroupTwentyFour.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"group.twenty-four", map[string]string{})
		}
	}
	if !data.IntegrityMd5.IsNull() && !data.IntegrityMd5.IsUnknown() {
		if data.IntegrityMd5.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"integrity.md5", map[string]string{})
		}
	}
	if !data.IntegritySha1.IsNull() && !data.IntegritySha1.IsUnknown() {
		if data.IntegritySha1.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"integrity.sha1", map[string]string{})
		}
	}
	if !data.IntegritySha256.IsNull() && !data.IntegritySha256.IsUnknown() {
		if data.IntegritySha256.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"integrity.sha256", map[string]string{})
		}
	}
	if !data.IntegritySha384.IsNull() && !data.IntegritySha384.IsUnknown() {
		if data.IntegritySha384.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"integrity.sha384", map[string]string{})
		}
	}
	if !data.IntegritySha512.IsNull() && !data.IntegritySha512.IsUnknown() {
		if data.IntegritySha512.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"integrity.sha512", map[string]string{})
		}
	}
	if !data.PrfMd5.IsNull() && !data.PrfMd5.IsUnknown() {
		if data.PrfMd5.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"prf.md5", map[string]string{})
		}
	}
	if !data.PrfSha1.IsNull() && !data.PrfSha1.IsUnknown() {
		if data.PrfSha1.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"prf.sha1", map[string]string{})
		}
	}
	if !data.PrfSha256.IsNull() && !data.PrfSha256.IsUnknown() {
		if data.PrfSha256.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"prf.sha256", map[string]string{})
		}
	}
	if !data.PrfSha384.IsNull() && !data.PrfSha384.IsUnknown() {
		if data.PrfSha384.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"prf.sha384", map[string]string{})
		}
	}
	if !data.PrfSha512.IsNull() && !data.PrfSha512.IsUnknown() {
		if data.PrfSha512.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"prf.sha512", map[string]string{})
		}
	}
	return body
}

func (data *CryptoIKEv2Proposal) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get(prefix + "encryption.en-3des"); !data.EncryptionEn3des.IsNull() {
		if value.Exists() {
			data.EncryptionEn3des = types.BoolValue(true)
		} else {
			data.EncryptionEn3des = types.BoolValue(false)
		}
	} else {
		data.EncryptionEn3des = types.BoolNull()
	}
	if value := res.Get(prefix + "encryption.aes-cbc-128"); !data.EncryptionAesCbc128.IsNull() {
		if value.Exists() {
			data.EncryptionAesCbc128 = types.BoolValue(true)
		} else {
			data.EncryptionAesCbc128 = types.BoolValue(false)
		}
	} else {
		data.EncryptionAesCbc128 = types.BoolNull()
	}
	if value := res.Get(prefix + "encryption.aes-cbc-192"); !data.EncryptionAesCbc192.IsNull() {
		if value.Exists() {
			data.EncryptionAesCbc192 = types.BoolValue(true)
		} else {
			data.EncryptionAesCbc192 = types.BoolValue(false)
		}
	} else {
		data.EncryptionAesCbc192 = types.BoolNull()
	}
	if value := res.Get(prefix + "encryption.aes-cbc-256"); !data.EncryptionAesCbc256.IsNull() {
		if value.Exists() {
			data.EncryptionAesCbc256 = types.BoolValue(true)
		} else {
			data.EncryptionAesCbc256 = types.BoolValue(false)
		}
	} else {
		data.EncryptionAesCbc256 = types.BoolNull()
	}
	if value := res.Get(prefix + "encryption.aes-gcm-128"); !data.EncryptionAesGcm128.IsNull() {
		if value.Exists() {
			data.EncryptionAesGcm128 = types.BoolValue(true)
		} else {
			data.EncryptionAesGcm128 = types.BoolValue(false)
		}
	} else {
		data.EncryptionAesGcm128 = types.BoolNull()
	}
	if value := res.Get(prefix + "encryption.aes-gcm-256"); !data.EncryptionAesGcm256.IsNull() {
		if value.Exists() {
			data.EncryptionAesGcm256 = types.BoolValue(true)
		} else {
			data.EncryptionAesGcm256 = types.BoolValue(false)
		}
	} else {
		data.EncryptionAesGcm256 = types.BoolNull()
	}
	if value := res.Get(prefix + "group.one"); !data.GroupOne.IsNull() {
		if value.Exists() {
			data.GroupOne = types.BoolValue(true)
		} else {
			data.GroupOne = types.BoolValue(false)
		}
	} else {
		data.GroupOne = types.BoolNull()
	}
	if value := res.Get(prefix + "group.two"); !data.GroupTwo.IsNull() {
		if value.Exists() {
			data.GroupTwo = types.BoolValue(true)
		} else {
			data.GroupTwo = types.BoolValue(false)
		}
	} else {
		data.GroupTwo = types.BoolNull()
	}
	if value := res.Get(prefix + "group.fourteen"); !data.GroupFourteen.IsNull() {
		if value.Exists() {
			data.GroupFourteen = types.BoolValue(true)
		} else {
			data.GroupFourteen = types.BoolValue(false)
		}
	} else {
		data.GroupFourteen = types.BoolNull()
	}
	if value := res.Get(prefix + "group.fifteen"); !data.GroupFifteen.IsNull() {
		if value.Exists() {
			data.GroupFifteen = types.BoolValue(true)
		} else {
			data.GroupFifteen = types.BoolValue(false)
		}
	} else {
		data.GroupFifteen = types.BoolNull()
	}
	if value := res.Get(prefix + "group.sixteen"); !data.GroupSixteen.IsNull() {
		if value.Exists() {
			data.GroupSixteen = types.BoolValue(true)
		} else {
			data.GroupSixteen = types.BoolValue(false)
		}
	} else {
		data.GroupSixteen = types.BoolNull()
	}
	if value := res.Get(prefix + "group.nineteen"); !data.GroupNineteen.IsNull() {
		if value.Exists() {
			data.GroupNineteen = types.BoolValue(true)
		} else {
			data.GroupNineteen = types.BoolValue(false)
		}
	} else {
		data.GroupNineteen = types.BoolNull()
	}
	if value := res.Get(prefix + "group.twenty"); !data.GroupTwenty.IsNull() {
		if value.Exists() {
			data.GroupTwenty = types.BoolValue(true)
		} else {
			data.GroupTwenty = types.BoolValue(false)
		}
	} else {
		data.GroupTwenty = types.BoolNull()
	}
	if value := res.Get(prefix + "group.twenty-one"); !data.GroupTwentyOne.IsNull() {
		if value.Exists() {
			data.GroupTwentyOne = types.BoolValue(true)
		} else {
			data.GroupTwentyOne = types.BoolValue(false)
		}
	} else {
		data.GroupTwentyOne = types.BoolNull()
	}
	if value := res.Get(prefix + "group.twenty-four"); !data.GroupTwentyFour.IsNull() {
		if value.Exists() {
			data.GroupTwentyFour = types.BoolValue(true)
		} else {
			data.GroupTwentyFour = types.BoolValue(false)
		}
	} else {
		data.GroupTwentyFour = types.BoolNull()
	}
	if value := res.Get(prefix + "integrity.md5"); !data.IntegrityMd5.IsNull() {
		if value.Exists() {
			data.IntegrityMd5 = types.BoolValue(true)
		} else {
			data.IntegrityMd5 = types.BoolValue(false)
		}
	} else {
		data.IntegrityMd5 = types.BoolNull()
	}
	if value := res.Get(prefix + "integrity.sha1"); !data.IntegritySha1.IsNull() {
		if value.Exists() {
			data.IntegritySha1 = types.BoolValue(true)
		} else {
			data.IntegritySha1 = types.BoolValue(false)
		}
	} else {
		data.IntegritySha1 = types.BoolNull()
	}
	if value := res.Get(prefix + "integrity.sha256"); !data.IntegritySha256.IsNull() {
		if value.Exists() {
			data.IntegritySha256 = types.BoolValue(true)
		} else {
			data.IntegritySha256 = types.BoolValue(false)
		}
	} else {
		data.IntegritySha256 = types.BoolNull()
	}
	if value := res.Get(prefix + "integrity.sha384"); !data.IntegritySha384.IsNull() {
		if value.Exists() {
			data.IntegritySha384 = types.BoolValue(true)
		} else {
			data.IntegritySha384 = types.BoolValue(false)
		}
	} else {
		data.IntegritySha384 = types.BoolNull()
	}
	if value := res.Get(prefix + "integrity.sha512"); !data.IntegritySha512.IsNull() {
		if value.Exists() {
			data.IntegritySha512 = types.BoolValue(true)
		} else {
			data.IntegritySha512 = types.BoolValue(false)
		}
	} else {
		data.IntegritySha512 = types.BoolNull()
	}
	if value := res.Get(prefix + "prf.md5"); !data.PrfMd5.IsNull() {
		if value.Exists() {
			data.PrfMd5 = types.BoolValue(true)
		} else {
			data.PrfMd5 = types.BoolValue(false)
		}
	} else {
		data.PrfMd5 = types.BoolNull()
	}
	if value := res.Get(prefix + "prf.sha1"); !data.PrfSha1.IsNull() {
		if value.Exists() {
			data.PrfSha1 = types.BoolValue(true)
		} else {
			data.PrfSha1 = types.BoolValue(false)
		}
	} else {
		data.PrfSha1 = types.BoolNull()
	}
	if value := res.Get(prefix + "prf.sha256"); !data.PrfSha256.IsNull() {
		if value.Exists() {
			data.PrfSha256 = types.BoolValue(true)
		} else {
			data.PrfSha256 = types.BoolValue(false)
		}
	} else {
		data.PrfSha256 = types.BoolNull()
	}
	if value := res.Get(prefix + "prf.sha384"); !data.PrfSha384.IsNull() {
		if value.Exists() {
			data.PrfSha384 = types.BoolValue(true)
		} else {
			data.PrfSha384 = types.BoolValue(false)
		}
	} else {
		data.PrfSha384 = types.BoolNull()
	}
	if value := res.Get(prefix + "prf.sha512"); !data.PrfSha512.IsNull() {
		if value.Exists() {
			data.PrfSha512 = types.BoolValue(true)
		} else {
			data.PrfSha512 = types.BoolValue(false)
		}
	} else {
		data.PrfSha512 = types.BoolNull()
	}
}

func (data *CryptoIKEv2ProposalData) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "encryption.en-3des"); value.Exists() {
		data.EncryptionEn3des = types.BoolValue(true)
	} else {
		data.EncryptionEn3des = types.BoolValue(false)
	}
	if value := res.Get(prefix + "encryption.aes-cbc-128"); value.Exists() {
		data.EncryptionAesCbc128 = types.BoolValue(true)
	} else {
		data.EncryptionAesCbc128 = types.BoolValue(false)
	}
	if value := res.Get(prefix + "encryption.aes-cbc-192"); value.Exists() {
		data.EncryptionAesCbc192 = types.BoolValue(true)
	} else {
		data.EncryptionAesCbc192 = types.BoolValue(false)
	}
	if value := res.Get(prefix + "encryption.aes-cbc-256"); value.Exists() {
		data.EncryptionAesCbc256 = types.BoolValue(true)
	} else {
		data.EncryptionAesCbc256 = types.BoolValue(false)
	}
	if value := res.Get(prefix + "encryption.aes-gcm-128"); value.Exists() {
		data.EncryptionAesGcm128 = types.BoolValue(true)
	} else {
		data.EncryptionAesGcm128 = types.BoolValue(false)
	}
	if value := res.Get(prefix + "encryption.aes-gcm-256"); value.Exists() {
		data.EncryptionAesGcm256 = types.BoolValue(true)
	} else {
		data.EncryptionAesGcm256 = types.BoolValue(false)
	}
	if value := res.Get(prefix + "group.one"); value.Exists() {
		data.GroupOne = types.BoolValue(true)
	} else {
		data.GroupOne = types.BoolValue(false)
	}
	if value := res.Get(prefix + "group.two"); value.Exists() {
		data.GroupTwo = types.BoolValue(true)
	} else {
		data.GroupTwo = types.BoolValue(false)
	}
	if value := res.Get(prefix + "group.fourteen"); value.Exists() {
		data.GroupFourteen = types.BoolValue(true)
	} else {
		data.GroupFourteen = types.BoolValue(false)
	}
	if value := res.Get(prefix + "group.fifteen"); value.Exists() {
		data.GroupFifteen = types.BoolValue(true)
	} else {
		data.GroupFifteen = types.BoolValue(false)
	}
	if value := res.Get(prefix + "group.sixteen"); value.Exists() {
		data.GroupSixteen = types.BoolValue(true)
	} else {
		data.GroupSixteen = types.BoolValue(false)
	}
	if value := res.Get(prefix + "group.nineteen"); value.Exists() {
		data.GroupNineteen = types.BoolValue(true)
	} else {
		data.GroupNineteen = types.BoolValue(false)
	}
	if value := res.Get(prefix + "group.twenty"); value.Exists() {
		data.GroupTwenty = types.BoolValue(true)
	} else {
		data.GroupTwenty = types.BoolValue(false)
	}
	if value := res.Get(prefix + "group.twenty-one"); value.Exists() {
		data.GroupTwentyOne = types.BoolValue(true)
	} else {
		data.GroupTwentyOne = types.BoolValue(false)
	}
	if value := res.Get(prefix + "group.twenty-four"); value.Exists() {
		data.GroupTwentyFour = types.BoolValue(true)
	} else {
		data.GroupTwentyFour = types.BoolValue(false)
	}
	if value := res.Get(prefix + "integrity.md5"); value.Exists() {
		data.IntegrityMd5 = types.BoolValue(true)
	} else {
		data.IntegrityMd5 = types.BoolValue(false)
	}
	if value := res.Get(prefix + "integrity.sha1"); value.Exists() {
		data.IntegritySha1 = types.BoolValue(true)
	} else {
		data.IntegritySha1 = types.BoolValue(false)
	}
	if value := res.Get(prefix + "integrity.sha256"); value.Exists() {
		data.IntegritySha256 = types.BoolValue(true)
	} else {
		data.IntegritySha256 = types.BoolValue(false)
	}
	if value := res.Get(prefix + "integrity.sha384"); value.Exists() {
		data.IntegritySha384 = types.BoolValue(true)
	} else {
		data.IntegritySha384 = types.BoolValue(false)
	}
	if value := res.Get(prefix + "integrity.sha512"); value.Exists() {
		data.IntegritySha512 = types.BoolValue(true)
	} else {
		data.IntegritySha512 = types.BoolValue(false)
	}
	if value := res.Get(prefix + "prf.md5"); value.Exists() {
		data.PrfMd5 = types.BoolValue(true)
	} else {
		data.PrfMd5 = types.BoolValue(false)
	}
	if value := res.Get(prefix + "prf.sha1"); value.Exists() {
		data.PrfSha1 = types.BoolValue(true)
	} else {
		data.PrfSha1 = types.BoolValue(false)
	}
	if value := res.Get(prefix + "prf.sha256"); value.Exists() {
		data.PrfSha256 = types.BoolValue(true)
	} else {
		data.PrfSha256 = types.BoolValue(false)
	}
	if value := res.Get(prefix + "prf.sha384"); value.Exists() {
		data.PrfSha384 = types.BoolValue(true)
	} else {
		data.PrfSha384 = types.BoolValue(false)
	}
	if value := res.Get(prefix + "prf.sha512"); value.Exists() {
		data.PrfSha512 = types.BoolValue(true)
	} else {
		data.PrfSha512 = types.BoolValue(false)
	}
}

func (data *CryptoIKEv2Proposal) getDeletedItems(ctx context.Context, state CryptoIKEv2Proposal) []string {
	deletedItems := make([]string, 0)
	if !state.EncryptionEn3des.IsNull() && data.EncryptionEn3des.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/encryption/en-3des", state.getPath()))
	}
	if !state.EncryptionAesCbc128.IsNull() && data.EncryptionAesCbc128.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/encryption/aes-cbc-128", state.getPath()))
	}
	if !state.EncryptionAesCbc192.IsNull() && data.EncryptionAesCbc192.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/encryption/aes-cbc-192", state.getPath()))
	}
	if !state.EncryptionAesCbc256.IsNull() && data.EncryptionAesCbc256.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/encryption/aes-cbc-256", state.getPath()))
	}
	if !state.EncryptionAesGcm128.IsNull() && data.EncryptionAesGcm128.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/encryption/aes-gcm-128", state.getPath()))
	}
	if !state.EncryptionAesGcm256.IsNull() && data.EncryptionAesGcm256.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/encryption/aes-gcm-256", state.getPath()))
	}
	if !state.GroupOne.IsNull() && data.GroupOne.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/group/one", state.getPath()))
	}
	if !state.GroupTwo.IsNull() && data.GroupTwo.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/group/two", state.getPath()))
	}
	if !state.GroupFourteen.IsNull() && data.GroupFourteen.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/group/fourteen", state.getPath()))
	}
	if !state.GroupFifteen.IsNull() && data.GroupFifteen.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/group/fifteen", state.getPath()))
	}
	if !state.GroupSixteen.IsNull() && data.GroupSixteen.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/group/sixteen", state.getPath()))
	}
	if !state.GroupNineteen.IsNull() && data.GroupNineteen.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/group/nineteen", state.getPath()))
	}
	if !state.GroupTwenty.IsNull() && data.GroupTwenty.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/group/twenty", state.getPath()))
	}
	if !state.GroupTwentyOne.IsNull() && data.GroupTwentyOne.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/group/twenty-one", state.getPath()))
	}
	if !state.GroupTwentyFour.IsNull() && data.GroupTwentyFour.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/group/twenty-four", state.getPath()))
	}
	if !state.IntegrityMd5.IsNull() && data.IntegrityMd5.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/integrity/md5", state.getPath()))
	}
	if !state.IntegritySha1.IsNull() && data.IntegritySha1.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/integrity/sha1", state.getPath()))
	}
	if !state.IntegritySha256.IsNull() && data.IntegritySha256.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/integrity/sha256", state.getPath()))
	}
	if !state.IntegritySha384.IsNull() && data.IntegritySha384.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/integrity/sha384", state.getPath()))
	}
	if !state.IntegritySha512.IsNull() && data.IntegritySha512.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/integrity/sha512", state.getPath()))
	}
	if !state.PrfMd5.IsNull() && data.PrfMd5.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/prf/md5", state.getPath()))
	}
	if !state.PrfSha1.IsNull() && data.PrfSha1.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/prf/sha1", state.getPath()))
	}
	if !state.PrfSha256.IsNull() && data.PrfSha256.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/prf/sha256", state.getPath()))
	}
	if !state.PrfSha384.IsNull() && data.PrfSha384.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/prf/sha384", state.getPath()))
	}
	if !state.PrfSha512.IsNull() && data.PrfSha512.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/prf/sha512", state.getPath()))
	}
	return deletedItems
}

func (data *CryptoIKEv2Proposal) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	if !data.EncryptionEn3des.IsNull() && !data.EncryptionEn3des.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/encryption/en-3des", data.getPath()))
	}
	if !data.EncryptionAesCbc128.IsNull() && !data.EncryptionAesCbc128.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/encryption/aes-cbc-128", data.getPath()))
	}
	if !data.EncryptionAesCbc192.IsNull() && !data.EncryptionAesCbc192.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/encryption/aes-cbc-192", data.getPath()))
	}
	if !data.EncryptionAesCbc256.IsNull() && !data.EncryptionAesCbc256.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/encryption/aes-cbc-256", data.getPath()))
	}
	if !data.EncryptionAesGcm128.IsNull() && !data.EncryptionAesGcm128.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/encryption/aes-gcm-128", data.getPath()))
	}
	if !data.EncryptionAesGcm256.IsNull() && !data.EncryptionAesGcm256.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/encryption/aes-gcm-256", data.getPath()))
	}
	if !data.GroupOne.IsNull() && !data.GroupOne.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/group/one", data.getPath()))
	}
	if !data.GroupTwo.IsNull() && !data.GroupTwo.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/group/two", data.getPath()))
	}
	if !data.GroupFourteen.IsNull() && !data.GroupFourteen.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/group/fourteen", data.getPath()))
	}
	if !data.GroupFifteen.IsNull() && !data.GroupFifteen.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/group/fifteen", data.getPath()))
	}
	if !data.GroupSixteen.IsNull() && !data.GroupSixteen.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/group/sixteen", data.getPath()))
	}
	if !data.GroupNineteen.IsNull() && !data.GroupNineteen.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/group/nineteen", data.getPath()))
	}
	if !data.GroupTwenty.IsNull() && !data.GroupTwenty.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/group/twenty", data.getPath()))
	}
	if !data.GroupTwentyOne.IsNull() && !data.GroupTwentyOne.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/group/twenty-one", data.getPath()))
	}
	if !data.GroupTwentyFour.IsNull() && !data.GroupTwentyFour.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/group/twenty-four", data.getPath()))
	}
	if !data.IntegrityMd5.IsNull() && !data.IntegrityMd5.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/integrity/md5", data.getPath()))
	}
	if !data.IntegritySha1.IsNull() && !data.IntegritySha1.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/integrity/sha1", data.getPath()))
	}
	if !data.IntegritySha256.IsNull() && !data.IntegritySha256.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/integrity/sha256", data.getPath()))
	}
	if !data.IntegritySha384.IsNull() && !data.IntegritySha384.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/integrity/sha384", data.getPath()))
	}
	if !data.IntegritySha512.IsNull() && !data.IntegritySha512.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/integrity/sha512", data.getPath()))
	}
	if !data.PrfMd5.IsNull() && !data.PrfMd5.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/prf/md5", data.getPath()))
	}
	if !data.PrfSha1.IsNull() && !data.PrfSha1.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/prf/sha1", data.getPath()))
	}
	if !data.PrfSha256.IsNull() && !data.PrfSha256.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/prf/sha256", data.getPath()))
	}
	if !data.PrfSha384.IsNull() && !data.PrfSha384.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/prf/sha384", data.getPath()))
	}
	if !data.PrfSha512.IsNull() && !data.PrfSha512.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/prf/sha512", data.getPath()))
	}
	return emptyLeafsDelete
}

func (data *CryptoIKEv2Proposal) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.EncryptionEn3des.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/encryption/en-3des", data.getPath()))
	}
	if !data.EncryptionAesCbc128.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/encryption/aes-cbc-128", data.getPath()))
	}
	if !data.EncryptionAesCbc192.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/encryption/aes-cbc-192", data.getPath()))
	}
	if !data.EncryptionAesCbc256.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/encryption/aes-cbc-256", data.getPath()))
	}
	if !data.EncryptionAesGcm128.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/encryption/aes-gcm-128", data.getPath()))
	}
	if !data.EncryptionAesGcm256.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/encryption/aes-gcm-256", data.getPath()))
	}
	if !data.GroupOne.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/group/one", data.getPath()))
	}
	if !data.GroupTwo.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/group/two", data.getPath()))
	}
	if !data.GroupFourteen.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/group/fourteen", data.getPath()))
	}
	if !data.GroupFifteen.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/group/fifteen", data.getPath()))
	}
	if !data.GroupSixteen.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/group/sixteen", data.getPath()))
	}
	if !data.GroupNineteen.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/group/nineteen", data.getPath()))
	}
	if !data.GroupTwenty.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/group/twenty", data.getPath()))
	}
	if !data.GroupTwentyOne.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/group/twenty-one", data.getPath()))
	}
	if !data.GroupTwentyFour.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/group/twenty-four", data.getPath()))
	}
	if !data.IntegrityMd5.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/integrity/md5", data.getPath()))
	}
	if !data.IntegritySha1.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/integrity/sha1", data.getPath()))
	}
	if !data.IntegritySha256.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/integrity/sha256", data.getPath()))
	}
	if !data.IntegritySha384.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/integrity/sha384", data.getPath()))
	}
	if !data.IntegritySha512.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/integrity/sha512", data.getPath()))
	}
	if !data.PrfMd5.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/prf/md5", data.getPath()))
	}
	if !data.PrfSha1.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/prf/sha1", data.getPath()))
	}
	if !data.PrfSha256.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/prf/sha256", data.getPath()))
	}
	if !data.PrfSha384.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/prf/sha384", data.getPath()))
	}
	if !data.PrfSha512.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/prf/sha512", data.getPath()))
	}
	return deletePaths
}
