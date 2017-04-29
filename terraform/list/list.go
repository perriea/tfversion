package tflist

import (
    "fmt"
    "net/http"
    "regexp"
    "bytes"

    "github.com/fatih/color"
)

var (
    url_tf    string
    cleaned   []string
    available []string
    tfversion []string
    good      *color.Color
)

func init()  {
    good = color.New(color.FgGreen, color.Bold)
    url_tf = "https://releases.hashicorp.com/terraform/"
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func stringInSlice(str string, list []string) bool {
    for _, v := range list {
        if v == str {
            return true
        }
    }

    return false
}

func Run()  {

    resp, err := http.Get(url_tf)
    check(err)
    defer resp.Body.Close()

    // Verify code equal 200
    if (err == nil) && (resp.StatusCode == 200) {
        r, err := regexp.Compile("[0-9]+\\.[0-9]+\\.[0-9]+(-(rc|beta)[0-9]+)?")
        check(err)

        // Convert byte to string
        buf := new(bytes.Buffer)
    	  buf.ReadFrom(resp.Body)
    	  newStr := buf.String()

        good.Printf("Versions availables of terraform (tfversion support <= 0.7) :\n")
        tfversion = r.FindAllString(newStr, -1)

        // Clean doublon
       	for _, value := range tfversion {
       	    if !stringInSlice(value, cleaned) {
         		     cleaned = append(cleaned, value)
       	    }
       	}

        // Show versions
     	  fmt.Println(cleaned)
    }
}
