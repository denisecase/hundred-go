# You must have GNU Make installed. 
# To run a target, run make <targetname>, e.g., make echo
# IMPORTANT: Use tabs not spaces to indent.
# @ means "don't print the command - just the results"

# Some Boilerplate for Makefile to work in Windows or Bash

ifeq ($(OS),Windows_NT) 
RM = del /Q /F
CP = copy /Y
ifdef ComSpec
SHELL := $(ComSpec)
endif
ifdef COMSPEC
SHELL := $(COMSPEC)
endif
else
RM = rm -rf
CP = cp -f
endif

BINARY_NAME=bin/hundred-go

help:
	@echo "======================================================"
	@echo "  A Makefile has targets (short aliases for commands)."
	@echo "  Run `make help` to run this target.                 "
	@echo "  See Makefile to find more targets.                  "
	@echo "     Run Go commands with:                          "
	@echo "        make download-deps                                    "
	@echo "        make download-deps-vendor                                    "
	@echo "        make vet                                    "
	@echo "        make build                                   "
	@echo "        make test                                   "
	@echo "        make run                                   "
	@echo "        make clean                                   "
	@echo "        make all                                   "
	@echo "        make update-deps                                    "
	@echo "  It may include commands to Dockerize the app.     "
	@echo "===================================================="

download-deps:
	go mod download

download-deps-vendor:
	go mod vendor

vet:
	go vet

lint:
	golangci-lint run --enable-all

build:
	go build -o ${BINARY_NAME}

test:
	go test -v main.go

run: build
	./${BINARY_NAME}

clean:
	go clean
	-$(RM) bin

all: build test

update-deps:
	go get -u
	go mod tidy
