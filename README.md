# RevBits - terraform-provider-pam
## Overview
The provider manages authentication with **RevBits-PAM**, allowing Terraform to fetch and use secrets stored in PAM. The provider includes the following features and benefits:

* Simple setup in the Terraform manifest.

* The provider authenticates to **RevBits-PAM**.

* A provider method fetches variable values and makes them available for use elsewhere in the manifest.

* The Terraform sensitive flag may be used against any fetched secret value to keep the value from appearing in logs and on-screen.

## Usage and Workflow  
Terraform can be run manually by users, but it is often run by machines. To access the values of secrets, the user/machine needs execute privileges on the RevBits PAM variables referenced in your Terraform manifests.
 
### Provider configuration
The provider uses **pam-api-go** to load its configuration. pam-api-go can be configured using environment variables or using the provider configuration in the .tf file. To run the example, create a terraform.tfvars:
``` 
domain = "Enter Partner Domain"
secret_id = "Enter Secret ID here"
api_key =  "Enter API Key here" 
```

## Fetch secrets  
An important thing to keep in mind is that by design Terraform state files can contain sensitive data (which may include credentials fetched by this plugin). 

For Terraform 0.13+, include the terraform block in your configuration or plan to that specifies the provider:


```
##main.tf

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

}
 ```
Secrets like data.pam_provider.secret.secret_value can be used in any Terraform resources.

