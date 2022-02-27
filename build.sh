#!/bin/bash

go clean --cache && go test -v -cover microservices/...

go build -o authentication/authsrv authentication/main.go

go build -o api/apisrv api/main.go
