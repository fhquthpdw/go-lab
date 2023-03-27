// using this
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	cdn20180510 "github.com/alibabacloud-go/cdn-20180510/v2/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const (
	FileName   = "t.png"
	BucketName = "s3-cl"
	ObjKey     = "t.png"
)

func main() {
	keyId := os.Getenv("ALIYUN_ACCESS_KEY_ID")
	keySecret := os.Getenv("ALIYUN_ACCESS_KEY_SECRET")
	if keyId == "" || keySecret == "" {
		log.Fatalf("set aliyun access credentials first")
	}

	S3Upload()
	RefreshCloudFront()
	RefreshAliCDN(keyId, keySecret)
}

func S3Upload() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("load aws config err: %s", err.Error())
	}

	cli := s3.New(s3.Options{
		Credentials: cfg.Credentials,
		Region:      cfg.Region,
	})

	file, err := os.Open(FileName)
	if err != nil {
		log.Fatalf("open file err: %s", err.Error())
	}
	defer func() {
		_ = file.Close()
	}()

	if _, err = cli.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(BucketName),
		Key:         aws.String(ObjKey),
		Body:        file,
		ContentType: aws.String("image/png"),
	}); err != nil {
		log.Fatalf("upload file to s3 error: %s", err.Error())
	}
	log.Println("upload file to s3 Successfully!")
}

func RefreshCloudFront() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("load aws default config err: %s", err.Error())
	}

	cli := cloudfront.New(cloudfront.Options{
		Credentials: cfg.Credentials,
		Region:      cfg.Region,
	})

	pattern := "/t.png"
	var quantity int32 = 1
	_, err = cli.CreateInvalidation(context.TODO(), &cloudfront.CreateInvalidationInput{
		DistributionId: aws.String(os.Getenv("CLOUD_FRONT_DISTRIBUTION")),
		InvalidationBatch: &types.InvalidationBatch{
			CallerReference: aws.String(fmt.Sprintf("goinvali%s", time.Now().Format("2006/01/02,15:04:05"))),
			Paths: &types.Paths{
				Quantity: &quantity,
				Items: []string{
					pattern,
				},
			},
		},
	})
	if err != nil {
		log.Fatal("create cloud front invalidation err: %s", err.Error())
	}
	log.Println("create cloud front invalidation Successfully")
}

func RefreshAliCDN(keyId, keySecret string) {
	cfg := &openapi.Config{
		AccessKeyId:     tea.String(keyId),
		AccessKeySecret: tea.String(keySecret),
	}
	cfg.Endpoint = tea.String("cdn.aliyuncs.com")

	cli, err := cdn20180510.NewClient(cfg)
	if err != nil {
		log.Fatalf("create ali yun client err: %s", err.Error())
	}

	refreshObjectCachesRequest := &cdn20180510.RefreshObjectCachesRequest{
		ObjectPath: tea.String("https://ali-cdn-poc.lego.cn/t.png"),
		ObjectType: tea.String("File"),
	}
	runtime := &util.RuntimeOptions{}

	_, err = cli.RefreshObjectCachesWithOptions(refreshObjectCachesRequest, runtime)
	if err != nil {
		log.Fatalf("refresh object in ali yun err: %s", err.Error())
	}
	log.Println("refresh object in ali yun Successfully")
}
