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
 - [x] Docker integration
 - [ ] Postgres schema and test data scripts prepared
 - [ ] Go debugger setuped for goland
 - [ ] Update graphql according to postgre schema 
 - [ ] Do actuall query to postgre https://github.com/go-pg/pg
 - [ ] Add CRUD mutations to main entities https://graphql.org/learn/queries/
 - [ ] Implement oauth login/registration https://github.com/golang/oauth2
 - [ ] Allow oauth for our service https://github.com/go-oauth2/oauth2
 - [ ] graphql+security=friendship https://www.howtographql.com/advanced/4-security/ , https://dev-blog.apollodata.com/auth-in-graphql-part-2-c6441bcc4302
 
 # Standard Go Project Layout
 
 This is a basic layout for Go application projects. It represents the most common directory structure with a number of small enhancements along with several supporting directories common to any real world application. 
 
 This project layout is intentionally generic and it doesn't try to impose a specific Go package structure.
 
 Clone the repository, keep what you need and delete everything else!
 
 [Go Project Layout](https://medium.com/golang-learn/go-project-layout-e5213cdcfaa2) - additional background information.
 
 ## Go Directories
 
 ### `/cmd`
 
 Main applications for this project.
 
 The directory name for each application should match the name of the executable you want to have (e.g., `/cmd/myapp`).
 
 Don't put a lot of code in the application directory. If you think the code can be imported and used in other projects, then it should live in the `/pkg` directory. If the code is not reusable or if you don't want others to reuse it, put that code in the `/internal` directory. You'll be surprised what others will do, so be explicit about your intentions!
 
 It's common to have a small `main` function that imports and invokes the code from the `/internal` and `/pkg` directories and nothing else.
 
 See the [`/cmd`](cmd/README.md) directory for examples.
 
 ## Service Application Directories
 
### `/api`

OpenAPI/Swagger specs, JSON schema files, protocol definition files.

See the [`/api`](api/README.md) directory for examples.

### `/deployments`

IaaS, PaaS, system and container orchestration deployment configurations and templates (docker-compose and etc).
See the [`/deployments`](deployments/README.md) directory for examples.