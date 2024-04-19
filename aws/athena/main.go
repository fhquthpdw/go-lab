package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/athena"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/aws/smithy-go/logging"
	"time"
)

// SESA uat account for test
func main() {
	region := "cn-northwest-1"
	var logger logging.Logger
	var optFns []func(*config.LoadOptions) error

	optFns = append(optFns, config.WithRegion(region), config.WithLogger(logger))
	awsCfg, err := config.LoadDefaultConfig(context.TODO(), optFns...)
	if err != nil {
		panic("failed to load default config: " + err.Error())
	}

	//
	accountId := "248598923617"
	stsRoleName := "CrossAccountTerraformAccessRole"
	stsSvc := sts.NewFromConfig(awsCfg)
	stsRole := fmt.Sprintf("arn:aws-cn:iam::%s:role/%s", accountId, stsRoleName)
	credential := stscreds.NewAssumeRoleProvider(stsSvc, stsRole)

	awsCfg.Credentials = aws.NewCredentialsCache(credential)

	// Create service client value configured for credentials, from assumed role.
	athenaClient := athena.NewFromConfig(awsCfg)
	//

	// build query
	query := "select op, dmstimestamp, id, created_at, ext_id, corp_id, name, avatar from weiban.b_wecom_ppd_weiban_external_user where id > 0 limit 5"
	outputLocation := "s3://aws-athena-query-results-weiban/"
	input := &athena.StartQueryExecutionInput{
		QueryString:         aws.String(query),
		ResultConfiguration: &types.ResultConfiguration{OutputLocation: &outputLocation},
	}

	// execute query
	startResult, err := athenaClient.StartQueryExecution(context.Background(), input)
	if err != nil {
		panic("failed start query execution: " + err.Error())
	}
	queryExecutionID := startResult.QueryExecutionId

	var nextToken *string = nil
	// get query results in loop

	//OuterLoop:
	for {
		queryExecutionResult, err := athenaClient.GetQueryExecution(context.Background(), &athena.GetQueryExecutionInput{
			QueryExecutionId: queryExecutionID,
		})
		if err != nil {
			panic("failed to get query execution status, " + err.Error())
		}

		state := queryExecutionResult.QueryExecution.Status.State
		if state == "SUCCEEDED" {
			// break
		} else if state == "FAILED" || state == "CANCELLED" {
			panic("query execution failed or cancelled")
		} else {
			// sleep for a while before polling again
			fmt.Println("waiting 1 second ...")
			time.Sleep(time.Second * 1)
			continue
		}

		// get the query results
		result, err := athenaClient.GetQueryResults(context.Background(), &athena.GetQueryResultsInput{
			QueryExecutionId: queryExecutionID,
			MaxResults:       aws.Int32(2),
			NextToken:        nextToken,
		})
		if err != nil {
			panic("failed to get query results, " + err.Error())
		}

		// Print the query results
		for _, row := range result.ResultSet.Rows {
			for _, data := range row.Data {
				if data.VarCharValue != nil {
					//break OuterLoop
					fmt.Printf("%s\t", *data.VarCharValue)
				}
			}
			fmt.Println()
		}

		// 分页处理
		if result.NextToken == nil {
			break
		}
		nextToken = result.NextToken

		// fmt.Println("===========")
		// fmt.Println("Page: ", *nextToken)
		// fmt.Println("===========")
	}

	return
}
