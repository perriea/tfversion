# tfversion

`tfversion` is a command created to switch between different versions of [terraform](https://www.terraform.io).   
Functional for all versions.

## Build Project

- [Install Golang](https://golang.org/doc/install) (add var ENV),
- Build `go build` or `go get github.com/perriea/tfversion`,
- Add in your `.bashrc` (Linux) or `.bash_profile` : `export PATH=$PATH:~/terraform/bin`

## Commands

``` shell
➜  ~ tfversion help
tfversion v0.1.2

Usage:

  tfversion <command> [option]

Options:

  -v, --version   Show version and check update

Commands:

  help        Display help informations
  install     Install new versions or switch.
  uninstall   Uninstall local version of Terraform
  list        List online or offline version of terraform
  test        Test provider cloud (AWS)
```

## Docker

### Require

- Docker,
- AWS credidencial (`~/.aws`),
- SSH folder (`~/.ssh`).

### Install

Pull image `docker pull perriea/tfversion`.   
Execute command in the terminal : `docker run -it -v ~/.aws:/root/.aws -v ~/.ssh:/root/.ssh perriea/tfversion sh`.   

## Dependancies

- [kardianos/govendor](https://github.com/kardianos/govendor),
- [google/go-github](https://github.com/google/go-github),
- [aws/aws-sdk-go](https://github.com/aws/aws-sdk-go) (modules: Session, EC2, AWSErr).

## License

The MIT License (MIT)   
Copyright (c) 2017 Aurelien PERRIER