package tfgcp

import (
	"fmt"
	"os"

	"cloud.google.com/go/storage"
	"golang.org/x/net/context"
)

var (
	ctx        context.Context
	bucketName string
	err        error
)

func init() {
	// The SDK has support for the shared configuration file, please export GOOGLE_APPLICATION_CREDENTIALS
	ctx = context.Background()
	bucketName = "tfversion"
}

// Run : Lauch test GCP
func Run(projectID string) {

	var (
		client *storage.Client
		bucket *storage.BucketHandle
	)

	fmt.Print("\033[1;37mGoogle Cloud Plateform Connection (Storage) :\n")

	// Creates a client.
	client, err = storage.NewClient(ctx)
	if err != nil {
		fmt.Printf("\033[0;31m- Failed to create client: %v\n", err)
		os.Exit(0)
	}

	fmt.Print("\033[0;32m- Successful connection\n")

	// Creates a Bucket instance.
	bucket = client.Bucket(bucketName)

	// Creates the new bucket.
	if err = bucket.Create(ctx, projectID, nil); err != nil {
		fmt.Printf("\033[0;31m- Failed to create bucket: %v", err)
	} else {
		fmt.Printf("\033[0;32m- Bucket %v created.\n", bucketName)

		if err = bucket.Delete(ctx); err != nil {
			fmt.Printf("\033[0;31m- Failed to delete bucket: %v", err)
		} else {
			fmt.Printf("\033[0;32m- Bucket %v deleted.\n", bucketName)
		}
	}
}
