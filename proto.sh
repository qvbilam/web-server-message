function proto {
  # 脚本所在路径
  SCRIPT_DIR=$(cd $(dirname "$0");pwd)
  DOMAIN=.
  PROTO_FILE=$1.proto
  VERSION=$2
  PROTO_PATH=${DOMAIN}/api/qvbilam/video/"$VERSION"
  OUT_PATH=./${DOMAIN}/api/qvbilam/video/"$VERSION"
  # 引入项目目录下不同级的 proto 需要指定参数 --proto_path=绝对路径
  protoc -I="$PROTO_PATH" --go_out "$OUT_PATH" --go_opt paths=source_relative --go-grpc_out "$OUT_PATH" --go-grpc_opt=paths=source_relative "$PROTO_FILE" --proto_path="$SCRIPT_DIR"
}

function userProto {
  # 脚本所在路径
  SCRIPT_DIR=$(cd $(dirname "$0");pwd)
  DOMAIN=.
  PROTO_FILE=$1.proto
  VERSION=$2
  PROTO_PATH=${DOMAIN}/api/qvbilam/user/"$VERSION"
  OUT_PATH=./${DOMAIN}/api/qvbilam/user/"$VERSION"
  # 引入项目目录下不同级的 proto 需要指定参数 --proto_path=绝对路径
  protoc -I="$PROTO_PATH" --go_out "$OUT_PATH" --go_opt paths=source_relative --go-grpc_out "$OUT_PATH" --go-grpc_opt=paths=source_relative "$PROTO_FILE" --proto_path="$SCRIPT_DIR"
}

function pageProto {
  # 脚本所在路径
  SCRIPT_DIR=$(cd $(dirname "$0");pwd)
  DOMAIN=.
  PROTO_FILE=$1.proto
  VERSION=$2
  PROTO_PATH=${DOMAIN}/api/qvbilam/page/"$VERSION"
  OUT_PATH=./${DOMAIN}/api/qvbilam/page/"$VERSION"
  # 引入项目目录下不同级的 proto 需要指定参数 --proto_path=绝对路径
  protoc -I="$PROTO_PATH" --go_out "$OUT_PATH" --go_opt paths=source_relative --go-grpc_out "$OUT_PATH" --go-grpc_opt=paths=source_relative "$PROTO_FILE" --proto_path="$SCRIPT_DIR"
}

pageProto page v1
userProto user v1

proto category v1
proto drama v1
proto region v1
proto video v1


