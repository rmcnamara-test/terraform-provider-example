terraform {
  required_providers {
    example = {
      version = "~> 1.0.0"
      source  = "localhost/example/example"
    }
  }
}
provider "example" {
}

data "example" "example" {}

output "test" {
  value = "test"
}
