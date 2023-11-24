terraform {
  required_providers {
    fresh = {
      source  = "registry.terraform.io/rahmnstein/fresh"
      version = "0.1.0"
    }
  }
}

data "fresh_asset" "test" {
  display_id = "7308"
}

output "test" {
  value = data.fresh_asset.test.name
}
