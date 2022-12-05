# Wallet Service

## How to run 

1. Run all dependency using 
```sh
docker-compose up
```
2. To make sure all protobuf is built, Please run this line

if you are linux or mac user
```sh
make proto-wallet
```

if you are windows user
```sh
protoc deposit.proto --proto_path=protos/ --go_out=.
protoc balance.proto --proto_path=protos/ --go_out=.
protoc threshold.proto --proto_path=protos/ --go_out=.
```

3. There are three different services. Http Service, Balance Processor and Threshold Processor. You must run all 3 service at CLI to make this service fully functional.

command to run http service
```sh
go run main.go http
```

command to run balance processor service
```sh
go run main.go balance-processor
```

command to run threshold processor service
```sh
go run main.go threshold-processor
```

## Configuration

You can change serval configuration at config.yml

```yml
threshold:
  # how many minute the threshold window
  period_window: 2 
  # how much limit for the threshold
  limit: 10000

deposit:
  # topic used for deposit
  topic: "deposit"

broker:
  # list of hosts of kafka broker
  hosts: 
    - "localhost:9092"

http:
  # port used for REST endpoint
  port: "80"
```

## Project Stucture

|Folder|Description|
| ------ | ------ |
|cmd|this folder contains all configuration CLI command|
|config|this folder contains config object|
|errors|this folder contains all customer error used for this service|
|protos|this folder contains all protos definition|
|src|this folder contains all business prosess related|
|src/dtos|this folder contains data transfer object struct|
|src/endpoints|this folder contains definition of REST Service|
|src/endpoints/handlers|this folder contains definition of all REST Endpoint Handlers|
|src/models|this folder contains all generate proto model|
|src/processors|this folder contains all definition for processor|
|src/services|this folder contains all services, where all the business process writter|
|src/utils|this folder contains all helper object|