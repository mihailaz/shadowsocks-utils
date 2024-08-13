BUILD := $(shell git rev-parse --short HEAD)
DATE = $(shell date)
CURRENT_DIR := $(shell pwd)
ARTIFACTS_DIR=$(CURRENT_DIR)/artifacts

all: clean build

build: build-ssinfo

build-ssinfo:
	mkdir -p $(ARTIFACTS_DIR)
	go build -C cmd/ssinfo -o $(ARTIFACTS_DIR)/ssinfo -ldflags "-X 'main.build=$(BUILD)' -X 'main.date=$(DATE)'"

clean:
	rm -rf $(ARTIFACTS_DIR)

.PHONY: all build build-ssinfo clean