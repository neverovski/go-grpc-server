TARGET=go-grpc-server
GOCMD := go

ifndef $(GOPATH)
    GOPATH=$(shell go env GOPATH)
    export GOPATH
endif

get: \
	protos \
	gofmt

protos:
	for file in $$(git ls-files '*.proto'); do \
		protoc -I/usr/local/include -I. \
        	-I${GOPATH}/src \
        	-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
        	--go_out=plugins=grpc:. $$file; \
        protoc -I/usr/local/include -I. \
            -I${GOPATH}/src \
            -I${GOPATH}/src/github.com/googleapis/googleapis/ \
            -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
            --plugin=protoc-gen-grpc-gateway=${GOPATH}/bin/protoc-gen-grpc-gateway \
            --grpc-gateway_out=logtostderr=true:. $$file; \
        protoc -I/usr/local/include -I. \
            -I${GOPATH}/src \
            -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
            --plugin=protoc-gen-swagger=${GOPATH}/bin/protoc-gen-swagger \
            --swagger_out=logtostderr=true:. $$file; \
	done

gofmt:
	$(GOCMD) fmt ./...