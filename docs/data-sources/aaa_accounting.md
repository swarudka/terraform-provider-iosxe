---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxe_aaa_accounting Data Source - terraform-provider-iosxe"
subcategory: "AAA"
description: |-
  This data source can read the AAA Accounting configuration.
---

# iosxe_aaa_accounting (Data Source)

This data source can read the AAA Accounting configuration.

## Example Usage

```terraform
data "iosxe_aaa_accounting" "example" {
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `device` (String) A device name from the provider configuration.

### Read-Only

- `execs` (Attributes List) For starting an exec (shell). (see [below for nested schema](#nestedatt--execs))
- `id` (String) The path of the retrieved object.
- `identity_default_start_stop_group1` (String) Use Server-group
- `identity_default_start_stop_group2` (String) Use Server-group
- `identity_default_start_stop_group3` (String) Use Server-group
- `identity_default_start_stop_group4` (String) Use Server-group
- `networks` (Attributes List) For network services. (PPP, SLIP, ARAP) (see [below for nested schema](#nestedatt--networks))
- `system_guarantee_first` (Boolean) Guarantee system accounting as first record.
- `update_newinfo_periodic` (Number) Periodic intervals to send accounting update records(in minutes)

<a id="nestedatt--execs"></a>
### Nested Schema for `execs`

Read-Only:

- `name` (String)
- `start_stop_group1` (String) Use Server-group


<a id="nestedatt--networks"></a>
### Nested Schema for `networks`

Read-Only:

- `id` (String)
- `start_stop_group1` (String) Use Server-group
- `start_stop_group2` (String) Use Server-group