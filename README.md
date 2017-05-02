# protobuf-intro

## A light introduction to Protocol Buffers

First make sure that protoc is intalled. On a Mac, this can be done with `Homebrew`:

`brew install protobuf`

1. Generate Python code

`protoc -I=protobuf --python_out=python protobuf/addressbook.proto`

See it in action by running `python python/app.py`

2. Generat Go code

First install the `Go` protobuf compiler (assumes `Go` development environment, i.e. $GOPATH, etc. is already setup):

`go get -u github.com/golang/protobuf/{proto,protoc-gen-go}`

And generate classes from `proto` files:

`go get github.com/jmarin/protobuf-intro` (ignore warning)

`protoc -I=protobuf --go_out=$GOPATH/src/github.com/jmarin/protobuf-intro protobuf/addressbook.proto`

Compile project with `go build go/app.go` and run it with `go run go/app.go`


