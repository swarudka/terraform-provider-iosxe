---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxe_flow_record Data Source - terraform-provider-iosxe"
subcategory: "Flow"
description: |-
  This data source can read the Flow Record configuration.
---

# iosxe_flow_record (Data Source)

This data source can read the Flow Record configuration.

## Example Usage

```terraform
data "iosxe_flow_record" "example" {
  name = "FNF1"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String)

### Optional

- `device` (String) A device name from the provider configuration.

### Read-Only

- `collect_counter_bytes_long` (Boolean) Total number of bytes (64 bit counter)
- `collect_counter_packets_long` (Boolean) Total number of packets (64 bit counter)
- `collect_interface_output` (Boolean) The output interface
- `collect_timestamp_absolute_first` (Boolean) Absolute time the first packet was seen (milliseconds)
- `collect_timestamp_absolute_last` (Boolean) Absolute time the most recent packet was seen (milliseconds)
- `collect_transport_tcp_flags` (Boolean) TCP flags
- `description` (String) Provide a description for this Flow Record
- `id` (String) The path of the retrieved object.
- `match_flow_direction` (Boolean) Direction the flow was monitored in
- `match_interface_input` (Boolean) The input interface
- `match_ipv4_destination_address` (Boolean) IPv4 destination address
- `match_ipv4_protocol` (Boolean) IPv4 protocol
- `match_ipv4_source_address` (Boolean) IPv4 source address
- `match_ipv4_tos` (Boolean) IPv4 type of service
- `match_transport_destination_port` (Boolean) Transport destination port
- `match_transport_source_port` (Boolean) Transport source port