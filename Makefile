########################################
#          TFVERSION MAKEFILE          #
#        Author: Aurelien PERRIER      #
########################################

GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test

DOCKERCMD=docker
DOCKERBUILD=$(DOCKERCMD) build
CONTNAME=perriea/tfversion

.PHONY: help test retest coverage build clean tools dist_tools deps update_deps dist docker

all:
	@echo "*******************************"
	@echo "**   tfversion build tools   **"
	@echo "*******************************"
	@echo "make <cmd>"
	@echo ""
	@echo "commands:"
	@echo "  test        - run go tests"
	@echo "  build       - build binaries into bin/ directory"
	@echo "  clean       - clean up bin/ directory"
	@echo ""
	@echo "  dist        - clean build with deps and tools"
	@echo "  tools       - go get's a bunch of tools for dev"
	@echo "  docker      - Launch container with binary"

##
## Tools
##
tools:
	$(GOCMD) get -u github.com/kardianos/govendor

docker:
	$(DOCKERBUILD) . -t $(CONTNAME)


##
## Development
##

test:
	$(GOTEST) $$(go list ./... | grep -v '/vendor/')


##
## Building
##
dist: clean
	$(MAKE) build

build:
	@mkdir -p ./bin
	$(GOBUILD) -i -o ./bin/tfversion ./

clean:
	@rm -rf $$GOPATH/pkg/*/github.com/perriea/tfversion{,.*}
	@rm -rf ./bin


##
## Dependency mgmt
##
vendor-list:
	@govendor list

vendor-update:
	@govendor update +vendor