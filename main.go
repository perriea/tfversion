package main

import (
	"fmt"
	"os"

	"github.com/mkideal/cli"
)

var (
	err error
)

func main() {

	err = cli.Root(root,
		cli.Tree(install),
		cli.Tree(uninstall),
		cli.Tree(list),
		cli.Tree(test),
	).Run(os.Args[1:])

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
