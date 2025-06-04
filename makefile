.PHONY: gen
gen:
	protoc --go_out=. --go-grpc_out=. api/messenger_user.proto