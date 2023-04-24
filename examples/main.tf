terraform {
  required_providers {
    cattlenames = {
      source  = "local/cattlenames"
      version = "1.0.0"
    }
  }
}

provider "cattlenames" {}

resource "cattlenames" "example" {
  count = 5
}

resource "cattlenames" "club" {
  count  = 5
  family = "27Club"
}

output "cattle_name" {
  value = cattlenames.example
}

output "cattle_name_27club" {
  value = cattlenames.club
}
