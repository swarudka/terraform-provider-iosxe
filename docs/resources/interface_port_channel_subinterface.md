---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxe_interface_port_channel_subinterface Resource - terraform-provider-iosxe"
subcategory: "Interface"
description: |-
  This resource can manage the Interface Port Channel Subinterface configuration.
---

# iosxe_interface_port_channel_subinterface (Resource)

This resource can manage the Interface Port Channel Subinterface configuration.

## Example Usage

```terraform
resource "iosxe_interface_port_channel_subinterface" "example" {
  name                        = "10.666"
  description                 = "My Interface Description"
  shutdown                    = false
  ip_proxy_arp                = false
  ip_redirects                = false
  unreachables                = false
  vrf_forwarding              = "VRF1"
  ipv4_address                = "192.0.2.2"
  ipv4_address_mask           = "255.255.255.0"
  encapsulation_dot1q_vlan_id = 666
  ip_access_group_in          = "1"
  ip_access_group_in_enable   = true
  ip_access_group_out         = "1"
  ip_access_group_out_enable  = true
  helper_addresses = [
    {
      address = "10.10.10.10"
      global  = false
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String)

### Optional

- `auto_qos_classify` (Boolean) Configure classification for untrusted devices
- `auto_qos_classify_police` (Boolean) Configure QoS policing for untrusted devices
- `auto_qos_trust` (Boolean) Trust the DSCP/CoS marking
- `auto_qos_trust_cos` (Boolean) Trust the CoS marking
- `auto_qos_trust_dscp` (Boolean) Trust the DSCP marking
- `auto_qos_video_cts` (Boolean) Trust the QoS marking of the Cisco Telepresence System
- `auto_qos_video_ip_camera` (Boolean) Trust the QoS marking of the Ip Video Surveillance camera
- `auto_qos_video_media_player` (Boolean) Trust the Qos marking of the Cisco Media Player
- `auto_qos_voip` (Boolean) Configure AutoQoS for VoIP
- `auto_qos_voip_cisco_phone` (Boolean) Trust the QoS marking of Cisco IP Phone
- `auto_qos_voip_cisco_softphone` (Boolean) Trust the QoS marking of Cisco IP SoftPhone
- `auto_qos_voip_trust` (Boolean) Trust the DSCP/CoS marking
- `delete_mode` (String) Configure behavior when deleting/destroying the resource. Either delete the entire object (YANG container) being managed, or only delete the individual resource attributes configured explicitly and leave everything else as-is. Default value is `all`.
  - Choices: `all`, `attributes`
- `description` (String) Interface specific description
- `device` (String) A device name from the provider configuration.
- `encapsulation_dot1q_vlan_id` (Number) - Range: `1`-`4094`
- `helper_addresses` (Attributes List) Specify a destination address for UDP broadcasts (see [below for nested schema](#nestedatt--helper_addresses))
- `ip_access_group_in` (String)
- `ip_access_group_in_enable` (Boolean) inbound packets
- `ip_access_group_out` (String)
- `ip_access_group_out_enable` (Boolean) outbound packets
- `ip_proxy_arp` (Boolean) Enable proxy ARP
- `ip_redirects` (Boolean) Enable sending ICMP Redirect messages
- `ipv4_address` (String)
- `ipv4_address_mask` (String)
- `shutdown` (Boolean) Shutdown the selected interface
- `trust_device` (String) trusted device class
  - Choices: `cisco-phone`, `cts`, `ip-camera`, `media-player`
- `unreachables` (Boolean) Enable sending ICMP Unreachable messages
- `vrf_forwarding` (String) Configure forwarding table

### Read-Only

- `id` (String) The path of the object.

<a id="nestedatt--helper_addresses"></a>
### Nested Schema for `helper_addresses`

Required:

- `address` (String)

Optional:

- `global` (Boolean) Helper-address is global
- `vrf` (String) VRF name for helper-address (if different from interface VRF)

## Import

Import is supported using the following syntax:

```shell
terraform import iosxe_interface_port_channel_subinterface.example "Cisco-IOS-XE-native:native/interface/Port-channel-subinterface/Port-channel=10.666"
```