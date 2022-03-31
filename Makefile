
PORT_=12345
SECRET_WORD_=banana
FUEL_BUDGET_=10

build: pkg/field.proto:
	cd pkg && protoc field.proto --go_out=. --go-grpc_out=.

run_server: cmd/server/game.go cmd/server/main.go
	cd ./cmd/server && go build -o server.bin . && PORT=$(PORT_) SECRET_WORD=$(SECRET_WORD_) FUEL_INIT=$(FUEL_BUDGET_) ./server.bin

run_client: cmd/client/main.go
	cd ./cmd/client && go build -o client.bin . && PORT=$(PORT_) ./client.bin
