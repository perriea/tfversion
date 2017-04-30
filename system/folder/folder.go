package tffolder

import (
  "os"
  "os/user"

  "github.com/perriea/tfversion/error"
)

var (
    usr *user.User
    err error
)

func init()  {
    usr, err = user.Current()
    tferror.Panic(err)
}

func Run(name string, chmod int) {
    err = os.MkdirAll(usr.HomeDir + name, os.FileMode(chmod))
    tferror.Panic(err)
}
