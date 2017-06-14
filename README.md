# tfversion

`tfversion` is a command created to switch between different versions of [terraform](https://www.terraform.io).   
Functional for all versions.

## Build Project

- [Install Golang](https://golang.org/doc/install) (add var ENV),
- Build with commands `go build` or `go get github.com/perriea/tfversion`,
- Add in your `.bashrc` (Linux) or `.bash_profile` : `export PATH=$PATH:~/terraform/bin`

## Commands

``` shell
➜  ~ ✗ tfversion
tfversion v0.1.3

Usage:

  tfversion <command> [option]

Options:

  -h, --help      display help information
  -v, --version   show version and check update

Commands:

  install     install new versions or switch.
  uninstall   uninstall local version of Terraform
  list        list online or offline version of terraform
  test        test provider cloud (AWS, GCP)
```

## Docker

### Require

- Docker,
- AWS access keys (`~/.aws`),
- GCP access keys (`~/.gcloud` or an other path) and SDK,
- SSH folder (`~/.ssh`).

### Install

Pull image `docker pull perriea/tfversion`.   
Execute command in the terminal : `docker run -it -v ~/.aws:/root/.aws -v ~/.ssh:/root/.ssh perriea/tfversion sh`.   

## Dependancies

- [kardianos/govendor](https://github.com/kardianos/govendor),
- [mkideal/cli](https://github.com/mkideal/cli)
- [google/go-github](https://github.com/google/go-github),
- [aws/aws-sdk-go](https://github.com/aws/aws-sdk-go) (modules: Session, EC2, AWSErr).
- [GoogleCloudPlatform/google-cloud-go](https://github.com/GoogleCloudPlatform/google-cloud-go)

## License

The MIT License (MIT)   
Copyright (c) 2017 Aurelien PERRIER