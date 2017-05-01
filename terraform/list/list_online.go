package tflist

import (
    "fmt"
    "net/http"
    "crypto/tls"
    "regexp"
    "bytes"

    "github.com/perriea/tfversion/error"
)

var (
    url_tf    string
    cleaned   []string
    available []string
    tfversions []string
    transport *http.Transport
    client    *http.Client
)

func init()  {
    transport = &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
    client = &http.Client{Transport: transport}

    url_tf = "https://releases.hashicorp.com/terraform/"
}

func stringInSlice(str string, list []string) bool {

    for _, v := range list {
        if v == str {
            return true
        }
    }

    return false
}

func ListOff()  {

    resp, err := client.Get(url_tf)
    tferror.Panic(err)
    defer resp.Body.Close()

    // Verify code equal 200
    if (err == nil) && (resp.StatusCode == 200) {

        r, err := regexp.Compile("[0-9]+\\.[0-9]+\\.[0-9]+(-(rc|beta)[0-9]+)?")
        tferror.Panic(err)

        // Convert byte to string
        buf := new(bytes.Buffer)
    	  buf.ReadFrom(resp.Body)
    	  newStr := buf.String()

        tferror.Run(1, "Versions availables of terraform (tfversion support <= 0.7) :")
        tfversions = r.FindAllString(newStr, -1)

        // Clean doublon
       	for _, value := range tfversions {
       	    if !stringInSlice(value, cleaned) {
         		     cleaned = append(cleaned, value)
       	    }
       	}

        // Show versions
     	  fmt.Println(cleaned)
    }
}
