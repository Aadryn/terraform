# [ Naming ]
# ----------------------------------------------------------------------------------------------------
module "naming" {
  source         = "git@github.com:Aadryn/terraform.git//modules/common/naming?ref=main"
  naming_options = var.naming_options

}

# [ Management Group ]
# ----------------------------------------------------------------------------------------------------
resource "azurerm_management_group" "default" {
  name = module.naming.name
  display_name = var.display_name
  parent_management_group_id = var.parent_management_group_id

}