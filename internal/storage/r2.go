package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const (
	BucketUrl string = "https://anvi.dev"

	storageContextTimeout   time.Duration = time.Second * 20
	storagePresignExpiresIn time.Duration = time.Minute * 3

	folderName string = "files"
)

type R2Storage struct {
	bucketName string
	accountID  string
	client     *s3.Client
}

func NewR2Storage(bucketName, accountID, accessKeyID, accessKeySecret string) (Storager, error) {
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyID, accessKeySecret, "")),
		config.WithRegion("auto"),
	)
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(fmt.Sprintf("https://%s.r2.cloudflarestorage.com", accountID))
	})

	return &R2Storage{
		bucketName: bucketName,
		accountID:  accountID,
		client:     client,
	}, nil
}

func (s *R2Storage) PutPresign(ctx context.Context, key, contentType string, contentLength int64) (string, error) {
	presignClient := s3.NewPresignClient(s.client)

	ctx, cancel := context.WithTimeout(ctx, storageContextTimeout)
	defer cancel()

	obj := &s3.PutObjectInput{
		Bucket:        &s.bucketName,
		Key:           &key,
		ContentType:   &contentType,
		ContentLength: &contentLength,
	}

	request, err := presignClient.PresignPutObject(ctx, obj, s3.WithPresignExpires(storagePresignExpiresIn))
	if err != nil {
		return "", err
	}

	return request.URL, nil
}
