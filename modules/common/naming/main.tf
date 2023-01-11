resource "random_id" "default" {
  count       = local.has_identifier ? 0 : 1
  byte_length = 4
}

locals {
  has_identifier = var.naming_options.identifier != null
  identifier     = local.has_identifier ? var.naming_options.identifier : random_id.default[0].hex
  items = [
    var.naming_options.prefix, var.naming_options.resource_name, var.naming_options.suffix, local.identifier
  ]
  join        = join(var.naming_options.separator, compact(local.items))
  join_simple = join(var.naming_options.separator, compact([var.naming_options.resource_name, local.identifier]))
  case        = var.naming_options.lower ? lower(local.join) : upper(local.join)
  case_simple = var.naming_options.lower ? lower(local.join_simple) : upper(local.join_simple)
  rendered    = length(local.case) > var.naming_options.length && var.naming_options.length != 0 ? local.case_simple : local.case
}
