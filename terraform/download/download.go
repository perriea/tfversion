package tfdownload

import (
  "net/http"
  "io"
  "os"
  "runtime"
  "fmt"

  "github.com/fatih/color"
)

var (
  url_tf string
  version string
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func Run() {

  fatal := color.New(color.FgRed, color.Bold)

  // Verify number parameters
  if len(os.Args) == 2 {
    
      // Formulation URL Terraform Website
      fmt.Printf("Attempting to download version: %s\n", os.Args[1])
      url_tf := "https://releases.hashicorp.com/terraform/" + os.Args[1] + "/terraform_" + os.Args[1] + "_" + runtime.GOOS + "_" + runtime.GOARCH + ".zip"
      resp, err := http.Get(url_tf)
      check(err)
      defer resp.Body.Close()

      // Verify code equal 200
      if (err == nil) && (resp.StatusCode == 200) {
          out, err := os.Create("/tmp/terraform-" + os.Args[1] + ".zip")
          check(err)
          defer out.Close()
          _, err = io.Copy(out, resp.Body)
          check(err)
      } else {
          fatal.Printf("Download impossible, this version %s doesn't exist !", os.Args[1])
          os.Exit(1)
      }
  } else {
      fatal.Println("Error, too many argument !")
      os.Exit(1)
  }
}
