// using this
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const (
	VersionListBucketName = "s3-cl"
	VersionListS3Key      = "t.png"
)

func main() {
	keyId := os.Getenv("ALIYUN_ACCESS_KEY_ID")
	keySecret := os.Getenv("ALIYUN_ACCESS_KEY_SECRET")
	if keyId == "" || keySecret == "" {
		log.Fatalf("set aliyun access credentials first")
	}

	S3ListVersions()
}

func S3ListVersions() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("load aws config err: %s", err.Error())
	}

	cli := s3.New(s3.Options{
		Credentials: cfg.Credentials,
		Region:      cfg.Region,
	})

	out, err := cli.ListObjectVersions(context.TODO(), &s3.ListObjectVersionsInput{
		Bucket: aws.String(VersionListBucketName),
		Prefix: aws.String(VersionListS3Key),
	})
	versionList := out.Versions
	fmt.Println(len(versionList))
	for _, v := range versionList {
		fmt.Printf("Version: %s, Time: %s\r\n", *v.VersionId, v.LastModified)
	}
}
