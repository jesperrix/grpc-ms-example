SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
PYTHON_PROTO_DIR=$SCRIPT_DIR/python
PROTO_DIR=$SCRIPT_DIR/protos

mkdir -p $PYTHON_PROTO_DIR

echo "Creating Python files for protobuf and gRPC"
echo "output dir: $PYTHON_PROTO_DIR"

pushd $PYTHON_PROTO_DIR
python -m grpc_tools.protoc -I${PROTO_DIR} --python_out=. --grpc_python_out=. ${PROTO_DIR}/my_model_api.proto
popd

