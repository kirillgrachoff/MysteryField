
PORT_=4040
SECRET_WORD_=spagetti
FUEL_INIT_=16

run_server: cmd/server/game.go cmd/server/main.go
	cd ./cmd/server && go build -o server.bin . && PORT=$(PORT_) SECRET_WORD=$(SECRET_WORD_) FUEL_INIT=$(FUEL_INIT_) ./server.bin

run_client: cmd/client/main.go
	cd ./cmd/client && go build -o client.bin . && PORT=$(PORT_) ./client.bin
