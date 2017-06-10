package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/mkideal/cli"
)

var (
	bold *color.Color
	err  error
)

func init() {
	bold = color.New(color.FgWhite, color.Bold)
}

func main() {

	err = cli.Root(root,
		cli.Tree(help),
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
