# grpc-ms-example
This is a small GRPC microservice example. The idea is to show how gRPC can be used to serve a ML/DL model in production as a microservice.


The example includes a proto file, defining the gRPC service and messages and server/client in:
- Python
- Go

The python client can speak to the Go server and vice-versa.

**Prequisites:**
- go [gRPC GO Prequisites](https://grpc.io/docs/languages/go/quickstart/#prerequisites)
- Python [gRPC Python Prequisites](https://grpc.io/docs/languages/python/quickstart/#prerequisites)

### File structure
```
├── go      (go implementation of the client and server)
│   ├── client
│   │   ├── client
│   │   └── main.go
│   ├── go.mod
│   ├── go.sum
│   ├── grpc
│   │   └── mymodelapi
│   │       ├── my_model_api_grpc.pb.go
│   │       └── my_model_api.pb.go
│   └── server
│       ├── main.go
│       └── server
├── protos      (gRPC/protobuf spec)
│   └── my_model_api.proto
├── python      (python implementation of the client and server)
│   ├── client.py
│   ├── __init__.py
│   ├── my_model_api_pb2_grpc.py
│   ├── my_model_api_pb2.py
│   └── server.py
├── create_golang_proto.sh
├── create_python_proto.sh
└── README.md
```
