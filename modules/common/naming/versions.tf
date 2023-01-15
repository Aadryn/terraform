terraform {
  required_providers {
    random = {
      source = "hashicorp/random"
      version = ">= 3.4.3"
    }
  }
  required_version = ">= 0.13"
}