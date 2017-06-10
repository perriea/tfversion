package main

import (
	"os"

	"github.com/mkideal/cli"
	"github.com/perriea/tfversion/error"
)

var (
	err error
)

func main() {

	err = cli.Root(root,
		cli.Tree(help),
		cli.Tree(install),
		cli.Tree(uninstall),
		cli.Tree(list),
		cli.Tree(test),
	).Run(os.Args[1:])

	tferror.Panic(err)
}
