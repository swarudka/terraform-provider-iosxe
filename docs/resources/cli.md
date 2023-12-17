---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxe_cli Resource - terraform-provider-iosxe"
subcategory: "General"
description: |-
  This resources is used to configure arbitrary CLI commands. This should be considered a last resort in case YANG models are not available, as it cannot read the state and therefore cannot reconcile changes.
---

# iosxe_cli (Resource)

This resources is used to configure arbitrary CLI commands. This should be considered a last resort in case YANG models are not available, as it cannot read the state and therefore cannot reconcile changes.

## Example Usage

```terraform
resource "iosxe_cli" "example" {
  cli = <<-EOT
  interface Loopback123
  description configured-via-restconf-cli
  EOT
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `cli` (String) This attribute contains the CLI commands.

### Optional

- `device` (String) A device name from the provider configuration.