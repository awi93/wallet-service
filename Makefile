proto-wallet:
	protoc deposit.proto --proto_path=protos/ --go_out=.
	protoc balance.proto --proto_path=protos/ --go_out=.
	protoc threshold.proto --proto_path=protos/ --go_out=.