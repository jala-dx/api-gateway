Define the Proto file
Compile the Protofile first
Compile the Protofile again


Follow this https://github.com/grpc-ecosystem/grpc-gateway

protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:. atom_service.proto ;protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:. atom_service.proto

Bkend and client code, refer this

https://github.com/grpc/grpc-go/tree/master/examples/helloworld

curl -s -XPOST -H "Content-Type: application/json" -d '{"value":"Hello atom"}' http://localhost:5050/v1/atomfrontendgw/echo | json_pp

or 

./atomclient

root@myvm:/vagrant/src/cto-github.cisco.com# ps -ef | grep atom
root      5248 28323  0 Nov29 pts/5    00:00:00 ./atombkend
root      5286 28323  0 Nov29 pts/5    00:00:00 ./atomgrpc-main -atomflag localhost:50051


start the grpcgw , by specifying the backend details as params.
- ./atomgrpc-main -atomflag localhost:50051

in here you are passing command line params saying localhost:50051 if your backend

start the backend
Issue the curl request to grpcgw
run the client 

