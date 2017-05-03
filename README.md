# protobuf-intro

## A light introduction to Protocol Buffers

First make sure that `protoc` is intalled. On a Mac, this can be done with `Homebrew`:

`brew install protobuf`

1. Generate Python code

`protoc -I=protobuf --python_out=python protobuf/addressbook.proto`

See it in action by running `python python/app.py`

2. Generate Go code

First install the `Go` protobuf compiler (assumes `Go` development environment, i.e. $GOPATH, etc. is already setup):

`go get -u github.com/golang/protobuf/{proto,protoc-gen-go}`

And generate classes from `proto` files:

`go get github.com/jmarin/protobuf-intro` (ignore warning)

`protoc -I=protobuf --go_out=$GOPATH/src/github.com/jmarin/protobuf-intro protobuf/addressbook.proto`

Compile project with `go build go/app.go` and run it with `go run go/app.go`

3. Generate Java code

`protoc -I=protobuf --java_out=java/protobuf/src/main/java protobuf/addressbook.proto`

From the `java/protobuf` directory run:

`mvn clean install`

`mvn exec:java`

4. Generate JavaScript code

From the javascript subdirectory run:

`npm install`

`node index.js`

5. Generate C++ code

`protoc -I=protobuf --cpp_out=cpp protobuf/addressbook.proto`
