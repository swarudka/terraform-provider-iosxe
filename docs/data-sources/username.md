---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxe_username Data Source - terraform-provider-iosxe"
subcategory: "System"
description: |-
  This data source can read the Username configuration.
---

# iosxe_username (Data Source)

This data source can read the Username configuration.

## Example Usage

```terraform
data "iosxe_username" "example" {
  name = "user1"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String)

### Optional

- `device` (String) A device name from the provider configuration.

### Read-Only

- `description` (String) description string with max 128 characters
- `id` (String) The path of the retrieved object.
- `password` (String)
- `password_encryption` (String)
- `privilege` (Number) Set user privilege level
- `secret` (String)
- `secret_encryption` (String)