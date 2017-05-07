package tfcloud

import (
	"flag"
	"fmt"

	"github.com/perriea/tfversion/cloud/aws"
)

var (
	testaws   bool
	testgcp   bool
	testazure bool
	err       error
)

func Run(params []string) error {

	cloud := flag.NewFlagSet("cloud", flag.ExitOnError)
	cloud.BoolVar(&testaws, "aws", false, "Test connection on AWS)")
	cloud.BoolVar(&testgcp, "gcp", false, "Test connection on GCP)")
	cloud.BoolVar(&testazure, "azure", false, "Test connection on Azure)")
	cloud.Parse(params)

	if testaws && testgcp && testazure {
		return fmt.Errorf("--aws, --gcp and --azure are mutually exclusive")
	}

	if len(params) != 1 {
		return fmt.Errorf("Too many arguments ...")
	}

	if testaws {
		aws.TestConnect()
	} else if testgcp {
		return fmt.Errorf("GCP test is not actived")
	} else if testazure {
		return fmt.Errorf("Azure test is not actived")
	} else {
		return fmt.Errorf("This provider doesn't exist !")
	}

	return nil
}
