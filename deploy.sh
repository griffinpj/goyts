#!/bin/bash

docker buildx build --platform linux/amd64 -t goyts:latest . && docker tag goyts:latest cougargriff/goyts:latest
