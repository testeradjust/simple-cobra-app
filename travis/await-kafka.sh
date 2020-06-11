#!/bin/bash

max_wait_seconds=60

echo "waiting for kafka start"

i=0
while [[ ${i} -le ${max_wait_seconds} ]]
do
  sleep 1
  curl localhost:9092
  # curl on kafka will produce: curl: (52) Empty reply from server
  # with exit status 52
  if [[ $? == 52 ]]; then
    exit 0
  else
    i=$(($i + 1))
    echo -n "."
  fi
done

echo "oh no, failed check for kafka/zookeeper start"
exit 1
