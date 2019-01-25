ARBOR_VERSION?=0.0.0-rc0
all: toolchain
	go build -o arbor ./cmd/arbor/main.go

toolchain: build run

build:
	go build -o plugins/build -buildmode=plugin ./Commands/build/

run:
	go build -o plugins/run -buildmode=plugin ./Commands/run/

publish: toolchain
	go build -ldflags "-X main.Version=$(ARBOR_VERSION)" -o arbor ./cmd/arbor/main.go

test_file: all
	./arbor build -wast -o test.wast test.ab
	
wast_test:
	wat2wasm -o test.wasm test.wast
	./arbor run --wasm --entrypoint main test.wasm

test_run: test_file wast_test