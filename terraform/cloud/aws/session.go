package tfaws

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/perriea/tfversion/error"
)

var (
	sess   *session.Session
	ec2svc *ec2.EC2
	err    error
)

func TestConnect() {

	// The SDK has support for the shared configuration file (~/.aws/config)
	// Create a session to share configuration, and load external configuration.
	sess = session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	params := &ec2.DescribeInstancesInput{}

	// Create the service's client with the session.
	ec2svc = ec2.New(sess)
	_, err := ec2svc.DescribeInstances(params)

	if err != nil {
		tferror.Run(2, "[WARN] Your AWS access is not correct")
		if awsErr, ok := err.(awserr.Error); ok {

			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Printf("%s : %s (%s)", awsErr.Code(), awsErr.Message(), reqErr.RequestID())
			} else {
				// Generic AWS Error with Code, Message, and original error (if any)
				fmt.Printf("%s : %s\n%s", awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			}

		} else {
			tferror.Panic(err)
		}

	} else {
		tferror.Run(1, "Your AWS access is correct")
	}
}
