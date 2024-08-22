package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

const bucketName = "aws-demo-test-bucket-12ab"
const region = "eu-west-1"

type S3Client interface {
	ListBuckets(ctx context.Context, params *s3.ListBucketsInput, optFns ...func(*s3.Options)) (*s3.ListBucketsOutput, error)
	CreateBucket(ctx context.Context, params *s3.CreateBucketInput, optFns ...func(*s3.Options)) (*s3.CreateBucketOutput, error)
}

type S3Uploader interface {
	Upload(ctx context.Context, input *s3.PutObjectInput, opts ...func(*manager.Uploader)) (*manager.UploadOutput, error)
}

type S3Downloader interface {
	Download(ctx context.Context, w io.WriterAt, input *s3.GetObjectInput, options ...func(*manager.Downloader)) (n int64, err error)
}

func main() {
	var (
		s3Client *s3.Client
		err      error
	)

	ctx := context.Background()

	if s3Client, err = initS3Client(ctx, region); err != nil {
		fmt.Printf("unable to get S3 client, %s", err)
		os.Exit(1)
	}

	if err = createS3Bucket(ctx, s3Client); err != nil {
		fmt.Printf("unable to create S3 bucket, %s", err)
		os.Exit(1)
	}

	if err = uploadToS3Bucket(ctx, manager.NewUploader(s3Client), "test.txt"); err != nil {
		fmt.Printf("upload to S3 bucket error, %s", err)
		os.Exit(1)
	}

	fmt.Println("Upload complete.")

	out, err := downloadFromS3Bucket(ctx, manager.NewDownloader(s3Client))
	if err != nil {
		fmt.Printf("download from S3 bucket error, %s", err)
		os.Exit(1)
	}

	fmt.Printf("Output: %s", out)

}

func initS3Client(ctx context.Context, region string) (*s3.Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		return nil, fmt.Errorf("unable to load SDK config, %s", err)
	}

	return s3.NewFromConfig(cfg), nil
}

func createS3Bucket(ctx context.Context, s3Client S3Client) error {

	allBuckets, err := s3Client.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		return fmt.Errorf("no S3 buckets available, %s", err)
	}

	found := false
	for _, bucket := range allBuckets.Buckets {
		if *bucket.Name == bucketName {
			found = true
		}
	}

	if !found {
		_, err = s3Client.CreateBucket(ctx, &s3.CreateBucketInput{
			Bucket: aws.String(bucketName),
			CreateBucketConfiguration: &types.CreateBucketConfiguration{
				LocationConstraint: types.BucketLocationConstraint(region),
			},
		})

		if err != nil {
			return fmt.Errorf("unable to create S3 bucket, %s", err)
		}
	}

	return nil
}

func uploadToS3Bucket(ctx context.Context, uploader S3Uploader, fileName string) error {
	testFile, err := os.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("read file error: %s", err)
	}

	_, err = uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Body:   bytes.NewReader(testFile),
		Key:    aws.String(fileName),
	})
	if err != nil {
		return fmt.Errorf("upload error to S3 bucket, %s", err)
	}

	return nil
}

func downloadFromS3Bucket(ctx context.Context, downloader S3Downloader) ([]byte, error) {

	buffer := manager.NewWriteAtBuffer([]byte{})

	numBytes, err := downloader.Download(ctx, buffer, &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String("test.txt"),
	})

	if err != nil {
		return nil, fmt.Errorf("download error to S3 bucket, %s", err)
	}

	if numBytesReceived := len(buffer.Bytes()); numBytes != int64(numBytesReceived) {
		return nil, fmt.Errorf("numbytes received doesn't match: %d : %d", numBytes, numBytesReceived)
	}

	return buffer.Bytes(), nil
}
