#!/bin/bash

PATH=$PATH:/Users/Max/go-examples/docker/gin/bin/
protoc $1 --go_out=plugins=grpc:.