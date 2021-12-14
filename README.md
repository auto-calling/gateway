- [Atuo-calling-gateway](#atuo-calling-gateway)
  - [Feature overview](#feature-overview)
  - [Building](#building)
    - [Building the Docker image](#building-the-docker-image)
  - [Running](#running)
  - [How to work](#how-to-work)
  - [Thiết kế API](#thiết-kế-api)
  - [Message send to queue kafka](#message-send-to-queue-kafka)

# Atuo-calling-gateway

## Feature overview
Gateway receive API request from Nagios, Alert-manager, send message to kafka and insert data to MongoDB

## Building
### Building the Docker image
```
docker build .
```

## Running
Update ENV ```docker-compose.yml ```

```
docker-compose up -d
```

## How to work
- set testing env
```shell
export GIN_MODE="debug"
export TOKEN="Bearer change_me"
export MONGO_URI="mongodb://root:example@mongo:27017/?authSource=auto-calling&authMechanism=SCRAM-SHA-1"
export AUTO_CALLING_COLLECTION="logging"
export TIME_OUT="5s"
export MONGODB_MAX_POOL_SIZE=128
export MONGODB_MIN_POOL_SIZE=5
export KAFKA_BROKER_LIST="127.0.0.1:9092"
export KAFKA_TOPIC="auto-calling"
export KAFKA_GROUP_CONSUMER="auto-calling"
export MONGO_DBNAME="auto-calling"
```

## Thiết kế API
- API receive request https://{DOMAIN}/api/v1/make/event
```shell
curl --location --request POST 'https://{DOMAIN}/api/v1/make/event' \
--header 'Authorization: Bearer change_me' \
--header 'Content-Type: application/json' \
--data-raw '{
"host": "192.168.199.199",
"owner": "Admin",
"state":"CRITICAL",
"msg":"Port 8080 down",
"created":"2021-12-12 12:12:12"
"makecall":"true"
}'
```
Response success and any things else are failed
```json
{
    "data": "60adb0c82ee6f722479890e4",
    "success": true
}
```

## Message send to queue kafka
```json
{
"host":"192.168.199.199",
"owner":"Admin",
"state":"CRITICAL",
"msg":"Port 8080 down",
"created":"2021-12-12 12:12:12",
"makecall":"true"
}
```