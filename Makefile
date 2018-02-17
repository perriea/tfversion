########################################
#          TFVERSION MAKEFILE          #
#        Author: Aurelien PERRIER      #
########################################

GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test

DOCKERCMD=docker
DOCKERBUILD=$(DOCKERCMD) build
CONTNAME=perriea/tfversion:latest

.PHONY: help test build docker vendor-list vendor-update

all:
	@echo "*******************************"
	@echo "**   tfversion build tools   **"
	@echo "*******************************"
	@echo "make <cmd>"
	@echo ""
	@echo "commands:"
	@echo "  test           - run go tests"
	@echo "  build          - build binaries into bin/ directory"
	@echo ""
	@echo "  docker         - Launch container with binary"
	@echo "  vendor-list    - List dependencies"
	@echo "  vendor-update  - Upgrade dependencies"

##
## Docker
##

docker:
	@GOOS=linux $(GOBUILD) -i -o ./tfversion ./
	$(DOCKERBUILD) . -t $(CONTNAME)


##
## Development
##

test:
	$(GOTEST) $$(go list ./... | grep -v '/vendor/')


##
## Building
##

build:
	$(GOBUILD) -i -o ./tfversion ./


##
## Dependency mgmt
##
vendor-list:
	@govendor list

vendor-update:
	@govendor update +vendor