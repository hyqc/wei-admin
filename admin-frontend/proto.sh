# https://www.npmjs.com/package/ts-proto/v/1.79.0#quickstart
# npm install ts-proto
cd src/proto
rm -rf admin_ts admin_proto
mkdir -p admin_ts admin_proto/google/protobuf
cp -f ../../../admin-backend/proto/admin_proto/*.proto ./admin_proto
cp -f ../../../admin-backend/third_party/google/*/* ./admin_proto/google/protobuf
cd admin_proto
protoc --plugin=protoc-gen-ts_proto --ts_proto_out=../admin_ts ./*.proto --ts_proto_opt=outputEncodeMethods=false,outputJsonMethods=false,outputClientImpl=false