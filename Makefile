CC=go
PROJECT_PATH=$(shell pwd)
PROJECT_NAME=golang-websocket-client
MODULE_NAME=chat-client
TARGET_DIR=bin
VERSION_NUM=$$(cat deploy/version_num.txt)
BUILD_NUM=$$(cat deploy/build_num.txt)
OUTPUT=$(PROJECT_PATH)/$(TARGET_DIR)/$(MODULE_NAME)_$(VERSION_NUM).$(BUILD_NUM)
MAIN_DIR=main
LDFLAGS=-X main.BUILD_TIME=`date -u '+%Y-%m-%d_%H:%M:%S'`
LDFLAGS+=-X main.GIT_HASH=`git rev-parse HEAD`
LDFLAGS+=-X main.VERSION_NUMBER=$(VERSION_NUM)
LDFLAGS+=-X main.BUILD_NUMBER=BUILD_NUM=$(BUILD_NUM)
LDFLAGS+=-s -w

all: build

build_num :
	@echo $$(($$(cat $(BUILD_NUM_FILE)) + 1 )) > $(BUILD_NUM_FILE)

build:
	CGO_ENABLED=0 GOOS=linux go build -ldflags "$(LDFLAGS)" -o $(OUTPUT) $(PROJECT_PATH)/$(MAIN_DIR)
	cp $(OUTPUT) ./ex_$(MODULE_NAME)

clean:
	rm -f $(PROJECT_PATH)/$(TARGET_DIR)/$(MODULE_NAME)_$(VERSION)*
	rm -f $(PROJECT_PATH)/ex_*

