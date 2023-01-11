package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"io"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

// EC2 status check
type Instance struct {
	InstanceIDs []string `json:"InstanceID"`
}

var client *ec2.Client

func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("cn-northwest-1"))
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	client = ec2.NewFromConfig(cfg)
}

func HandleRequest(instances Instance) ([]string, error) {

	result, err := client.DescribeInstances(context.TODO(), &ec2.DescribeInstancesInput{
		InstanceIds: instances.InstanceIDs,
	})
	if err != nil {
		return []string{}, err
	}

	var status []string
	for _, r := range result.Reservations {
		for _, i := range r.Instances {
			status = append(status, fmt.Sprintf("InstanceID: %v State: %v", *i.InstanceId, i.State.Name))
		}

		fmt.Println("")
	}

	return status, nil
}

// HTTP
var (
	DefaultHttpGetAddress = "https://6tcusl0x96.execute-api.cn-northwest-1.amazonaws.com.cn/default/http-trigger"
	ErrNon200Response     = fmt.Errorf("non 200 Response found")
	ErrNoIp               = fmt.Errorf("no IP in HTTP response")
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	resp, err := http.Get(DefaultHttpGetAddress)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	if resp.StatusCode != 200 {
		return events.APIGatewayProxyResponse{}, ErrNon200Response
	}

	ip, err := io.ReadAll(resp.Body)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	if len(ip) == 0 {
		return events.APIGatewayProxyResponse{}, ErrNoIp
	}

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("Hello, %v", string(ip)),
		StatusCode: http.StatusOK,
	}, nil
}

func main() {
	//lambda.Start(HandleRequest)
	lambda.Start(handler)
}
