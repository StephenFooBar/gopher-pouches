#!/bin/bash
# run redis-server
REDIS=$(redis-cli ping)
if [ "$REDIS" != "PONG" ]
then
    DTSTAMP=$(date +%Y%m%d%H%M%S)
    nohup redis-server > ./log/$DTSTAMP.log 2>&1 &
fi

