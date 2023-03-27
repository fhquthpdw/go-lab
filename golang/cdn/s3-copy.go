// using this
package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const (
	SourceBucketName = "s3-cl"
	TargetBucketName = "s3-cl"
	S3Key            = "t.png"
	SourceVersionId  = ""
)

func main() {
	keyId := os.Getenv("ALIYUN_ACCESS_KEY_ID")
	keySecret := os.Getenv("ALIYUN_ACCESS_KEY_SECRET")
	if keyId == "" || keySecret == "" {
		log.Fatalf("set aliyun access credentials first")
	}

	S3Copy()
}

func S3Copy() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("load aws config err: %s", err.Error())
	}

	cli := s3.New(s3.Options{
		Credentials: cfg.Credentials,
		Region:      cfg.Region,
	})

	if _, err = cli.CopyObject(context.TODO(), &s3.CopyObjectInput{
		Bucket: aws.String(TargetBucketName), // target bucket
		Key:    aws.String(S3Key),            // target object key

		CopySource: aws.String(SourceBucketName + "/" + S3Key + "?versionId=" + SourceVersionId), // source object bucket and key
	}); err != nil {
		log.Fatalf("copy s3 file error: %s", err.Error())
	}
	log.Println("copy s3 file to Successfully!")
}
