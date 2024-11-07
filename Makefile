GO=go

version=$(shell git rev-parse --short HEAD)
build_date=$(shell date -u "+%Y-%m-%dT%H:%M:%SZ")
goVersion=$(shell $(GO) version | cut -d' ' -f3)
GO_BUILD=$(GO) build -ldflags "-s -w -X github.com/luzen23141/mouse/pkg.Version=$(version) -X github.com/luzen23141/mouse/pkg.BuildDate=$(build_date) -X github.com/luzen23141/mouse/pkg.GoVersion=$(goVersion)" -tags=sonic

SERVER_PATH=./main.go
BIN_PATH=./bin

build:
	$(GO_BUILD) -o $(BIN_PATH)/main $(SERVER_PATH)

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO_BUILD) -o $(BIN_PATH)/linux-main $(SERVER_PATH)
	upx -9 $(BIN_PATH)/linux-main

# 設置 ".PHONY" 目標，以確保這些目標不是實際的檔案名稱
.PHONY: build build-linux
