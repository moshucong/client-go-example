#!/bin/bash

docker build -t 172.30.10.185:15000/gears/client-go-example:v0.0.1 .
docker push 172.30.10.185:15000/gears/client-go-example:v0.0.1

