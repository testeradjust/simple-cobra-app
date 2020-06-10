#!/bin/bash
set -o pipefail # abort on errors in pipeline
set -e          # abort on errors

docker-compose -f ./travis/docker-compose-kafka.yml up &
