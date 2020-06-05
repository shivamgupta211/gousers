#!/bin/bash

set -e

. ./export-envs.sh

DOCKER_HUB_USERNAME=${DOCKER_HUB_USERNAME}

if [ "$DOCKER_HUB_USERNAME" == "" ]
then
    echo "Please provide the Doker hub username to push the Docker Images"
    exit 1;
fi

echo "Pushing to the Docker Hub User: ${DOCKER_HUB_USERNAME}"

docker build -t ${DOCKER_HUB_USERNAME}/gousers:latest .

docker push ${DOCKER_HUB_USERNAME}/gousers:latest

exit 0
