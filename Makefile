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
	go get -u github.com/kardianos/govendor

docker:
	@docker build . -t perriea/tfversion


##
## Development
##

test:
	@GOGC=off go test $$(go list ./... | grep -v '/vendor/')

dist-test:
	@$(MAKE) test


##
## Building
##
dist: clean
	$(MAKE) build

build:
	@mkdir -p ./bin
	GOGC=off go build -i -o ./bin/tfversion ./

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