FROM golang:1.15.1 as builder
# Install xz
RUN apt-get update && apt-get install -y xz-utils \
    && rm -rf /var/lib/apt/lists/*
# Install UPX
ADD https://github.com/upx/upx/releases/download/v3.94/upx-3.94-amd64_linux.tar.xz /usr/local
RUN xz -d -c /usr/local/upx-3.94-amd64_linux.tar.xz | \
    tar -xOf - upx-3.94-amd64_linux/upx > /bin/upx && \
    chmod a+x /bin/upx

# Install dep
RUN go get github.com/golang/dep/cmd/dep


# Set working directory

# Install packages
WORKDIR /Go/src/protocol-go-svcs
# Add contents
COPY . .
RUN dep ensure --vendor-only


# Build the source lorem-consul.d
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o proto_server server/server.go
# Strip and compress the binary
RUN strip --strip-unneeded proto_server
RUN upx proto_server


# Use a minimal alpine image
FROM alpine:3.7
# Add ca-certificates in case you need them
RUN apk add --no-cache ca-certificates
# Set working directory
# Copy the binary from builder
COPY --from=builder /Go/src/protocol-go-svcs /usr/local/bin

#Or docker run --expose=8080-8080 or docker run -p 8080-8080
EXPOSE  8080

# Run the binary
CMD ["-help"] 
ENTRYPOINT ["proto_server"]



