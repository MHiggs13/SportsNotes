# Go params
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test

BINARY_NAME=ballNotes
BUILD_DIR=build/


all: build
build:
	$(GOBUILD) -o $(BUILD_DIR)$(BINARY_NAME) -v src/*

clean:
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)


