## ✨ Get started

- See the [examples](./examples/) directory for detailed examples.

```tf
terraform {
  required_providers {
    fresh = {
      source  = "registry.terraform.io/rahmnstein/fresh"
      version = "0.1.0"
    }
  }
}

provider "fresh" {
  # Configuration options
  address = "https://MYDOMAIN.freshservice.com/api/v2"
  api_key = "MyApiKey"
}

# Retrieve an asset by its display ID.
data "fresh_asset" "test" {
  display_id = "7308"
}

# Create a new asset type for a VMware VM.
data "fresh_asset_type" "vmware" {
  name = "VMware VCenter VM"
}

resource "fresh_asset" "test" {
  name          = "TestAssetTerraform"
  asset_type_id = data.fresh_asset_type.vmware.id
  description   = "Description of TestAssetTerraform"
}
```
