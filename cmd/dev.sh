
if [[ $1 == "generate" ]];then
    root=$(git rev-parse --show-toplevel)
    protoc -I=${root}/cmd/user \
    --go_out=${root}/cmd/user/pb \
    --go-grpc_out=${root}/cmd/user/pb \
    --go-grpc_opt=require_unimplemented_servers=false \
    ${root}/cmd/user/user.proto
fi