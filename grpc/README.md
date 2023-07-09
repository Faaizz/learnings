# gRPC: An Example

A gRPC example with a Go server and Python client that exhibits the 3 (of the 4) kinds of service methods.

## Prerequisites
### Server
Install the protocol compiler plugins for Go:
```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```
### Client
Install required python libraries:
```shell
pip install -r client/requirements.txt
```

## Generate files
Generate Go and Python files from [.proto](model/todo.proto) service definition.
```shell
# Server (Go) files
protoc --go_out=server --go_opt=paths=source_relative --go-grpc_out=server --go-grpc_opt=paths=source_relative model/todo.proto
# Client (Python) files
python -m grpc_tools.protoc -I model --python_out=client --pyi_out=client --grpc_python_out=client model/todo.proto
```

## Usage
### Server
Start server with:
```shell
cd server
go mod tidy
go run . -port=8080
```
### Client
In a different terminal window, run client with:
```shell
cd client
python client.py
```

## References
- [https://grpc.io/docs/languages/go/basics/](https://grpc.io/docs/languages/go/basics/)
- [https://grpc.io/docs/languages/python/basics/](https://grpc.io/docs/languages/python/basics/)
- [https://www.udemy.com/course/fundamentals-of-backend-communications-and-protocols/learn/lecture/34630280](https://www.udemy.com/course/fundamentals-of-backend-communications-and-protocols/learn/lecture/34630280)
