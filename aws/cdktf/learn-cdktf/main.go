package main

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/cdktf-provider-aws-go/aws/v9"
	"github.com/hashicorp/cdktf-provider-aws-go/aws/v9/ec2"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func NewMyStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	aws.NewAwsProvider(stack, jsii.String("AWS"), &aws.AwsProviderConfig{
		Region: jsii.String("cn-northwest-1"),
	})

	instance := ec2.NewInstance(stack, jsii.String("daochunCompute"), &ec2.InstanceConfig{
		Ami:          jsii.String("ami-01dee4c3e06906189"),
		InstanceType: jsii.String("t2.micro"),
		Tags: &map[string]*string{
			"Name": jsii.String("CDKTF-Demo"),
		},
	})

	cdktf.NewTerraformOutput(stack, jsii.String("public_ip"), &cdktf.TerraformOutputConfig{
		Value: instance.PublicIp(),
	})

	return stack
}

func main() {
	app := cdktf.NewApp(nil)

	stack := NewMyStack(app, "aws_instance")
	cdktf.NewCloudBackend(stack, &cdktf.CloudBackendProps{
		Hostname:     jsii.String("app.terraform.io"),
		Organization: jsii.String("daochun"),
		Workspaces:   cdktf.NewNamedCloudWorkspace(jsii.String("learn-cdktf")),
	})

	app.Synth()
}
