package aws

import (
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
	//log.Fatal(err)
	if err != nil {
		tferror.Run(2, "[WARN] Your AWS access is not correct")
		tferror.Panic(err)
	} else {
		tferror.Run(1, "Your AWS access is correct")
	}
}
