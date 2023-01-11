terraform {
  required_providers {
    random = {
      source = "hashicorp/random"
      version = ">= 3.4.6"
    }
  }
  required_version = ">= 1.3.7"
}