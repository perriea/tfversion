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

all: build

docker:
	@GOOS=linux $(GOBUILD) -i -o ./tfversion ./
	$(DOCKERBUILD) . -t $(CONTNAME)

test:
	$(GOTEST) $$(go list ./... | grep -v '/vendor/')

build:
	$(GOBUILD) -i -o ./tfversion ./

vendor-list:
	@govendor list

vendor-update:
	@govendor update +vendor