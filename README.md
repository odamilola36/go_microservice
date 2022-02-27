# go_microservice

### Build protoc and generate go code
protoc -I=./messages --go_out=plugins=grpc:. ./messages/*.proto 