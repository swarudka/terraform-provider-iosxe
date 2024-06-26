---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxe_crypto_ikev2_policy Resource - terraform-provider-iosxe"
subcategory: "Crypto"
description: |-
  This resource can manage the Crypto IKEv2 Policy configuration.
---

# iosxe_crypto_ikev2_policy (Resource)

This resource can manage the Crypto IKEv2 Policy configuration.

## Example Usage

```terraform
resource "iosxe_crypto_ikev2_policy" "example" {
  name                   = "policy1"
  match_address_local_ip = ["1.2.3.4"]
  match_fvrf_any         = true
  proposals = [
    {
      proposals = "proposal123"
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String)
- `proposals` (Attributes List) Specify Proposal (see [below for nested schema](#nestedatt--proposals))

### Optional

- `device` (String) A device name from the provider configuration.
- `match_address_local_ip` (List of String) Local address
- `match_fvrf` (String)
- `match_fvrf_any` (Boolean) Any fvrf
- `match_inbound_only` (Boolean) inbound only for controller

### Read-Only

- `id` (String) The path of the object.

<a id="nestedatt--proposals"></a>
### Nested Schema for `proposals`

Required:

- `proposals` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import iosxe_crypto_ikev2_policy.example "Cisco-IOS-XE-native:native/crypto/Cisco-IOS-XE-crypto:ikev2/policy=policy1"
```
