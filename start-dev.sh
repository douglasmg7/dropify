#!/usr/bin/env bash
# [[ `systemctl status redis | awk '/Active/{print $2}'` == inactive ]] && sudo systemctl start redis

# Start dockerizeds.
$ZUNKA_DOCKER_SCRIPTS/start-mongo.sh
$ZUNKA_DOCKER_SCRIPTS/start-redis.sh

go build && ./dropify