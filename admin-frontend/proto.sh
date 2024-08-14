# yarn global add typescript ts-node @types/node # 安装 ts 和 ts-node
# yarn global add grpc-tools grpc_tools_node_protoc_ts # 安装 proto 工具到全局
# --plugin=proto-gen-ts=`which protoc-gen-ts`
cd src/proto
rm -rf admin_ts
mkdir admin_ts
cd admin_proto
protoc --plugin=proto-gen-ts --ts_out=../admin_ts *.proto

