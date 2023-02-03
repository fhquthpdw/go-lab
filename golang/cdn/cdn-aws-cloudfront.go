package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
)

func viewListInvalidations(cli *cloudfront.Client) error {
	resp, err := cli.ListInvalidations(context.TODO(), &cloudfront.ListInvalidationsInput{
		DistributionId: aws.String(os.Getenv("CLOUD_FRONT_DISTRIBUTION")),
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("list invalidations: ")
	fmt.Printf("list invalidations: %v\n\n", resp)
	return nil
}

func createInvalidationRequest(cli *cloudfront.Client, pattern string) error {
	now := time.Now()
	var quantity int32 = 1
	resp, err := cli.CreateInvalidation(context.TODO(), &cloudfront.CreateInvalidationInput{
		DistributionId: aws.String(os.Getenv("CLOUD_FRONT_DISTRIBUTION")),
		InvalidationBatch: &types.InvalidationBatch{
			CallerReference: aws.String(fmt.Sprintf("goinvali%s", now.Format("2006/01/02,15:04:05"))),
			Paths: &types.Paths{
				Quantity: &quantity,
				Items: []string{
					pattern,
				},
			},
		},
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("list create invalidation result: ")
	fmt.Printf("create invalidation result: %v\n\n", resp)
	return nil
}

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err.Error())
	}

	svc := cloudfront.New(cloudfront.Options{
		Credentials: cfg.Credentials,
		Region:      cfg.Region,
	})

	if err := viewListInvalidations(svc); err != nil {
		log.Fatal(err.Error())
	}

	pattern := "/t.png"
	if err := createInvalidationRequest(svc, pattern); err != nil {
		log.Fatal(err.Error())
	}

	if err := viewListInvalidations(svc); err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("all done")
}
