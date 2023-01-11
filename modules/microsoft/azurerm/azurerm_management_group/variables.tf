variable "naming_options" {
  type = object({
    resource_name = optional(string, "mg")
    suffix        = optional(string, null)
    prefix        = optional(string, null)
    separator     = optional(string, "-")
    identifier    = optional(string, null)
    lower         = optional(bool, true)
    length        = optional(number, 0)
  })
  default     = {}
  description = "(Optional) Options to pass to the name generator"
}

variable "display_name" {
  type        = string
  description = "(Optional) The display name of the resource"
}

variable "parent_management_group_id" {
  type        = string
  default     = null
  description = "(Optional) The ID of the parent management group. If not set, the resource will be created at the root level."
}