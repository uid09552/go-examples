#!/bin/bash
function build {
    docker build . -f $1/Dockerfile -t $2:$3
    docker tag $2:$3 $2:latest
    docker push $2:$3
    docker push $2:latest
}
build greet_server $1/greetserver $2
build greet_client $1/greetclient $2