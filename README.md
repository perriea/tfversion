# tfversion

`tfversion` is a command created to switch between different versions of [Terraform](https://www.terraform.io).

## Build Project

- [Install Golang](https://golang.org/doc/install) (add var ENV),
- Build with commands `go build`, `make build` or `go get github.com/perriea/tfversion`,
- Add in your `.bashrc` (Linux), `.bash_profile` (Darwin, MacOS) or `.zshrc` : `export PATH=$PATH:~/.tfversion/bin` and `GOOGLE_APPLICATION_CREDENTIALS` (contain path of Google JSON credential, **optional**).

## Commands

``` shell
➜  ~ ✗ tfversion
tfversion v0.1.4-dev

Usage:

  tfversion <command> [option]

Options:

  -h, --help      display help information
  -v, --version   show version and check update

Commands:

  install     install new versions or switch
  uninstall   uninstall local version of Terraform
  list        list of terraform versions
```

## Docker

### Install

Pull image `docker pull perriea/tfversion:latest`.   
Execute command in the terminal : `docker run -it perriea/tfversion`.   

## Dependancies

- [spf13/cobra](https://github.com/spf13/cobra),
- [google/go-github](https://github.com/google/go-github)

## License

The MIT License (MIT)   
Copyright (c) 2017-2018 Aurelien PERRIER