module github.com/RevBits/terraform-provider-pam

go 1.16

require (
	github.com/RevBits/pam-api-go v1.0.0
	github.com/hashicorp/terraform-plugin-sdk v1.17.2
)

//replace github.com/RevBits/pam-api-go => ../pam
