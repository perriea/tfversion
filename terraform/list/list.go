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
    tfversion []string
    fatal     *color.Color
)

func init()  {
    fatal = color.New(color.FgRed, color.Bold)
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

        buf := new(bytes.Buffer)
    	  buf.ReadFrom(resp.Body)
    	  newStr := buf.String()

        fmt.Printf("Versions availables :\n")
        tfversion = r.FindAllString(newStr, -1)

       	for _, value := range tfversion {

       		if !stringInSlice(value, cleaned) {
       			cleaned = append(cleaned, value)
       		}
       	}

     	  fmt.Println(cleaned)
    }
}
