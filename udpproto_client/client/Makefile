 # Go parameters
    GOCMD=go
    GOBUILD=go build
    GOCLEAN=go clean
    GOTEST=go test
    CLIENT=client
 
    all: build
    build: 
		$(GOBUILD) -o $(CLIENT) -v
    test: 
		$(GOTEST) -v ./...
    clean: 
		$(GOCLEAN)
			rm -f $(CLIENT)
    run:
		$(GOBUILD) -o $(CLIENT) -v
		./$(CLIENT)

    deps:
		dep ensure -v --vendor-only
