package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"io"
	"net/http"
)

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
	lambda.Start(handler)
}
