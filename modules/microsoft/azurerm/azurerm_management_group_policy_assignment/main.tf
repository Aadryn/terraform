# [ Naming ]
# ----------------------------------------------------------------------------------------------------
module "naming" {
  source         = "git@github.com:Aadryn/terraform.git//modules/common/naming?ref=main"
  naming_options = var.naming_options

}

# [ Management Group ]
# ----------------------------------------------------------------------------------------------------
resource "azurerm_management_group_policy_assignment" "default" {
  management_group_id  = var.management_group_id
  policy_definition_id = var.policy_definition_id
  enforce              = var.enforce
  name                 = "default"
  parameters           = <<PARAMETERS
    ${jsonencode(var.parameters)}
  PARAMETERS
}