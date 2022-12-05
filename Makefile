proto-wallet:
	protoc deposit.proto --proto_path=protos/ --go_out=.
	protoc balance.proto --proto_path=protos/ --go_out=.
	protoc threshold.proto --proto_path=protos/ --go_out=.

start-http:
	go run main.go http

start-balance:
	go run main.go balance-processor

start-threshold:
	go run main.go threshold-processor