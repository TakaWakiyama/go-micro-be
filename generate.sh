
service=user
root=$(git rev-parse --show-toplevel)
cd $root

generate() {
    protoc proto/${1}.proto \
    -I=proto \
    --go_out=cmd/$1/pb \
    --go-grpc_out=cmd/$1/pb \
    --go-grpc_opt=require_unimplemented_servers=false
}

generate ${service}