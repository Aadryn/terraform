# [ Naming ]
# ----------------------------------------------------------------------------------------------------
module "naming" {
  source         = "git@github.com:Aadryn/terraform.git//modules/common/naming?ref=main"
  naming_options = var.naming_options

}

# [ Policy Definition ]
# ----------------------------------------------------------------------------------------------------
resource "azurerm_policy_definition" "default" {
  name         = module.naming.name
  policy_type  = var.policy_type
  mode         = var.mode
  display_name = var.display_name
  description  = var.description
  metadata     = var.metadata
  policy_rule  = var.policy_rule
  parameters   = var.parameters
}