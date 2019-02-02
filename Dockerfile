# Accept the Go version for the image to be set as a build argument.
# Default to Go 1.11
ARG GO_VERSION=1.11

# First stage: build the executable.
FROM golang:${GO_VERSION}-alpine AS builder

# Set the environment variables for the go command:
# * CGO_ENABLED=0 to build a statically-linked executable
# * GOFLAGS=-mod=vendor to force `go build` to look into the `/vendor` forlder.
ENV CGO_ENABLED=0

ADD . /go/src/go-grpc-server

RUN echo GOPATH=$GOPATH
WORKDIR /go/src/go-grpc-server

# Import the code from the context.
COPY ./ ./

# The -gcflags "all=-N -l" flag helps us get a better debug experience
RUN go build -gcflags "all=-N -l" -o /go-grpc-server ./cmd/api-server

# Add git
RUN \
    apk update && \
    apk add --no-cache git

# Compile protobuf, grpc and context
RUN \
    go get github.com/segmentio/ksuid && \
    go get github.com/golang/protobuf/proto && \
    go get github.com/golang/protobuf/protoc-gen-go && \
    go get google.golang.org/grpc && \
    go get golang.org/x/net/context

# Final stage: the running container.
FROM scratch AS final

# Port 8080 belongs to our application
EXPOSE 8080

WORKDIR /
COPY --from=builder /go-grpc-server .

# Run
CMD ["/go-grpc-server"]