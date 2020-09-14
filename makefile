# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
SERVER_BINARY=users
CLIENT_BINARY=users-cli

all: build test
build: 
	$(GOBUILD) ./cmd/$(SERVER_BINARY) && $(GOBUILD) ./cmd/$(CLIENT_BINARY)
test: 
	$(GOTEST) -v ./...
clean: 
	$(GOCLEAN)
	rm -f $(SERVER_BINARY)
	rm -f $(CLIENT_BINARY)
run:
	$(GORUN) ./cmd/$(SERVER_BINARY)

