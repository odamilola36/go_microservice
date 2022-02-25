#!/bin/bash

go clean --cache && go test -v -cover microservices/authentication/...

go build -o authentication/authsrv authentication/main.go
