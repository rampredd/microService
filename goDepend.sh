#!/bin/sh
DEPENDENCIES="go get github.com/micro/go-micro;go get github.com/nats-io/go-nats"

echo "getting go dependines"
eval $DEPENDENCIES
sleep 1
echo "dendencies are completed"
