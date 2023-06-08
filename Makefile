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

backend-test:
	@echo > Start backend testing
	@cd backend && \
	go test ./internal/utils/test -tags release && \
	go test ./dao/test -tags release && \
	go test ./server/routers/test -tags release && \
	go test ./server/middlewares/test -tags verify

clean:
	rm -rf build

test: backend-test

release: prepare backend-release

debug: prepare backend-debug

all: clean release

.PHONY: prepare clean
