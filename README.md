# tfversion

`tfversion` is a command created to switch between different versions of [terraform](https://www.terraform.io).   
Functional started from version 0.7.

## Build Project

- [Install Golang](https://golang.org/doc/install) (add var ENV),
- [Install Govendor](https://github.com/kardianos/govendor),
- Build `go build .`,
- Add in your `.bashrc` (Linux) or `.bash_profile` : `export PATH=$PATH:~/terraform/bin`

## Usage

### Install version

**Example :**   
`tfversion --version 0.9.4`  

**Result :**
``` shell
aurelien@localhost:~ tfversion --version 0.9.4                  
Attempting to download version: 0.9.4
Start download ...
Unzip file ...
Install the binary file ...
Installed 0.9.4, Thanks ! â™¥ 
```

### List versions

**Example :**   
`tfversion --list`  

**Result :**
``` shell
tfversion --list         
Versions availables of terraform (tfversion support <= 0.7) :
[0.9.4 0.9.3 0.9.2 0.9.1 0.9.0 0.9.0-beta2 0.9.0-beta1 0.8.8 0.8.7 0.8.6 0.8.5 0.8.4 0.8.3 0.8.2 0.8.1 0.8.0 0.8.0-rc3 0.8.0-rc2 0.8.0-rc1 0.8.0-beta2 0.8.0-beta1 0.7.13 0.7.12 0.7.11 0.7.10 0.7.9 0.7.8 0.7.7 0.7.6 0.7.5 0.7.4 0.7.3 0.7.2 0.7.1 0.7.0 0.7.0-rc4 0.7.0-rc3 0.7.0-rc2 0.7.0-rc1 0.6.16 0.6.15 0.6.14 0.6.13 0.6.12 0.6.11 0.6.10 0.6.9 0.6.8 0.6.7 0.6.6 0.6.5 0.6.4 0.6.3 0.6.2 0.6.1 0.6.0 0.5.3 0.5.1 0.5.0 0.4.2 0.4.1 0.4.0 0.3.7 0.3.6 0.3.5 0.3.1 0.3.0 0.2.2 0.2.1 0.2.0 0.1.1 0.1.0]
```

## Roadmap
  
- Switch between version without internet (if the zip archive is allready in folder tmp),
- List the local versions of terraform,
- Command cleanup tmp folder,
- Install script.

And other things ...

## Dependancies

- [kardianos/govendor](https://github.com/kardianos/govendor),
- [fatih/color](https://github.com/fatih/color).

## License

The MIT License (MIT)

Copyright (c) 2017 Aurelien Perrier

