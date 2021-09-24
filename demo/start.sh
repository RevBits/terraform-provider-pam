#!/bin/bash
mkdir -p registry.terraform.io/revbits/pam/1.0.0/darwin_amd64
go build -o ./registry.terraform.io/revbits/pam/1.0.0/darwin_amd64/terraform-provider-pam ../main.go
rm -rf .terraform.lock.hcl
terraform init -plugin-dir=.
TF_LOG=TRACE terraform apply

