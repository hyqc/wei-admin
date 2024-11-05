# yarn global add typescript ts-node @types/node # 安装 ts 和 ts-node
# yarn global add grpc-tools grpc_tools_node_protoc_ts # 安装 proto 工具到全局
# --plugin=proto-gen-ts=`which protoc-gen-ts`
cd src/proto
rm -rf admin_ts admin_proto
mkdir -p admin_ts admin_proto/google/protobuf
cp -f ../../../admin-backend/proto/admin_proto/*.proto ./admin_proto
cp -f ../../../admin-backend/third_party/google/*/* ./admin_proto/google/protobuf
cd admin_proto
protoc --plugin=proto-gen-ts --ts_out=../admin_ts *.proto --ts_opt=no_grpc,no_namespace,json_names,npm:google-protobuf