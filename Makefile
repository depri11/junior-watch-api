api:
	go run api_gateway/cmd/main.go -config=./api_gateway/config/config.yaml

user:
	go run user_service/cmd/main.go -config=./user_service/config/config.yaml

gen-proto:
	@echo Generating user microservice proto
	cd pkg/proto && protoc --go_out=. --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=. user.proto