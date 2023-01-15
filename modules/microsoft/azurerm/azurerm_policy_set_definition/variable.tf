variable "naming_options" {
  type = object({
    resource_name = optional(string, "psd")
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

variable "policy_type" {
  type        = string
  default     = "NotSpecified"
  description = "(Optional) The policy set type. Valid values are NotSpecified, BuiltIn, and Custom. Defaults to NotSpecified."
  validation {
    condition     = can(regex("^(BuiltIn|Custom|NotSpecified|Static)$", var.policy_type))
    error_message = "The value must be one of BuiltIn, Custom, NotSpecified, Static."
  }
}

variable "display_name" {
  type        = string
  default     = null
  description = "(Optional) The display name of the policy set definition."
}

variable "management_group_id" {
  type        = string
  default     = null
  description = "(Optional) The ID of the management group where the policy set definition should be created."
}

variable "policy_definition_references" {
  type        = list(string)
  description = "(Optional) A list of policy definition references in the policy set definition."
  default     = [ ]
}