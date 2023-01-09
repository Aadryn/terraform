variable "naming_options" {
  type = object({
    resource_name = optional(string, null)
    suffix        = optional(string, null)
    prefix        = optional(string, null)
    separator     = optional(string, "-")
    identifier    = optional(string, null)
    lower         = optional(bool, true)
    length        = optional(number, 0)
  })
  default = {}
  description = "(Optional) Options to pass to the name generator"
}