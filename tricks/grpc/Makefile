.PHONY: compile
compile: ## Compile the proto file.
	protoc --go_out=pkg --go_opt=paths=source_relative --go-grpc_out=pkg --go-grpc_opt=paths=source_relative api/proto/credit.proto

.PHONY: server
server: ## Build and run server.
	go build -race -ldflags "-s -w" -o bin/server cmd/server/main.go
	bin/server

.PHONY: client
client: ## Build and run client.
	go build -race -ldflags "-s -w" -o bin/client client/main.go
	bin/client

.PHONY: client_evans
client_evans: ## Build and run client.
	evans api/proto/credit.proto -p 3200




