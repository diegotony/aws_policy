variable "name" {
  type = string
  description = "(optional) describe your variable"
  default = null
}

variable "policies" {
  type = list(string)
}

variable "roles" {
  type        = list(string)
  description = "roles list"
  default = []
}

variable "users" {
  type        = list(string)
  description = "users list"
  default = []
}

variable "groups" {
  type        = list(string)
  description = "groups list"
  default = []
}
