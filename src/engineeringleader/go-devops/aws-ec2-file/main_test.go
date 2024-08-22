package main

import (
	"context"
	"io"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type MockS3Client struct {
	ListBucketsOutput  *s3.ListBucketsOutput
	CreateBucketOutput *s3.CreateBucketOutput
}

func (m MockS3Client) ListBuckets(ctx context.Context, params *s3.ListBucketsInput, optFns ...func(*s3.Options)) (*s3.ListBucketsOutput, error) {
	return m.ListBucketsOutput, nil
}

func (m MockS3Client) CreateBucket(ctx context.Context, params *s3.CreateBucketInput, optFns ...func(*s3.Options)) (*s3.CreateBucketOutput, error) {
	return m.CreateBucketOutput, nil
}

type MockS3Uploader struct {
	UploadOutput *manager.UploadOutput
}

type MockS3Downloader struct {
	DownloadOutput *manager.Downloader
}

func (m MockS3Uploader) Upload(ctx context.Context, input *s3.PutObjectInput, opts ...func(*manager.Uploader)) (*manager.UploadOutput, error) {
	return m.UploadOutput, nil
}

func (m MockS3Downloader) Download(ctx context.Context, w io.WriterAt, input *s3.GetObjectInput, options ...func(*manager.Downloader)) (n int64, err error) {
	return m.DownloadOutput.PartSize, nil
}

func TestCreateS3Bucket(t *testing.T) {
	ctx := context.Background()
	err := createS3Bucket(ctx, MockS3Client{
		ListBucketsOutput: &s3.ListBucketsOutput{
			Buckets: []types.Bucket{
				{
					Name: aws.String("test-bucket"),
				},
				{
					Name: aws.String("test-bucket-2"),
				},
			},
		},
		CreateBucketOutput: &s3.CreateBucketOutput{},
	})
	if err != nil {
		t.Fatalf("createS3Bucket error: %s", err)
	}
}

func TestUploadToS3Bucket(t *testing.T) {
	mockUploader := MockS3Uploader{
		UploadOutput: &manager.UploadOutput{},
	}
	err := uploadToS3Bucket(context.Background(), mockUploader, "testdata/test.txt")
	if err != nil {
		t.Fatalf("upload to bucket error: %s", err)
	}
}

func TestDownloadFromS3Bucket(t *testing.T) {
	mockDownloader := MockS3Downloader{
		DownloadOutput: &manager.Downloader{},
	}
	_, err := downloadFromS3Bucket(context.Background(), mockDownloader)
	if err != nil {
		t.Fatalf("download from bucket error: %s", err)
	}
}
