#!/usr/bin/env bash

echo "Building image locally..."
docker build --no-cache -t "manager" -f ./Dockerfile .

echo "Deploying agent docker..."
docker run -d --name testdev \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -v /var/lib/docker/volumes:/var/lib/docker/volumes \
  -p 8080:8080 \
  "manager"

echo "Starting log"
docker logs -f testdev