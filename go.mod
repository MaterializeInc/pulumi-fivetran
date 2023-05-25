module github.com/benesch/pulumi-fivetran

go 1.16

require (
	github.com/benesch/terraform-provider-fivetran v0.1.3-0.20220822140138-40ecdacb1b3a
	github.com/hashicorp/terraform-plugin-sdk v1.17.2 // indirect
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.47.0
	github.com/pulumi/pulumi/pkg/v3 v3.68.0 // indirect
	github.com/pulumi/pulumi/sdk/v3 v3.68.0
	golang.org/x/mod v0.10.0
)

replace github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20230327102345-3fa930f86570
