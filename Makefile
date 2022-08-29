ARBOR_VERSION?=0.0.0-rc0
export ARBOR_VERSION := $(ARBOR_VERSION)

export PATH := $(HOME)/go/bin/:$(PWD)/utils:$(PATH)
.PHONY: utils

all: build publish

build: test generate
	go build -o arbor ./cmd/arbor/main.go

test: 
	go test -v ./...

test-lite:
	go test ./...

utils: 
	go build -o ./utils/gen_visitors ./utils/

generate: utils
	go generate ./...

ldflags: utils
	./utils/gen_visitors -ldFlags

debug: test generate ldflags
	go build -tags debug -ldflags '$(shell ARBOR_VERSION=$(ARBOR_VERSION) ./utils/gen_visitors -ldFlags)' -o arbor ./cmd/arbor/main.go


publish: test generate ldflags
	go build -ldflags '$(shell ARBOR_VERSION=$(ARBOR_VERSION) ./utils/gen_visitors -ldFlags)' -o arbor ./cmd/arbor/main.go

delve: debug
	dlv exec ./arbor build docs/example/example.ab

skip_test:
	go build -o arbor ./cmd/arbor/main.go
