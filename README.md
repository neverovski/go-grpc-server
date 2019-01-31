# Quick start

1. Install Go https://golang.org/doc/install
2. Install GoLand https://www.jetbrains.com/go/nextversion/
3. Install Docker https://docs.docker.com/install/
4. Istall protoc https://github.com/google/protobuf/releases
    - `go get -u github.com/golang/protobuf/{proto,protoc-gen-go}`
    - `go get -u google.golang.org/grpc`
    - `go get -u golang.org/x/net/context`
5. Prepare FS before `git clone`
   - `mkdir -p go-grpc-server/src`
   - `export GOPATH=$PWD/go-grpc-server`
   - `cd go-grpc-server/src`
   - `git clone git@github.com:neverovski/go-grpc-server.git`
6. Generate protoc for Go
    - `protoc --go_out=plugins=grpc:. *.proto`
        - --go_out - parameter set to the directory you want to output the Go code to.
        - plugins=grpc - specifies the list of sub-plugins to load. The only plugin in this repo is grpc.

# Items to do

 - [ ] Create protobuf 
 - [ ] Graphql integration
 - [ ] Docker integration
 - [ ] Postgres schema and test data scripts prepared
 - [ ] Go debugger setuped for goland
 - [ ] Update graphql according to postgre schema 
 - [ ] Do actuall query to postgre https://github.com/go-pg/pg
 - [ ] Add CRUD mutations to main entities https://graphql.org/learn/queries/
 - [ ] Implement oauth login/registration https://github.com/golang/oauth2
 - [ ] Allow oauth for our service https://github.com/go-oauth2/oauth2
 - [ ] graphql+security=friendship https://www.howtographql.com/advanced/4-security/ , https://dev-blog.apollodata.com/auth-in-graphql-part-2-c6441bcc4302
