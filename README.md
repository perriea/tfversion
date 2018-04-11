# tfversion
[![Go Report Card](https://goreportcard.com/badge/github.com/perriea/tfversion)](https://goreportcard.com/report/github.com/perriea/tfversion) [![Build Status](https://travis-ci.org/perriea/tfversion.svg?branch=master)](https://travis-ci.org/perriea/tfversion) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

`tfversion` is a command created to switch between different versions of [Terraform](https://www.terraform.io).

## Build Project

- [Install Golang](https://golang.org/doc/install) (add var ENV),
- Build with commands `go build`, `make` or `go get -u github.com/perriea/tfversion`,
- Add in your `.bashrc` (Linux), `.bash_profile` (Darwin, MacOS) or `.zshrc` : `export PATH=$PATH:$HOME/.tfversion/bin`.

## Commands

``` shell
➜  ~ ✗ tfversion
tfversion v0.1.5 - Switcher Terraform

Usage:
  tfversion [command]

Available Commands:
  help        Help about any command
  install     Install new versions or switch
  list        List of terraform versions
  remove      Remove local version of Terraform
  version     Version installed of switcher Terraform

Flags:
  -h, --help   help for tfversion

Use "tfversion [command] --help" for more information about a command.
```

## Docker

### Install

Pull image `docker pull perriea/tfversion:latest`.   
Execute command in the terminal : `docker run -it perriea/tfversion`.   

## Dependancies

- [spf13/cobra](https://github.com/spf13/cobra),
- [google/go-github](https://github.com/google/go-github)

## Used by

- [perriea/tfwrapper](https://github.com/perriea/tfwrapper)

## License

The MIT License (MIT)   
Copyright (c) 2017-2018 Aurelien PERRIER
