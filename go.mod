module github.com/moneymeets/pulumi-heroku

go 1.12

replace (
	git.apache.org/thrift.git => github.com/apache/thrift v0.0.0-20180902110319-2566ecd5d999
	github.com/Nvveen/Gotty => github.com/ijc25/Gotty v0.0.0-20170406111628-a8b993ba6abd
	github.com/golang/glog => github.com/pulumi/glog v0.0.0-20180820174630-7eaa6ffb71e4
)

require (
	github.com/hashicorp/terraform v0.12.7
	github.com/pulumi/pulumi v1.0.0
	github.com/pulumi/pulumi-terraform v0.18.4-0.20190905205929-ed0b5c29edd1
	github.com/pulumi/scripts v0.0.0-20190821175317-4a3c5c021139 // indirect
	github.com/stretchr/testify v1.3.1-0.20190311161405-34c6fa2dc709
	github.com/terraform-providers/terraform-provider-heroku v0.0.0-20190809201753-73779273aff7
)
