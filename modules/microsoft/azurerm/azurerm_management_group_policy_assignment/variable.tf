variable "naming_options" {
  type = object({
    resource_name = optional(string, "mgpa")
    suffix        = optional(string, null)
    prefix        = optional(string, null)
    separator     = optional(string, "-")
    identifier    = optional(string, null)
    lower         = optional(bool, true)
    length        = optional(number, 24)
  })
  default     = {}
  description = "(Optional) Options to pass to the name generator"
}

variable "management_group_id" {
  type        = string
  description = "(Required) The ID of the Management Group to which to attach this Policy Assignment."
}
variable "policy_definition_id" {
  type        = string
  description = "(Required) The ID of the policy definition to assign to the management group."
}
variable "enforce" {
  type        = bool
  default     = false
  description = "(Optional) Specifies whether the policy assignment should be enforced. Defaults to false."
}

variable "parameters" {
  type        = map(object({ value = any }))
  default     = {}
  description = "(Optional) A map of key/value pairs to pass to the policy definition as parameters."
}