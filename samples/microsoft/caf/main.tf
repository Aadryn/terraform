# [ Terraform - Providers ]
# ----------------------------------------------------------------------------------------------------
provider "azurerm" {
  features {}
}

locals {
  naming_options = {
    prefix = "en"
    suffix = "par"
  }
}

# [ Management Group - V0 ]
# ----------------------------------------------------------------------------------------------------
# { v0 }
module "v0" {
  source         = "git@github.com:Aadryn/terraform.git//modules/microsoft/azurerm/azurerm_management_group?ref=main"
  display_name   = "V0"
  naming_options = merge(local.naming_options, { suffix = join("-", [ local.naming_options.suffix, "v0" ]) })
}

# [ Management Group - V1]
# ----------------------------------------------------------------------------------------------------
# { v1 }
module "v1" {
  source         = "git@github.com:Aadryn/terraform.git//modules/microsoft/azurerm/azurerm_management_group?ref=main"
  display_name   = "V1"
  naming_options = merge(local.naming_options, { suffix = join("-", [ local.naming_options.suffix, "v1" ]) })
}

# { sandbox }
module "sandbox" {
  source                     = "git@github.com:Aadryn/terraform.git//modules/microsoft/azurerm/azurerm_management_group?ref=main"
  display_name               = "Sandbox"
  naming_options             = merge(local.naming_options, {
    suffix = join("-", [ local.naming_options.suffix, "sandbox" ])
  })
  parent_management_group_id = module.v1.id
}

# { platform }
module "platform" {
  source                     = "git@github.com:Aadryn/terraform.git//modules/microsoft/azurerm/azurerm_management_group?ref=main"
  display_name               = "Platform"
  naming_options             = merge(local.naming_options, {
    suffix = join("-", [ local.naming_options.suffix, "platform" ])
  })
  parent_management_group_id = module.v1.id
}

# { legacy }
module "legacy" {
  source                     = "git@github.com:Aadryn/terraform.git//modules/microsoft/azurerm/azurerm_management_group?ref=main"
  display_name               = "Legacy"
  naming_options             = merge(local.naming_options, {
    suffix = join("-", [ local.naming_options.suffix, "legacy" ])
  })
  parent_management_group_id = module.v1.id
}

# { lz }
module "lz" {
  source                     = "git@github.com:Aadryn/terraform.git//modules/microsoft/azurerm/azurerm_management_group?ref=main"
  display_name               = "Landing Zones"
  naming_options             = merge(local.naming_options, {
    suffix = join("-", [ local.naming_options.suffix, "lz" ])
  })
  parent_management_group_id = module.v1.id
}