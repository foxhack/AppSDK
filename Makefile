#
# Copyright (c) 2018 Cavium
#
# SPDX-License-Identifier: Apache-2.0
#


.PHONY: build clean test docker run


GO=CGO_ENABLED=0 GO111MODULE=on go
GOCGO=CGO_ENABLED=1 GO111MODULE=on go


MICROSERVICES=apps/nsdemo/cmd/nsdemo

.PHONY: $(MICROSERVICES)

VERSION=$(shell cat ./VERSION)

GOFLAGS=-ldflags "-X github.com/edgexfoundry/nsplussdk.Version=$(VERSION)"

build: $(MICROSERVICES)

apps/nsdemo/cmd/nsdemo:
	$(GOCGO) build $(GOFLAGS) -o $@ ./apps/nsdemo/cmd/


clean:
	rm -f $(MICROSERVICES)

test:
	GO111MODULE=on go test -cover ./...
	GO111MODULE=on go vet ./...
	gofmt -l .
	[ "`gofmt -l .`" = "" ]


