package tfdownload

import (
    "net/http"
    "io"
    "os"
    "runtime"
    "fmt"
    "regexp"
    "crypto/tls"

    "github.com/perriea/tfversion/error"
)

var (
    // URL and Path (file)
    do_url_tf       string
    url_tf          string
    do_path_tf      string
    path_tf         string
    file_unzip      *os.File
    // check if format version is *.*.* and more ...
    match           bool
    // HTTP request
    transport       *http.Transport
    client          *http.Client
    resp            *http.Response
    // Errors
    err_fatal_msg   string
    err_msg         string
    err             error
)

func init()  {
    // Dont check certificate SSL + new path
    transport = &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
    client      = &http.Client{Transport: transport}
    do_url_tf   = "https://releases.hashicorp.com/terraform/%s/terraform_%s_%s_%s.zip"
    do_path_tf  = "/tmp/terraform-%s.zip"

    // Color response
    err_fatal_msg = "[FATAL] Download impossible, this version doesn't exist !"
    err_msg       = "[ERROR] This version (%s) is not supported !"
}

func Run(version string) bool {

    match, err := regexp.MatchString("[0-9]+\\.[0-6]+\\.[0-9]+(-(rc|beta)[0-9]+)?", version)
    tferror.Panic(err)

    if !match {
        // Formulation URL Terraform Website
        fmt.Printf("Attempting to download version: %s\n", version)
        url_tf = fmt.Sprintf(do_url_tf, version, version, runtime.GOOS, runtime.GOARCH)

        // Request GET URL
        resp, err = client.Get(url_tf)
        tferror.Panic(err)
        defer resp.Body.Close()

        // Verify code equal 200
        if (err == nil) && (resp.StatusCode == 200) {
            tferror.Run(1, "Start download ...")
            path_tf = fmt.Sprintf(do_path_tf, version)
            file_unzip, err = os.Create(path_tf)
            tferror.Panic(err)
            defer file_unzip.Close()

            // Copy reponse in file
            _, err = io.Copy(file_unzip, resp.Body)
            tferror.Panic(err)

            return true

        } else {
            tferror.Run(3, err_fatal_msg)
            return false
        }

    } else {
        tferror.Run(3, err_msg)
        return false
    }

    return false
}
