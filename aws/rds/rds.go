package rdsx

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/rds"
)

const Regin = "cn-northwest-1"

func ModifyPG() error {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return err
	}

	rdsSvc := rds.NewFromConfig(cfg)
	input := &rds.DescribeDBInstancesInput{
		DBInstanceIdentifier: aws.String("dev-tpp-aurora-postgres-instance-1"),
	}
	r, err := rdsSvc.DescribeDBInstances(context.TODO(), input)
	if err != nil {
		return err
	}
	fmt.Println(*r.DBInstances[0].DBInstanceClass)
	return nil
}
