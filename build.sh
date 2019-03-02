#!/bin/bash

cd $1
ID=$(docker build  -t $2:$3  .  | tail -1 | sed 's/.*Successfully built \(.*\)$/\1/')
docker tag $2:$3 $2:latest
docker push $2:$3
docker push $2:latest
