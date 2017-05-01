# tfversion

`tfversion` is a command created to switch between different versions of [terraform](https://www.terraform.io).   
Functional started from version 0.7.

## Build Project

- [Install Golang](https://golang.org/doc/install) (add var ENV),
- [Install Govendor](https://github.com/kardianos/govendor),
- Build `go build .`,
- Add in your `.bashrc` (Linux) or `.bash_profile` : `export PATH=$PATH:~/terraform/bin`

## Docker

### Require

- Docker,
- AWS credidencial (`~/.aws`),
- SSH folder (`~/.ssh`).

### Install

Pull image `docker pull perriea/tfversion`.   
Execute command in the terminal : `docker run -it -v ~/.aws:/root/.aws -v ~/.ssh:/root/.ssh perriea/tfversion sh`.


## Roadmap
  
- List the local versions of terraform and **show curent version**,
- Install script.

And other things ...

## Dependancies

- [kardianos/govendor](https://github.com/kardianos/govendor),
- [fatih/color](https://github.com/fatih/color).

## License

The MIT License (MIT)

Copyright (c) 2017 Aurelien Perrier

