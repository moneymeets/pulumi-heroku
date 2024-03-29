module github.com/moneymeets/pulumi-heroku/provider

go 1.16

replace (
	github.com/hashicorp/go-getter v1.5.0 => github.com/hashicorp/go-getter v1.4.0
	github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20210629210550-59d24255d71f
)

require (
	github.com/docker/docker v1.6.1
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/hashicorp/terraform-plugin-sdk v1.15.0 // indirect
	github.com/heroku/terraform-provider-heroku/v4 v4.8.0
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.16.1
	github.com/pulumi/pulumi/sdk/v3 v3.22.0
)
