SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
GO_GRPC_DIR=$SCRIPT_DIR/go/grpc
PROTO_DIR=$SCRIPT_DIR/protos

mkdir -p $GO_GRPC_DIR

echo "Creating go files gRPC"
echo "output dir: $GO_GRPC_DIR"

pushd $GO_GRPC_DIR
protoc -I$PROTO_DIR --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    $PROTO_DIR/my_model_api.proto
popd

