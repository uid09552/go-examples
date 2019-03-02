#!/bin/bash

cd $0
ID=$(docker build  -t $1:$2  .  | tail -1 | sed 's/.*Successfully built \(.*\)$/\1/')
docker tag $(ID) $1:latest
docker tag $(ID) $1:$2
docker push $1:$2
docker push $1:latest
