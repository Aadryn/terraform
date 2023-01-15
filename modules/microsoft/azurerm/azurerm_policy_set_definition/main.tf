# [ Naming ]
# ----------------------------------------------------------------------------------------------------
module "naming" {
  source         = "git@github.com:Aadryn/terraform.git//modules/common/naming?ref=main"
  naming_options = var.naming_options

}

# [ Management Group ]
# ----------------------------------------------------------------------------------------------------
data "azurerm_policy_definition" "default" {
  count        = length(var.policy_definition_references)
  display_name = var.policy_definition_references[ count.index ]
}

resource "azurerm_policy_set_definition" "default" {
  name                = module.naming.name
  policy_type         = var.policy_type
  display_name        = var.display_name == null ? module.naming.name : var.display_name
  management_group_id = var.management_group_id

  metadata = <<METADATA
    {
    "category": "${var.policy_type}"
    }
METADATA

  dynamic "policy_definition_reference" {
    for_each = data.azurerm_policy_definition.default
    content {
      policy_definition_id = policy_definition_reference.value[ "id" ]
      reference_id         = policy_definition_reference.value[ "id" ]
    }
  }
}