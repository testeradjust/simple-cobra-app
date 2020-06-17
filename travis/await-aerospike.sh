#!/bin/bash

max_wait_seconds=60

echo "waiting for aerospike start"

i=0
while [[ ${i} -le ${max_wait_seconds} ]]
do
  sleep 1
  curl localhost:3000
  # curl on aerospike will produce: curl: (56) Recv failure: Connection reset by peer
  # with exit status 56
  if [[ $? == 56 ]]; then
    echo "aerospike seems started <3"
    exit 0
  else
    i=$(($i + 1))
    echo -n "."
  fi
done

echo "oh no, failed check for aerospike start"
exit 1
