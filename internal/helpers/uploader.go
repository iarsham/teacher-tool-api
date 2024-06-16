package helpers

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/iarsham/teacher-tool-api/configs"
	"io"
	"mime/multipart"
	"strings"
	"time"
)

func UploadAwsS3(cfg *configs.Config, file multipart.File, folder, fileName string) (string, error) {
	s3Config, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(cfg.S3.Region))
	if err != nil {
		return "", err
	}
	s3Config.Credentials = aws.CredentialsProviderFunc(func(ctx context.Context) (aws.Credentials, error) {
		return aws.Credentials{
			AccessKeyID:     cfg.S3.AccessKey,
			SecretAccessKey: cfg.S3.SecretKey,
		}, nil
	})
	s3Config.BaseEndpoint = aws.String(cfg.S3.Endpoint)
	client := s3.NewFromConfig(s3Config)
	dst := DstPath(folder, fileName)
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(cfg.S3.BucketName),
		Key:    aws.String(dst),
		Body:   bytes.NewReader(fileBytes),
	})
	if err != nil {
		return "", err
	}
	return CreateS3Url(cfg, dst), nil
}

func CreateS3Url(cfg *configs.Config, fileName string) string {
	return cfg.S3.StorageDomain + "/" + fileName
}

func DstPath(folder, fileName string) string {
	name := strings.Split(fileName, ".")
	timeNow := time.Now().Format("2006-01-02")
	return folder + "/" + strings.ReplaceAll(name[0], " ", "") + "-" + timeNow + "." + name[1]
}
