.PHONY: build test clean prepare update docker

GO = CGO_ENABLED=0 GO111MODULE=on go

MICROSERVICES=cmd/helloworld-device

.PHONY: $(MICROSERVICES)

VERSION=$(shell cat ./VERSION)
GIT_SHA=$(shell git rev-parse HEAD)
GOFLAGS=-ldflags "-X github.com/edgexfoundry/helloworld-device.Version=$(VERSION)"

build: $(MICROSERVICES)
	$(GO) build ./...

cmd/helloworld-device:
	$(GO) build $(GOFLAGS) -o $@ ./cmd

test:
	$(GO) test ./... -cover

clean:
	rm -f $(MICROSERVICES)

run:
	./cmd/helloworld-device