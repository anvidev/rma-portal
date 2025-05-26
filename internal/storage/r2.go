package storage

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const (
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

func (s *R2Storage) Get() {}

func (s *R2Storage) Put(ctx context.Context, body io.Reader, namespace, filename string) (*StorageResponse, error) {
	key := aws.String(fmt.Sprintf("%s/%s/%s", folderName, namespace, timestampedKey(filename)))

	_, err := s.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: &s.bucketName,
		Key:    key,
		Body:   body,
	})
	if err != nil {
		return nil, err
	}

	return &StorageResponse{
		Key: *key,
		URL: url(s.accountID, s.bucketName, *key),
	}, nil
}

func timestampedKey(s string) string {
	return fmt.Sprintf("%d-%s", time.Now().Unix(), s)
}

func url(accountID, bucketName, key string) string {
	return fmt.Sprintf("https://%s.r2.cloudflarestorage.com/%s/%s", accountID, bucketName, key)
}
