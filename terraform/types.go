package terraform

import "net/http"

// Release struct : information switch release
type Release struct {
	Home         string
	Version      string
	HTTPclient   *http.Client
	HTTPResponse *http.Response
}

type Path string

const (
	PathTerraform      Path = "https://releases.hashicorp.com/terraform/%s/terraform_%s_%s_%s.zip"
	PathTerraformIndex Path = "https://releases.hashicorp.com/terraform/"
	PathBin            Path = "/.tfversion/bin/"
	PathTmp            Path = "/.tfversion/tmp/"
)

func (f Path) toString() string {
	return string(f)
}
