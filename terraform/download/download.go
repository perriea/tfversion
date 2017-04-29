package tfdownload

import (
    "net/http"
    "io"
    "os"
    "runtime"
    "fmt"
    "regexp"

    "github.com/fatih/color"
)

var (
    url_tf    string
    version   string
    denied    []string
    fatal     *color.Color
    good      *color.Color
)

func init()  {
    fatal = color.New(color.FgRed, color.Bold)
    good = color.New(color.FgGreen, color.Bold)
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func Run(version string) bool {
    match, err := regexp.MatchString("[0-9]+\\.[0-6]+\\.[0-9]+(-(rc|beta)[0-9]+)?", version)
    check(err)

    if !match {
        // Formulation URL Terraform Website
        fmt.Printf("Attempting to download version: %s\n", version)
        url_tf = "https://releases.hashicorp.com/terraform/" + version + "/terraform_" + version + "_" + runtime.GOOS + "_" + runtime.GOARCH + ".zip"
        resp, err := http.Get(url_tf)
        check(err)
        defer resp.Body.Close()

        // Verify code equal 200
        if (err == nil) && (resp.StatusCode == 200) {
            good.Printf("Start download ...\n")
            out, err := os.Create("/tmp/terraform-" + version + ".zip")
            check(err)
            defer out.Close()

            _, err = io.Copy(out, resp.Body)
            check(err)

            return true
        } else {
            fatal.Printf("[FATAL] Download impossible, this version doesn't exist !")
            return false
        }
    } else {
        fatal.Printf("[ERROR] This version (%s) is not supported !", version)
        return false
    }

    return false
}
