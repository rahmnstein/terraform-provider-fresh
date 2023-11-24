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
