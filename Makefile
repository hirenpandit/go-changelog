# Makefile for changelog project

# Variables
BINARY_NAME=changelog
BINARY_DIR=bin
CMD_DIR=cmd

.PHONY: all build clean

all: build

build:
	@echo "Building $(BINARY_NAME)..."
	go build -o $(BINARY_DIR)/$(BINARY_NAME) ./$(CMD_DIR)

clean:
	@echo "Cleaning..."
	@rm -rf $(BINARY_DIR)
