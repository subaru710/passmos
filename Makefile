PACKAGES=$(shell go list ./... | grep -v '/vendor/')

all: get_tools get_vendor_deps build test

get_tools:
	go get github.com/golang/dep/cmd/dep

build:
	go build -o bin/passmos-cli cmd/passmos-cli/main.go && go build -o bin/passmosd cmd/passmosd/main.go

get_vendor_deps:
	@rm -rf vendor/
	@dep ensure

test:
	@go test -v $(PACKAGES)

benchmark:
	@go test -bench=. $(PACKAGES)

.PHONY: all build test benchmark
