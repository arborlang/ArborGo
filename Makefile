ARBOR_VERSION?=0.0.0-rc0
all: toolchain
	go build -o arbor ./cmd/arbor/main.go

toolchain: build

build:
	go build -o plugins/build -buildmode=plugin ./Commands/build/

publish: toolchain
	go build -ldflags "-X main.Version=$(ARBOR_VERSION)" -o arbor ./cmd/arbor/main.go