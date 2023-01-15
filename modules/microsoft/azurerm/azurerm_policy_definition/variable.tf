variable "naming_options" {
  type = object({
    resource_name = optional(string, "pd")
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
  description = "(Optional) The display name of the policy definition."
}

variable "description" {
  type        = string
  default     = null
  description = "(Optional) The description of the policy definition."
}

variable "mode" {
  type        = string
  default     = "Indexed"
  description = "(Optional) The mode of the policy definition."
  validation {
    condition     = can(regex("^(All|Indexed|Microsoft.ContainerService.Data|Microsoft.CustomerLockbox.Data|Microsoft.DataCatalog.Data|Microsoft.KeyVault.Data|Microsoft.Kubernetes.Data|Microsoft.MachineLearningServices.Data|Microsoft.Network.Data and Microsoft.Synapse.Data)$", var.mode))
    error_message = "The value must be one of All, Indexed, Microsoft.ContainerService.Data, Microsoft.CustomerLockbox.Data, Microsoft.DataCatalog.Data, Microsoft.KeyVault.Data, Microsoft.Kubernetes.Data, Microsoft.MachineLearningServices.Data, Microsoft.Network.Data and Microsoft.Synapse.Data."
  }
}

variable "metadata" {
  type        = string
  default     = null
  description = "(Optional) The metadata of the policy definition."
}

variable "parameters" {
  type        = string
  default     = null
  description = "(Optional) The parameters of the policy definition."
}
variable "policy_rule" {
  type        = string
  default     = null
  description = "(Optional) The policy rule of the policy definition."
}