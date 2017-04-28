package main

import (
  "github.com/perriea/tfversion/terraform/download"
  "github.com/perriea/tfversion/terraform/install"
)

func main()  {
  tfdownload.Run()
  tfinstall.Run()
}
