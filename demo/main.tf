terraform {
  required_providers {
    pam = {
      source = "registry.terraform.io/revbits/pam"
      version = "~> 1.0"
    }
  }
}

variable "domain" {
  type = string
}
variable "secret_id" {
  type = string
}
variable "api_key" {
  type = string
}

provider "pam" {
  domain = var.domain
  secret_id = var.secret_id
  api_key = var.api_key
}

data "pam_provider" "secret" {}


output "secret_output" {
  value = [data.pam_provider.secret.secret_value]
}
