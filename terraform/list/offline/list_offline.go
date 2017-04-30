package tflist_offline

import (
    "regexp"
    "os"
    "os/user"
    "io/ioutil"

    "github.com/perriea/tfversion/error"
)

var (
    path      string
    tfversion string
    count     int
    usr       *user.User
    err       error
)

func init()  {
    usr, err = user.Current()
    tferror.Panic(err)

    count = 0
    path = usr.HomeDir + "/terraform/tmp/"
}

func Run()  {

    r, err := regexp.Compile("[0-9]+\\.[0-9]+\\.[0-9]+(-(rc|beta)[0-9]+)?")
    tferror.Panic(err)

    tferror.Run(0, "[INFO] All local version:")

    files, err := ioutil.ReadDir(path)
    tferror.Panic(err)
    for _, f := range files {
        tfversion = r.FindString(f.Name())
        tferror.Run(-1, tfversion)

        count++
    }

    if count == 0 {
        tferror.Run(2, "No local versions !")
    }
}

func Cleanup()  {

    files, err := ioutil.ReadDir(path)
    tferror.Panic(err)
    for _, f := range files {
        err = os.Remove(path + f.Name())
        tferror.Panic(err)

        count++
    }

    if count == 0 {
        tferror.Run(2, "Nothing deleted !")
    } else {
        tferror.Run(1, "All files are deleted !")
    }
}