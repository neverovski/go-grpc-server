TARGET=go-grpc-server

ifndef $(GOPATH)
    GOPATH=$(shell go env GOPATH)
    export GOPATH
endif

get:
	echo $(GOPATH)

protos_news:
	echo "Start build proto news app"
	protoc -I/usr/local/include -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:. \
		api/news/v1/pb/news.proto
	protoc -I/usr/local/include -I. \
    	-I${GOPATH}/src \
    	-I${GOPATH}/src/github.com/googleapis/googleapis/ \
    	-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    	--plugin=protoc-gen-grpc-gateway=${GOPATH}/bin/protoc-gen-grpc-gateway \
    	--grpc-gateway_out=logtostderr=true:. \
    	api/news/v1/pb/news.proto
	protoc -I/usr/local/include -I. \
    	-I${GOPATH}/src \
    	-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    	--swagger_out=logtostderr=true:. \
    	api/news/v1/pb/news.proto
	echo "End build proto news app"

protos_user:
	echo "Start build proto user app"
	protoc -I/usr/local/include -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:. \
		api/user/v1/pb/user.proto
	protoc -I/usr/local/include -I. \
    	-I${GOPATH}/src \
    	-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    	--grpc-gateway_out=logtostderr=true:. \
    	api/user/v1/pb/user.proto
	protoc -I/usr/local/include -I. \
    	-I${GOPATH}/src \
    	-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    	--swagger_out=logtostderr=true:. \
    	api/user/v1/pb/user.proto
	echo "End build proto user app"