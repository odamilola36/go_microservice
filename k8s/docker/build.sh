#!/bin/bash

cp ../../api/apisrv .
cp ../../authentication/authsrv .

docker build -t microservice-v1 .
docker inspect microservice-v1