package provider_test

import (
	"bytes"
	"context"
	"io/ioutil"

	"devops.test/cli/provider"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("S3 Mock API", func() {
	var mockClient provider.MockS3API
	BeforeEach(func() {
		mockClient = provider.NewClient()
	})

	Describe("PutObject", func() {
		var res *s3.PutObjectOutput
		var err error
		BeforeEach(func() {
			res, err = provider.PutObject(context.TODO(), &mockClient, &s3.PutObjectInput{
				Bucket: provider.Ptr("my:bucket:arn"),
				Key:    provider.Ptr("readme.txt"),
				Body:   bytes.NewReader([]byte("Hello World")),
			})
		})

		It("will not return an error", func() {
			Expect(err).To(BeNil())
			Expect(res).NotTo(BeNil())
		})

		It("will insert the file", func() {
			file := mockClient.Buckets["my:bucket:arn"]["readme.txt"]

			Expect(err).To(BeNil())
			Expect(file).ToNot(BeNil())
			Expect(*file).To(BeEquivalentTo([]byte("Hello World")))
		})
	})

	Describe("GetObject", func() {
		BeforeEach(func() {
			file := []byte("Hello World")
			mockClient.Buckets["my:bucket:arn"] = map[string]*[]byte{}
			mockClient.Buckets["my:bucket:arn"]["readme.txt"] = &file
		})

		It("will return the object", func() {
			res, err := provider.GetObject(context.TODO(), &mockClient, &s3.GetObjectInput{
				Bucket: provider.Ptr("my:bucket:arn"),
				Key:    provider.Ptr("readme.txt"),
			})

			Expect(err).To(BeNil())
			Expect(res).NotTo(BeNil())

			file, err := ioutil.ReadAll(res.Body)
			Expect(err).To(BeNil())
			Expect(file).To(BeEquivalentTo([]byte("Hello World")))
		})
	})

	Describe("DeleteObject", func() {
		BeforeEach(func() {
			file := []byte("Hello World")
			mockClient.Buckets["my:bucket:arn"] = map[string]*[]byte{}
			mockClient.Buckets["my:bucket:arn"]["readme.txt"] = &file
		})

		It("will delete the object", func() {
			res, err := provider.DeleteObject(context.TODO(), &mockClient, &s3.DeleteObjectInput{
				Bucket: provider.Ptr("my:bucket:arn"),
				Key:    provider.Ptr("readme.txt"),
			})
			Expect(err).To(BeNil())
			Expect(res).NotTo(BeNil())
			Expect(mockClient.Buckets["my:bucket:arn"]["readme.txt"]).To(BeNil())
		})
	})
})
