package provider

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type MockS3API struct {
	Buckets map[string]map[string]*[]byte
}

func NewClient() MockS3API {
	return MockS3API{
		Buckets: map[string]map[string]*[]byte{},
	}
}

func (p *MockS3API) PutObject(
	ctx context.Context,
	params *s3.PutObjectInput,
	optFns ...func(*s3.Options),
) (*s3.PutObjectOutput, error) {
	files := p.Buckets[*params.Bucket]
	if files == nil {
		p.Buckets[*params.Bucket] = map[string]*[]byte{}
		files = p.Buckets[*params.Bucket]
	}

	file, err := ioutil.ReadAll(params.Body)
	if err != nil {
		return nil, fmt.Errorf("reading body: %w", err)
	}

	files[*params.Key] = &file

	return &s3.PutObjectOutput{}, nil
}

func (p *MockS3API) GetObject(
	ctx context.Context,
	params *s3.GetObjectInput,
	optFns ...func(*s3.Options),
) (*s3.GetObjectOutput, error) {
	files := p.Buckets[*params.Bucket]
	if files == nil {
		return nil, fmt.Errorf("bucket not found: %s", *params.Bucket)
	}

	file := files[*params.Key]
	if file == nil {
		return nil, fmt.Errorf("file not found: %s", *params.Key)
	}

	return &s3.GetObjectOutput{
		Body: io.NopCloser(bytes.NewReader(*file)),
	}, nil
}

func (p *MockS3API) DeleteObject(
	ctx context.Context,
	params *s3.DeleteObjectInput,
	optFns ...func(*s3.Options),
) (*s3.DeleteObjectOutput, error) {
	files := p.Buckets[*params.Bucket]
	if files == nil {
		return nil, fmt.Errorf("bucket not found: %s", *params.Bucket)
	}

	delete(files, *params.Key)

	return &s3.DeleteObjectOutput{}, nil
}
