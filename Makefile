# Makefile用于交叉编译 在确定要发布时再编译，避免远程仓库体积过大
SHELL := /bin/bash
BASEDIR = $(shell pwd)

# Go parameters
GOCMD=go
GOENV=GOOS=linux GOARCH=amd64
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME_SIGNALER=gosignaler

main:
	@echo "------build signaler--------"
	$(GOENV) $(GOBUILD) -o $(BINARY_NAME_SIGNALER) -v main.go hub.go handler.go client.go
test:
	$(GOTEST) -v ./...
clean:
	@echo "------clean--------"
	rm -f $(BINARY_NAME_SIGNALER)
help:
	@echo "make - compile the source code"
	@echo "make clean - remove binary file and vim swp files"
.PHONY: clean help



