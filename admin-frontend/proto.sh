# https://www.npmjs.com/package/ts-proto/v/1.79.0#quickstart
# npm install ts-proto -g

rm -rf src/proto
mkdir -p src/proto/admin_proto/google/protobuf
mkdir -p src/proto/admin_ts

cp -f ../admin-backend/proto/admin_proto/*.proto ./src/proto/admin_proto
cp -f ../admin-backend/third_party/google/*/* ./src/proto/admin_proto/google/protobuf

cd ./src/proto/admin_proto

protoc  --plugin=`where protoc-gen-ts_proto | grep proto$` --ts_proto_opt=useOptionals=all --ts_proto_out=../admin_ts *.proto  