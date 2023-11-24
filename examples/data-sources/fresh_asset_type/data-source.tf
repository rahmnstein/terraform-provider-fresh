terraform {
  required_providers {
    fresh = {
      source  = "registry.terraform.io/rahmnstein/fresh"
      version = "0.1.0"
    }
  }
}

data "fresh_asset_type" "Laptop" {
  name = "VMware VCenter VM"
}

output "name" {
  value = data.fresh_asset_type.Laptop.id

}
