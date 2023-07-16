---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxe_msdp Resource - terraform-provider-iosxe"
subcategory: "Multicast"
description: |-
  This resource can manage the MSDP configuration.
---

# iosxe_msdp (Resource)

This resource can manage the MSDP configuration.

## Example Usage

```terraform
resource "iosxe_msdp" "example" {
  originator_id = "Loopback100"
  passwords = [
    {
      addr       = "10.1.1.1"
      encryption = 0
      password   = "Cisco123"
    }
  ]
  peers = [
    {
      addr                    = "10.1.1.1"
      remote_as               = 65000
      connect_source_loopback = 100
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `delete_mode` (String) Configure behavior when deleting/destroying the resource. Either delete the entire object (YANG container) being managed, or only delete the individual resource attributes configured explicitly and leave everything else as-is. Default value is `all`.
  - Choices: `all`, `attributes`
- `device` (String) A device name from the provider configuration.
- `originator_id` (String) Configure MSDP Originator ID
- `passwords` (Attributes List) MSDP peer on which the password is to be set (see [below for nested schema](#nestedatt--passwords))
- `peers` (Attributes List) Configure an MSDP peer (see [below for nested schema](#nestedatt--peers))

### Read-Only

- `id` (String) The path of the object.

<a id="nestedatt--passwords"></a>
### Nested Schema for `passwords`

Required:

- `addr` (String)
- `password` (String)

Optional:

- `encryption` (Number) - Range: `0`-`7`


<a id="nestedatt--peers"></a>
### Nested Schema for `peers`

Required:

- `addr` (String)

Optional:

- `connect_source_loopback` (Number) Loopback interface
  - Range: `0`-`2147483647`
- `remote_as` (Number) Configured AS number
  - Range: `1`-`65535`

## Import

Import is supported using the following syntax:

```shell
terraform import iosxe_msdp.example "Cisco-IOS-XE-native:native/ip/Cisco-IOS-XE-multicast:msdp"
```