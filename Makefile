GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

.DEFAULT_GOAL := all

prepare:
	mkdir -p build

backend-debug:
	cd backend && go build -o ../build/narravo-backend-debug.exe \
		-ldflags="-s -w" \
		-v \
		.

backend-release:
	cd backend && go build -o ../build/narravo-backend-release.exe \
		-ldflags="-s -w" \
		-tags=release,verify \
		-v \
		.

clean:
	rm -rf build

release: prepare backend-release

debug: prepare backend-debug

all: clean release

.PHONY: prepare clean
