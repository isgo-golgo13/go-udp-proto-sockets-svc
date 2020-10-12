 # Go parameters
    GOCMD=go
    GOBUILD=go build
    GOCLEAN=go clean
    GOTEST=go test
    SERVER=server
  
    all: build
    build: 
		$(GOBUILD) -o $(SERVER) -v
    test: 
		$(GOTEST) -v ./...
    clean: 
		$(GOCLEAN)
			rm -f $(SERVER)
    run:
		$(GOBUILD) -o $(SERVER) -v ./...
		.$(SERVER)
    deps:
		dep ensure -v --vendor-only
