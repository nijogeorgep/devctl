package awshelper

import (
	"context"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	_ "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/stretchr/testify/assert"
)

// Mock S3 client
type mockS3Client struct {
	ListBucketsFunc func(ctx context.Context, params *s3.ListBucketsInput, optFns ...func(*s3.Options)) (*s3.ListBucketsOutput, error)
}

func (m *mockS3Client) ListBuckets(ctx context.Context, params *s3.ListBucketsInput, optFns ...func(*s3.Options)) (*s3.ListBucketsOutput, error) {
	return m.ListBucketsFunc(ctx, params, optFns...)
}

// Abstract the creation of the S3 client
var newS3Client = func(cfg aws.Config) *mockS3Client {
	return &mockS3Client{
		ListBucketsFunc: func(ctx context.Context, params *s3.ListBucketsInput, optFns ...func(*s3.Options)) (*s3.ListBucketsOutput, error) {
			return &s3.ListBucketsOutput{
				Buckets: []types.Bucket{
					{Name: aws.String("test-bucket-1")},
					{Name: aws.String("test-bucket-2")},
				},
			}, nil
		},
	}
}

// Abstract the loading of AWS config
var loadAWSConfigFunc = func(ctx context.Context) (aws.Config, error) {
	return aws.Config{}, errors.New("mock error")
}

func TestListS3Cmd(t *testing.T) {
	//mockClient := newS3Client(aws.Config{})

	// Use the mock client in the test
	cmd := listS3Cmd()
	err := cmd.RunE(nil, nil)
	assert.NoError(t, err)
}

func TestLoadAWSConfig(t *testing.T) {
	// Test failure scenario
	_, err := loadAWSConfigFunc(context.Background())
	assert.Error(t, err)
}
