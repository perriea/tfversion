.PHONY: tfversion release clean test race graph

GOLANG_ENV?=CGO_ENABLED=0
IMAGE_TAG:=$(shell ./hack/image-tag.sh)
LDFLAGS?=-v -ldflags="-w -s -X main.version=$(IMAGE_TAG)"
TEST_FLAGS?=
PKG_LIST := $(shell go list ./... | sort -u)

# NB default target architecture is amd64. If you would like to try the
# other one -- pass an ARCH variable, e.g.,
#  `make ARCH=arm64`
ifeq ($(ARCH),)
	ARCH=amd64
endif

all: tfversion

tfversion:
	$(GOLANG_ENV) go build $(LDFLAGS) -o build/tfversion $(PWD)/cmd

release:
	for arch in amd64; do \
		for os in linux darwin; do \
			$(GOLANG_ENV) GOOS=$$os GOARCH=$$arch go build -o "build/tfversion_"$$os"_$$arch" $(LDFLAGS) -ldflags "-X main.version=$(IMAGE_TAG)" $(PWD)/cmd; \
		done; \
	done;

clean:
	go clean
	rm -rf ./build

race:
	go test -race -coverprofile=coverage.txt -covermode=atomic ${PKG_LIST}

graph:
	go-callvis -file=data -format=png -group pkg -focus="" -limit github.com/perriea/tfversion $(PWD)/cmd

test:
	go test ${TEST_FLAGS} ${PKG_LIST}
