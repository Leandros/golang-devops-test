package main

import (
	"devops.test/cli/cmd"
)

func main() {
	cmd.Execute()
}

// Doxy.me Senior DevOps Coding Challenge:
//
// You're looking at a Go CLI application to interact with AWS.
//
// It's using https://cobra.dev to parse and process command line arguments.
// Furthermore, the application depends on the AWS Golang SDK. For any
// further dependencies check the `go.mod`.
//
//
// Task:
//
// Your task is to add a new sub-command called 'upload'. This command must
// take a single parameter, the path to a file (e.g., `cli upload <path-to-file>`).
// The file specified must be uploaded to an AWS S3 bucket, using the already
// provided interface in `provider/s3.go`.
// The file must be renamed and uploaded to S3 as `<unixtimestamp>.<original-file-ending>`.
//
// ** Please implement the subcommand and the required tests. **
//
//
// Further information:
//
// The S3 interface has two implementations, the first implementation (`provider/s3.go`)
// is using the aws-sdk, the second implementation (`provider/s3_mock.go`)
// is a mock implementation to be used for testing. It's guaranteed to make
// no network requests.
