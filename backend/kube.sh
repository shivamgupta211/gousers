#!/bin/bash

set -e

. ./export-envs.sh

DOCKER_HUB_USERNAME=${DOCKER_HUB_USERNAME}

SERVER_ENV=${SERVER_ENV}

DB_CONNECTION_URL=${DB_CONNECTION_URL}

DB_USER=${DB_USER}

DB_PASSWORD=${DB_PASSWORD}

DB_DATABASE=${DB_DATABASE}

echo "Deploying Gousers Backend"
template=`cat "./kubernetes/deployment.yaml" | sed "s/{{DOCKER_HUB_USERNAME}}/$DOCKER_HUB_USERNAME/g" | sed "s/{{SERVER_ENV}}/$SERVER_ENV/g" | sed "s/{{DB_CONNECTION_URL}}/$DB_CONNECTION_URL/g" | sed "s/{{DB_USER}}/$DB_USER/g" | sed "s/{{DB_PASSWORD}}/$DB_PASSWORD/g" | sed "s/{{DB_DATABASE}}/$DB_DATABASE/g"`
kubectl apply -f ./kubernetes/service.yaml
echo "$template" | kubectl apply -f -

echo "Sleeping for 30s"
sleep 30

exit 0
