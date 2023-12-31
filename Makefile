api:
	go run api_gateway/cmd/main.go -config=./api_gateway/config/config.yaml

proto-user:
	@echo Generating user microservice proto
	cd user_service/proto && protoc --go_out=. --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=. user.proto