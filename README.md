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

### How to run 
Specifies how to run the client and servers.

You can play around with running the Python server and Go Client or the other way around.

##### Go 
How to run the Go server/client.

###### Server
```
# Make sure you are in the root of the repo
> cd go
> go run server/main.go
2020/11/24 10:35:37 Starting listener at: %s localhost:56789
2020/11/24 10:35:37 Registering gRPC server to listener
```

The server is now ready to receive messages from a client.


###### Client
```
# Make sure you are in the root of the repo
> cd go
> go run client/main.go
2020/11/24 10:35:39 sent/received 0 messages.
2020/11/24 10:35:39 model_id:"model_xyz" y:1 y:2 y:3 y:1
2020/11/24 10:35:39 sent/received 100 messages.
2020/11/24 10:35:39 model_id:"model_xyz" y:1 y:2 y:3 y:101
2020/11/24 10:35:39 sent/received 200 messages.
....
```
#### Python
How to run the Python server/client.

##### Server
```
# Make sure you are in the root of the repo
> python python/server
INFO:root:Starting server at [::]:56789
```

The server is now ready to receive messages from a client.

##### Client
```
# Make sure you are in the root of the repo
> python python/client.py
INFO:root:sent/recieved 0 messages.
INFO:root:model_id=model_xyz, y=[1, 2, 3, 1]
INFO:root:sent/recieved 100 messages.
INFO:root:model_id=model_xyz, y=[1, 2, 3, 101]
INFO:root:sent/recieved 200 messages.
INFO:root:model_id=model_xyz, y=[1, 2, 3, 201]
....
```
