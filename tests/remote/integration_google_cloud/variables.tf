variable "client_email" {
  type = string
}

variable "private_key" {
  type      = string
  sensitive = true
}

variable "name" {
  type = string
}

variable "ownership_tag_keys" {
  type    = list(string)
  default = null
}

variable "ownership_tag_overrides" {
  type    = bool
  default = null
}
