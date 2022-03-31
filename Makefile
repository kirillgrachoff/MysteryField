
PORT_=12345
SECRET_WORD_=cheburashka
FUEL_BUDGET_=10
HOST_=localhost

build: pkg/field.proto
	cd pkg && protoc field.proto --go_out=. --go-grpc_out=.

run_server: cmd/server/game.go cmd/server/main.go
	cd ./cmd/server && go build -o server.bin . && PORT=$(PORT_) SECRET_WORD=$(SECRET_WORD_) FUEL_BUDGET=$(FUEL_BUDGET_) ./server.bin

run_client: cmd/client/main.go
	cd ./cmd/client && go build -o client.bin . && PORT=$(PORT_) HOST=$(HOST_) ./client.bin
