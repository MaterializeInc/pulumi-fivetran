module github.com/benesch/pulumi-fivetran

go 1.16

require (
	github.com/benesch/terraform-provider-fivetran v0.1.3-0.20211109215722-69ad9bd0b286
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/hashicorp/terraform-plugin-sdk v1.17.2 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.10.0
	github.com/pulumi/pulumi/sdk/v3 v3.16.0
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/mod v0.4.2
	golang.org/x/term v0.0.0-20210220032956-6a3ed077a48d // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20201218231525-9cca98608a5e
