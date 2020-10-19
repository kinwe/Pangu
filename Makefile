.PHONY: build clean tool lint help

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_PATH=./build
BINARY_NAME=main
BINARY_UNIX=$(BINARY_PATH)/$(BINARY_NAME)_linux

all: run
build:
	$(GOBUILD) -o $(BINARY_UNIX)
test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run:
	$(GOBUILD) -o $(BINARY_UNIX) -v
	./$(BINARY_UNIX)
deps:
	$(GOGET) github.com/markbates/goth
	$(GOGET) github.com/markbates/pop

help:
    @echo "make: run all"
    @echo "make build: run build"
    @echo "make test: run test"
    @echo "make clean: remove object files and cached files"