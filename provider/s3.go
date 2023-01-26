package provider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// S3API is the interface for the AWS S3 API.
//
// Usage:
// 		The s3 client from the AWS SDK adheres to this interface. This will
// 		allow us to dynamically replace it with a mock implementation.
//
// Example:
// 		client := s3.newFromConfig(config)
// 		out, err := PutObject(context.TODO(), &s3.PutObjectInput { ... })
//		if err != nil {
//			...
// 		}
type S3API interface {
	PutObject(
		ctx context.Context,
		params *s3.PutObjectInput,
		optFns ...func(*s3.Options),
	) (*s3.PutObjectOutput, error)
	GetObject(
		ctx context.Context,
		params *s3.GetObjectInput,
		optFns ...func(*s3.Options),
	) (*s3.GetObjectOutput, error)
	DeleteObject(
		ctx context.Context,
		params *s3.DeleteObjectInput,
		optFns ...func(*s3.Options),
	) (*s3.DeleteObjectOutput, error)
}

// PutObject wraps the `put` S3 API call.
func PutObject(
	ctx context.Context,
	api S3API,
	params *s3.PutObjectInput,
) (*s3.PutObjectOutput, error) {
	return api.PutObject(ctx, params)
}

// GetObject wraps the `get` S3 API call.
func GetObject(
	ctx context.Context,
	api S3API,
	params *s3.GetObjectInput,
) (*s3.GetObjectOutput, error) {
	return api.GetObject(ctx, params)
}

// GetObject wraps the `delete` S3 API call.
func DeleteObject(
	ctx context.Context,
	api S3API,
	params *s3.DeleteObjectInput,
) (*s3.DeleteObjectOutput, error) {
	return api.DeleteObject(ctx, params)
}
