package tfaws

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	sess       *session.Session
	s3svc      *s3.S3
	bucketName string
	err        error
)

func init() {
	// The SDK has support for the shared configuration file (~/.aws/config)
	// Create a session to share configuration, and load external configuration.
	sess = session.Must(session.NewSessionWithOptions(session.Options{SharedConfigState: session.SharedConfigEnable}))
	bucketName = "tfversion"
}

// Run : Lauch test AWS
func Run() {

	fmt.Print("\033[1;37mAWS Connection (S3) :\n")

	// sess with s3
	s3svc = s3.New(sess)

	// creation bucket
	_, err := s3svc.CreateBucket(&s3.CreateBucketInput{Bucket: &bucketName})
	if err != nil {
		fmt.Printf("\033[0;31m- Failed to create client\n")
		fmt.Printf("\033[0;31m- Failed to create bucket: %v", err)
	} else {
		fmt.Print("\033[0;32m- Successful connection\n")
		fmt.Printf("\033[0;32m- Bucket %v created.\n", bucketName)

		// delete bucket
		if _, err := s3svc.DeleteBucket(&s3.DeleteBucketInput{Bucket: &bucketName}); err != nil {
			fmt.Printf("\033[0;31m- Failed to delete bucket: %v", err)
		} else {
			fmt.Printf("\033[0;32m- Bucket %v deleted.\n", bucketName)
		}
	}
}
