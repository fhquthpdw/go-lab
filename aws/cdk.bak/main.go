package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseventstargets"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type LambdaCronStackProps struct {
	awscdk.StackProps
}

func NewLambdaCronStack(scope constructs.Construct, id string, props *LambdaCronStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	lambdaFn := awslambda.NewFunction(stack, jsii.String("Singleton"), &awslambda.FunctionProps{
		Code:    awslambda.NewAssetCode(jsii.String("lambda"), nil),
		Handler: jsii.String("handler.main"),
		Timeout: awscdk.Duration_Seconds(jsii.Number(300)),
		Runtime: awslambda.Runtime_PYTHON_3_6(),
	})

	rule := awsevents.NewRule(stack, jsii.String("Rule"), &awsevents.RuleProps{
		Schedule: awsevents.Schedule_Expression(jsii.String("cron(0 18 ? * MON-FRI *)")),
	})
	rule.AddTarget(awseventstargets.NewLambdaFunction(lambdaFn, nil))

	return stack
}

func main() {
	app := awscdk.NewApp(nil)

	NewLambdaCronStack(app, "LambdaCronStack", &LambdaCronStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	return nil
}
