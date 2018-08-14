#!/bin/sh
NATSRV="MICRO_REGISTRY=mdns ./gnatsd &"
MICROAPI="MICRO_REGISTRY=mdns micro api --handler=rpc &"
GATEWAY="MICRO_REGISTRY=mdns go run gateway/gateway.go &"
BUSINESS="MICRO_REGISTRY=mdns go run business/business.go &"
DATASTORE="MICRO_REGISTRY=mdns go run datastore/datastore.go &"

#echo "starting NATS server "
#eval $NATSRV
#sleep 1
#echo "started NATS server"
#echo "starting micro api"
#eval $MICROAPI
#sleep 1
#echo "started micro api"
echo "starting gateway"
eval $GATEWAY
sleep 1
echo "started gateway"
echo "starting business"
eval $BUSINESS
sleep 1
echo "started business"
echo "starting datastore"
eval $DATASTORE
sleep 1
echo "started datastore"
